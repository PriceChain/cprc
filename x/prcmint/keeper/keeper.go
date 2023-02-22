package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/PriceChain/cprc/x/prcmint/types"
	rtypes "github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc                      codec.BinaryCodec
		storeKey                 storetypes.StoreKey
		paramSpace               paramtypes.Subspace
		stakingKeeper            types.StakingKeeper
		bankKeeper               types.BankKeeper
		registryKeeper           types.RegistryKeeper
		feeCollectorName         string
		registryFeeCollectorName string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, ps paramtypes.Subspace,
	sk types.StakingKeeper, ak types.AccountKeeper, bk types.BankKeeper, rk types.RegistryKeeper,
	feeCollectorName string, registryFeeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:                      cdc,
		storeKey:                 key,
		paramSpace:               ps,
		stakingKeeper:            sk,
		bankKeeper:               bk,
		registryKeeper:           rk,
		feeCollectorName:         feeCollectorName,
		registryFeeCollectorName: registryFeeCollectorName,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// get the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

// set the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.MinterKey, b)
}

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// StakingTokenSupply to be used in BeginBlocker.
func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdk.Int {
	return k.stakingKeeper.StakingTokenSupply(ctx)
}

// BondedRatio implements an alias call to the underlying staking keeper's
// BondedRatio to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	return k.stakingKeeper.BondedRatio(ctx)
}

// BondedRatio the fraction of the staking tokens which are currently bonded
func (k Keeper) TotalBondedTokens(ctx sdk.Context) sdk.Int {
	return k.stakingKeeper.TotalBondedTokens(ctx)
}

// TokenSupply implements an alias call to the underlying bank keeper's
// TokenSupply to be used in BeginBlocker.
func (k Keeper) TokenSupply(ctx sdk.Context, denom string) sdk.Int {
	return k.bankKeeper.GetSupply(ctx, denom).Amount
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) AddCollectedFees(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, fees)
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) AddCollectedFeesToRegistry(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.registryFeeCollectorName, fees)
}

// Get All registries
func (k Keeper) GetAllRegistry(ctx sdk.Context) (list []rtypes.Registry) {
	return k.registryKeeper.GetAllRegistry(ctx)
}

// Get All staked amount per wallet
func (k Keeper) GetAllStakedAmountPerWallet(ctx sdk.Context) (list []rtypes.StakedAmountPerWallet) {
	return k.registryKeeper.GetAllStakedAmountPerWallet(ctx)
}

// Get all registry staked amount
func (k Keeper) GetAllRegistryStakedAmount(ctx sdk.Context) (list []rtypes.RegistryStakedAmount) {
	return k.registryKeeper.GetAllRegistryStakedAmount(ctx)
}

// Get all registry memeber - price validators
func (k Keeper) GetAllRegistryMember(ctx sdk.Context) (list []rtypes.RegistryMember) {
	return k.registryKeeper.GetAllRegistryMember(ctx)
}

// Get all price data
func (k Keeper) GetAllPriceData(ctx sdk.Context) (list []rtypes.PriceData) {
	return k.registryKeeper.GetAllPriceData(ctx)
}

// Get registry using Id
func (k Keeper) GetRegistry(ctx sdk.Context, index uint64) (rtypes.Registry, bool) {
	return k.registryKeeper.GetRegistry(ctx, index)
}

// Get registry staked amount using Id
func (k Keeper) GetRegistryStakedAmount(ctx sdk.Context, index string) (rtypes.RegistryStakedAmount, bool) {
	return k.registryKeeper.GetRegistryStakedAmount(ctx, index)
}

// Get total staked amount per wallet
func (k Keeper) GetStakedAmountPerWallet(ctx sdk.Context, index string) (rtypes.StakedAmountPerWallet, bool) {
	return k.registryKeeper.GetStakedAmountPerWallet(ctx, index)
}

// Update registry member
func (k Keeper) SetRegistryMember(ctx sdk.Context, rm rtypes.RegistryMember) {
	k.registryKeeper.SetRegistryMember(ctx, rm)
}
