package registry

import (
	"github.com/PriceChain/cprc/x/registry/keeper"
	"github.com/PriceChain/cprc/x/registry/types"
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
	// Set all the registryMember
	for _, elem := range genState.RegistryMemberList {
		k.SetRegistryMember(ctx, elem)
	}

	// Set registryMember count
	k.SetRegistryMemberCount(ctx, genState.RegistryMemberCount)
	// Set all the priceConsensus
	for _, elem := range genState.PriceConsensusList {
		k.SetPriceConsensus(ctx, elem)
	}

	// Set priceConsensus count
	k.SetPriceConsensusCount(ctx, genState.PriceConsensusCount)
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
	genesis.RegistryMemberList = k.GetAllRegistryMember(ctx)
	genesis.RegistryMemberCount = k.GetRegistryMemberCount(ctx)
	genesis.PriceConsensusList = k.GetAllPriceConsensus(ctx)
	genesis.PriceConsensusCount = k.GetPriceConsensusCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
