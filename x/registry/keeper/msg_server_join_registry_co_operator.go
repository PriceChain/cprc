package keeper

import (
	"context"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) JoinRegistryCoOperator(goCtx context.Context, msg *types.MsgJoinRegistryCoOperator) (*types.MsgJoinRegistryCoOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgJoinRegistryCoOperatorResponse{}, nil
}
