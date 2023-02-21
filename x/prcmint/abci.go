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

	// recalculate inflation rate
	totalStakingSupply := k.StakingTokenSupply(ctx)
	totalSupply := k.TokenSupply(ctx, params.MintDenom)
	bondedRatio := k.BondedRatio(ctx)
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

func CalculateRewards(ctx sdk.Context, k keeper.Keeper, totalSupply sdk.Int, blocksPerYear uint64) {
	allPriceData := k.GetAllPriceData(ctx)
	allRegistryMembers := k.GetAllRegistryMember(ctx)

	numberOfBlocks := (float64)(ctx.BlockHeight())
	fTotalPoPSent := (float64)(len(allPriceData))

	registryStakedAmount, bFound := k.GetRegistryStakedAmount(ctx, "total")
	if !bFound {
		return
	}

	// Parse total staked amount
	amount, err := strconv.ParseUint(registryStakedAmount.Amount, 10, 64)
	if err != nil {
		return
	}

	fTotalStakedToken := (float64)(amount)
	for _, pv := range allRegistryMembers {
		registryId, _ := strconv.ParseUint(pv.RegistryId, 10, 64)
		registry, bFound := k.GetRegistry(ctx, registryId)
		if !bFound {
			continue
		}

		// Total staked token at the registry
		totalDeposited, _ := strconv.ParseUint(registry.StakedAmount, 10, 64)
		fTotalDeposited := (float64)(totalDeposited)

		// PoP Sent by wallet
		popCount, _ := strconv.ParseUint(pv.PopCount, 10, 64)
		fPopCount := (float64)(popCount)

		// Total staked amount per wallet
		stakedAmountPerWallet, bFound := k.GetStakedAmountPerWallet(ctx, pv.Wallet)
		if !bFound {
			continue
		}

		// Total staked amount per wallet
		stakedAmount, _ := strconv.ParseUint(stakedAmountPerWallet.StakedAmount, 10, 64)
		fStakedAmount := (float64)(stakedAmount)

		// Rewards calculation
		fRewards := fTotalDeposited * math.Pow(0.5, (numberOfBlocks/(float64)(blocksPerYear))) * ((fPopCount / fTotalPoPSent) * (fStakedAmount / fTotalStakedToken))
		Rewards := (uint64)(fRewards)

		// Set RewardSum
		pv.RewardSum = fmt.Sprintf("%d", Rewards)
	}
}
