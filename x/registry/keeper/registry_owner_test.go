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

func createNRegistryOwner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RegistryOwner {
	items := make([]types.RegistryOwner, n)
	for i := range items {
		items[i].Id = keeper.AppendRegistryOwner(ctx, items[i])
	}
	return items
}

func TestRegistryOwnerGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryOwner(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRegistryOwner(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRegistryOwnerRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryOwner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRegistryOwner(ctx, item.Id)
		_, found := keeper.GetRegistryOwner(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRegistryOwnerGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryOwner(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRegistryOwner(ctx)),
	)
}

func TestRegistryOwnerCount(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryOwner(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRegistryOwnerCount(ctx))
}
