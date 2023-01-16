package keeper_test

import (
	"testing"

	keepertest "github.com/PriceChain/cprc/testutil/keeper"
	"github.com/PriceChain/cprc/testutil/nullify"
	"github.com/PriceChain/cprc/x/registry/keeper"
	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNRegistry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Registry {
	items := make([]types.Registry, n)
	for i := range items {
		items[i].Id = keeper.AppendRegistry(ctx, items[i])
	}
	return items
}

func TestRegistryGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistry(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRegistry(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRegistryRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRegistry(ctx, item.Id)
		_, found := keeper.GetRegistry(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRegistryGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRegistry(ctx)),
	)
}

func TestRegistryCount(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistry(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRegistryCount(ctx))
}
