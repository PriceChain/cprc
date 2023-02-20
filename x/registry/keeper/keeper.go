package keeper

import (
	"fmt"
	"strconv"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	// ensure registry stake collector module account is set
	if addr := ak.GetModuleAddress(types.RegistryStakeCollectorName); addr == nil {
		panic("the registry stake collect account has not been set")
	}

	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Preprocesssing
func (k msgServer) PreProcess(ctx sdk.Context, sender string, stakeAmount string) (sdk.AccAddress, sdk.Int, sdk.Coins, error) {
	// Parse creator bech32 address
	creator, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return sdk.AccAddress{}, sdk.ZeroInt(), sdk.Coins{}, err
	}

	// Parse amount of CPRC token
	collateral, err := sdk.ParseCoinsNormalized(stakeAmount)
	if err != nil {
		return sdk.AccAddress{}, sdk.ZeroInt(), sdk.Coins{}, err
	}

	// Get denomination
	denom := collateral.GetDenomByIndex(0)

	// Get Int amount
	amount := collateral.AmountOf(denom)

	// If equals to 0
	if amount.Equal(sdk.ZeroInt()) {
		return sdk.AccAddress{}, sdk.ZeroInt(), sdk.Coins{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid token amount")
	}

	// fetch stored params
	params := k.GetParams(ctx)

	// Parse amount of Minstake
	minStakeCoin, err := sdk.ParseCoinsNormalized(params.MinimumStakeAmount)
	if err != nil {
		return sdk.AccAddress{}, sdk.ZeroInt(), sdk.Coins{}, err
	}

	// Get denomination
	denomMinStake := minStakeCoin.GetDenomByIndex(0)

	// Get Int amount
	amtMinStake := minStakeCoin.AmountOf(denomMinStake)

	// if it is below than minimum stake amount
	if denom != denomMinStake || amount.LT(amtMinStake) {
		return sdk.AccAddress{}, sdk.ZeroInt(), sdk.Coins{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "It shouldn't be below than minimum staking amount!")
	}

	// return normal value
	return creator, amount, collateral, nil
}

// Update Total Staked amount
func (k msgServer) UpdateTotalStakedAmount(ctx sdk.Context, stakeAmount string) bool {
	amount, err := strconv.ParseUint(stakeAmount, 10, 64)
	if err != nil {
		return false
	}

	// Fetch previous total staked amount
	prevStakedAmount := (uint64)(0)
	stakedAmount, bFound := k.GetRegistryStakedAmount(ctx, "total")
	if bFound {
		amt, _ := strconv.ParseUint(stakedAmount.Amount, 10, 64)
		prevStakedAmount = amt
	}

	// Initalize a new total staked amount item
	rsa := types.RegistryStakedAmount{
		Index:  "total",
		Amount: fmt.Sprintf("%d", prevStakedAmount+amount),
	}

	// Update total staked amount
	k.SetRegistryStakedAmount(ctx, rsa)

	return true
}

// Update staked amount per wallet
func (k msgServer) UpdateStakedAmountPerWallet(ctx sdk.Context, sender string, stakeAmount string) bool {
	amount, err := strconv.ParseUint(stakeAmount, 10, 64)
	if err != nil {
		return false
	}

	// Staked amount per wallet
	stakedAmtPerWallet := k.GetAllStakedAmountPerWallet(ctx)

	// Find if there is already existing wallet
	foundIndex := -1
	for i, s := range stakedAmtPerWallet {
		// If already exists
		if s.Index == sender {
			foundIndex = i
		}
	}

	// Initialize Staked amount per wallet value
	sapw := types.StakedAmountPerWallet{
		Index:        sender,
		Wallet:       sender,
		StakedAmount: stakeAmount,
	}

	// If there is already an item found
	if foundIndex > -1 {
		sapw = stakedAmtPerWallet[foundIndex]

		// Sum the amount
		amt, _ := strconv.ParseUint(sapw.StakedAmount, 10, 64)
		amt += amount

		sapw.StakedAmount = fmt.Sprintf("%d", amt)
	}

	// Update Keeper Store
	k.SetStakedAmountPerWallet(ctx, sapw)

	return true
}
