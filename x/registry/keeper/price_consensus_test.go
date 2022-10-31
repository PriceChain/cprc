package keeper_test

import (
	"testing"

	keepertest "github.com/PriceChain/rd_net/testutil/keeper"
	"github.com/PriceChain/rd_net/testutil/nullify"
	"github.com/PriceChain/rd_net/x/registry/keeper"
	"github.com/PriceChain/rd_net/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNPriceConsensus(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PriceConsensus {
	items := make([]types.PriceConsensus, n)
	for i := range items {
		items[i].Id = keeper.AppendPriceConsensus(ctx, items[i])
	}
	return items
}

func TestPriceConsensusGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNPriceConsensus(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPriceConsensus(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPriceConsensusRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNPriceConsensus(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePriceConsensus(ctx, item.Id)
		_, found := keeper.GetPriceConsensus(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPriceConsensusGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNPriceConsensus(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPriceConsensus(ctx)),
	)
}

func TestPriceConsensusCount(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNPriceConsensus(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPriceConsensusCount(ctx))
}
