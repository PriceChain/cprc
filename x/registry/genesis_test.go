package registry_test

import (
	"testing"

	keepertest "github.com/PriceChain/rd_net/testutil/keeper"
	"github.com/PriceChain/rd_net/testutil/nullify"
	"github.com/PriceChain/rd_net/x/registry"
	"github.com/PriceChain/rd_net/x/registry/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RegistryList: []types.Registry{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RegistryCount: 2,
		RegistryOwnerList: []types.RegistryOwner{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RegistryOwnerCount: 2,
		RegistryMemberList: []types.RegistryMember{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RegistryMemberCount: 2,
		PriceConsensusList: []types.PriceConsensus{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PriceConsensusCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RegistryKeeper(t)
	registry.InitGenesis(ctx, *k, genesisState)
	got := registry.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RegistryList, got.RegistryList)
	require.Equal(t, genesisState.RegistryCount, got.RegistryCount)
	require.ElementsMatch(t, genesisState.RegistryOwnerList, got.RegistryOwnerList)
	require.Equal(t, genesisState.RegistryOwnerCount, got.RegistryOwnerCount)
	require.ElementsMatch(t, genesisState.RegistryMemberList, got.RegistryMemberList)
	require.Equal(t, genesisState.RegistryMemberCount, got.RegistryMemberCount)
	require.ElementsMatch(t, genesisState.PriceConsensusList, got.PriceConsensusList)
	require.Equal(t, genesisState.PriceConsensusCount, got.PriceConsensusCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
