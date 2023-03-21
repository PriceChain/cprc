package prcmint

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/PriceChain/cprc/x/prcmint/keeper"
	"github.com/PriceChain/cprc/x/prcmint/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, ic types.InflationCalculationFn) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	totalSupply := k.TokenSupply(ctx, params.MintDenom)

	// Get the total amount of rewards to be minted newly
	rewardsTokenAmount := CalculateRewards(ctx, k, totalSupply, params.BlocksPerYear)

	// If there is rewards to be minted
	if rewardsTokenAmount.GT(sdk.ZeroInt()) {
		rewardsTokens := sdk.NewCoin(params.MintDenom, rewardsTokenAmount)
		mintedRewardCoins := sdk.NewCoins(rewardsTokens)

		// telemetry log
		if rewardsTokens.Amount.IsInt64() {
			defer telemetry.ModuleSetGauge(types.ModuleName, float32(rewardsTokens.Amount.Int64()), "minted_tokens_rewards")
		}

		// Mint Coins for rewards of price validators
		err := k.MintCoins(ctx, mintedRewardCoins)
		if err != nil {
			panic(err)
		}

		// send the minted coins to the fee registry fee collector account
		err = k.AddCollectedFeesToRegistry(ctx, mintedRewardCoins)
		if err != nil {
			panic(err)
		}

		// Emit Rewards minted events
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMint,
				sdk.NewAttribute(types.AttributeKeyRewards, rewardsTokenAmount.String()),
			),
		)
	}

	// recalculate inflation rate
	totalStakingSupply := k.StakingTokenSupply(ctx)

	registryTotalStakedSupply := sdk.ZeroInt()
	registryTotalStakedAmount, bFound := k.GetRegistryStakedAmount(ctx, "total")
	if bFound {
		// Parse total staked amount
		amount, err := strconv.ParseInt(registryTotalStakedAmount.Amount, 10, 64)
		if err == nil {
			registryTotalStakedSupply = sdk.NewInt(amount)
		}
	}

	// Increase total staked supply
	totalStakingSupply = totalStakingSupply.Add(registryTotalStakedSupply)

	// BondedRatio the fraction of the staking tokens which are currently bonded
	bondedRatio := sdk.NewDecFromInt(k.TotalBondedTokens(ctx)).QuoInt(totalStakingSupply)

	minter.Inflation = ic(ctx, minter, params, bondedRatio, totalSupply)
	// If no inflation, just skip
	if minter.Inflation == sdk.ZeroDec() {
		return
	}

	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)
	k.SetMinter(ctx, minter)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
			sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
			sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}

// Calculate rewards per wallet
func CalculateRewards(ctx sdk.Context, k keeper.Keeper, totalSupply sdk.Int, blocksPerYear uint64) sdk.Int {
	allRegistries := k.GetAllRegistry(ctx)

	totalStakedAmount := uint64(0)
	// Loop all registries
	for _, ar := range allRegistries {
		sa, _ := strconv.ParseUint(ar.StakedAmount, 10, 64)
		totalStakedAmount += sa
	}

	// Get total staked tokens by registry owners
	fTotalStakedToken := float64(totalStakedAmount)

	// All price data
	allPriceData := k.GetAllPriceData(ctx)
	allRegistryMembers := k.GetAllRegistryMember(ctx)

	numberOfBlocks := (float64)(ctx.BlockHeight())
	fTotalPoPSent := (float64)(len(allPriceData))

	registryTotalStakedAmount, bFound := k.GetRegistryStakedAmount(ctx, "total")
	if !bFound {
		return sdk.ZeroInt()
	}

	// Parse total staked amount
	amount, err := strconv.ParseUint(registryTotalStakedAmount.Amount, 10, 64)
	if err != nil {
		return sdk.ZeroInt()
	}

	newlyMinted := int64(0)
	fTotalDeposited := (float64)(amount)
	for _, pv := range allRegistryMembers {
		registryId, _ := strconv.ParseUint(pv.RegistryId, 10, 64)
		_, bFound := k.GetRegistry(ctx, registryId)
		if !bFound {
			continue
		}

		// PoP Sent by wallet
		popCount, _ := strconv.ParseUint(pv.PopCount, 10, 64)
		fPopCount := (float64)(popCount)

		// Total staked amount per wallet
		stakedAmountPerWallet, bFound := k.GetRegistryStakedAmountPerWallet(ctx, pv.Wallet)
		if !bFound {
			continue
		}

		// Total staked amount per wallet
		stakedAmount, _ := strconv.ParseUint(stakedAmountPerWallet.StakedAmount, 10, 64)
		fStakedAmount := (float64)(stakedAmount)

		// Rewards calculation
		fRewards := fTotalDeposited * math.Pow(0.5, (numberOfBlocks/float64(blocksPerYear))) * ((fPopCount / fTotalPoPSent) * (fStakedAmount / fTotalStakedToken))
		rewards := (int64)(fRewards)
		newlyMinted += rewards

		prevRewardSum, _ := strconv.ParseInt(pv.RewardSum, 10, 64)
		// Set RewardSum
		pv.RewardSum = fmt.Sprintf("%d", prevRewardSum+rewards)

		// Update Registry Member Rewards
		k.SetRegistryMember(ctx, pv)
	}

	// return the amount tokens to be minted newly
	return sdk.NewInt(newlyMinted)
}
