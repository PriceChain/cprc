package keeper_test

import (
	"testing"

    "github.com/PriceChain/cprc/x/prcibc/keeper"
    "github.com/PriceChain/cprc/x/prcibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/PriceChain/cprc/testutil/keeper"
	"github.com/PriceChain/cprc/testutil/nullify"
	"github.com/stretchr/testify/require"
)

func createNIbcMsg(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.IbcMsg {
	items := make([]types.IbcMsg, n)
	for i := range items {
		items[i].Id = keeper.AppendIbcMsg(ctx, items[i])
	}
	return items
}

func TestIbcMsgGet(t *testing.T) {
	keeper, ctx := keepertest.PrcibcKeeper(t)
	items := createNIbcMsg(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetIbcMsg(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestIbcMsgRemove(t *testing.T) {
	keeper, ctx := keepertest.PrcibcKeeper(t)
	items := createNIbcMsg(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIbcMsg(ctx, item.Id)
		_, found := keeper.GetIbcMsg(ctx, item.Id)
		require.False(t, found)
	}
}

func TestIbcMsgGetAll(t *testing.T) {
	keeper, ctx := keepertest.PrcibcKeeper(t)
	items := createNIbcMsg(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllIbcMsg(ctx)),
	)
}

func TestIbcMsgCount(t *testing.T) {
	keeper, ctx := keepertest.PrcibcKeeper(t)
	items := createNIbcMsg(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetIbcMsgCount(ctx))
}
