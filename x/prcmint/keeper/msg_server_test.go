package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/PriceChain/cprc/testutil/keeper"
	"github.com/PriceChain/cprc/x/prcmint/keeper"
	"github.com/PriceChain/cprc/x/prcmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PrcmintKeeper(t)
	return keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
}
