package types

import (
	rtypes "github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// StakingKeeper defines the expected staking keeper
type StakingKeeper interface {
	StakingTokenSupply(ctx sdk.Context) sdk.Int
	BondedRatio(ctx sdk.Context) sdk.Dec
	TotalBondedTokens(ctx sdk.Context) sdk.Int
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress

	// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862
	SetModuleAccount(sdk.Context, types.ModuleAccountI)
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

// RegistryKeeper defines expected interface
type RegistryKeeper interface {
	GetAllRegistry(ctx sdk.Context) (list []rtypes.Registry)
	GetAllRegistryMember(ctx sdk.Context) (list []rtypes.RegistryMember)
	GetAllPriceData(ctx sdk.Context) (list []rtypes.PriceData)
	GetRegistry(sdk.Context, uint64) (rtypes.Registry, bool)
	GetRegistryStakedAmount(sdk.Context, string) (rtypes.RegistryStakedAmount, bool)
	GetAllRegistryStakedAmount(ctx sdk.Context) (list []rtypes.RegistryStakedAmount)
	GetRegistryStakedAmountPerWallet(sdk.Context, string) (rtypes.RegistryStakedAmountPerWallet, bool)
	GetAllRegistryStakedAmountPerWallet(ctx sdk.Context) (list []rtypes.RegistryStakedAmountPerWallet)
	SetRegistryMember(sdk.Context, rtypes.RegistryMember)
}
