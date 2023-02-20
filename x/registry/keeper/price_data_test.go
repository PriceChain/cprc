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

func createNPriceData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PriceData {
	items := make([]types.PriceData, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPriceData(ctx, items[i])
	}
	return items
}

func TestPriceDataGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNPriceData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPriceData(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPriceDataRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNPriceData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePriceData(ctx,
			item.Index,
		)
		_, found := keeper.GetPriceData(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPriceDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNPriceData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPriceData(ctx)),
	)
}
