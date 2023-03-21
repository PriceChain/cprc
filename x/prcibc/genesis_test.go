package prcibc_test

import (
	"testing"

	keepertest "github.com/PriceChain/cprc/testutil/keeper"
	"github.com/PriceChain/cprc/testutil/nullify"
	"github.com/PriceChain/cprc/x/prcibc"
	"github.com/PriceChain/cprc/x/prcibc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		IbcMsgList: []types.IbcMsg{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		IbcMsgCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PrcibcKeeper(t)
	prcibc.InitGenesis(ctx, *k, genesisState)
	got := prcibc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.IbcMsgList, got.IbcMsgList)
	require.Equal(t, genesisState.IbcMsgCount, got.IbcMsgCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
