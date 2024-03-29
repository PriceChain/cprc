package keeper_test

import (
	"testing"

	testkeeper "github.com/PriceChain/cprc/testutil/keeper"
	"github.com/PriceChain/cprc/x/prcibc/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PrcibcKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
