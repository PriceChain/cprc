package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/PriceChain/cprc/testutil/keeper"
	"github.com/PriceChain/cprc/testutil/nullify"
	"github.com/PriceChain/cprc/x/registry/keeper"
	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRegistryStakedAmount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RegistryStakedAmount {
	items := make([]types.RegistryStakedAmount, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRegistryStakedAmount(ctx, items[i])
	}
	return items
}

func TestRegistryStakedAmountGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryStakedAmount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRegistryStakedAmount(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRegistryStakedAmountRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryStakedAmount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRegistryStakedAmount(ctx,
			item.Index,
		)
		_, found := keeper.GetRegistryStakedAmount(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRegistryStakedAmountGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryStakedAmount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRegistryStakedAmount(ctx)),
	)
}
