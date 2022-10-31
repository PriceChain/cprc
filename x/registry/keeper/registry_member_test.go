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

func createNRegistryMember(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RegistryMember {
	items := make([]types.RegistryMember, n)
	for i := range items {
		items[i].Id = keeper.AppendRegistryMember(ctx, items[i])
	}
	return items
}

func TestRegistryMemberGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryMember(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRegistryMember(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRegistryMemberRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryMember(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRegistryMember(ctx, item.Id)
		_, found := keeper.GetRegistryMember(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRegistryMemberGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryMember(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRegistryMember(ctx)),
	)
}

func TestRegistryMemberCount(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNRegistryMember(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRegistryMemberCount(ctx))
}
