package prcmint

import (
	"github.com/PriceChain/cprc/x/prcmint/keeper"
	"github.com/PriceChain/cprc/x/prcmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, ak types.AccountKeeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.SetMinter(ctx, genState.Minter)

	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Minter = k.GetMinter(ctx)

	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
