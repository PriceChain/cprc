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

func createNStakedAmountPerWallet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StakedAmountPerWallet {
	items := make([]types.StakedAmountPerWallet, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStakedAmountPerWallet(ctx, items[i])
	}
	return items
}

func TestStakedAmountPerWalletGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNStakedAmountPerWallet(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStakedAmountPerWallet(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStakedAmountPerWalletRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNStakedAmountPerWallet(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStakedAmountPerWallet(ctx,
			item.Index,
		)
		_, found := keeper.GetStakedAmountPerWallet(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStakedAmountPerWalletGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNStakedAmountPerWallet(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStakedAmountPerWallet(ctx)),
	)
}
