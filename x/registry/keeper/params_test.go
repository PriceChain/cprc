package keeper_test

import (
	"testing"

	testkeeper "github.com/PriceChain/rd_net/testutil/keeper"
	"github.com/PriceChain/rd_net/x/registry/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RegistryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}