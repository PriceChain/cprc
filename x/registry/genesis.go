package registry

import (
	"github.com/PriceChain/rd_net/x/registry/keeper"
	"github.com/PriceChain/rd_net/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the registry
	for _, elem := range genState.RegistryList {
		k.SetRegistry(ctx, elem)
	}

	// Set registry count
	k.SetRegistryCount(ctx, genState.RegistryCount)
	// Set all the registryOwner
	for _, elem := range genState.RegistryOwnerList {
		k.SetRegistryOwner(ctx, elem)
	}

	// Set registryOwner count
	k.SetRegistryOwnerCount(ctx, genState.RegistryOwnerCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RegistryList = k.GetAllRegistry(ctx)
	genesis.RegistryCount = k.GetRegistryCount(ctx)
	genesis.RegistryOwnerList = k.GetAllRegistryOwner(ctx)
	genesis.RegistryOwnerCount = k.GetRegistryOwnerCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
