package keeper

import (
	"context"

	"github.com/PriceChain/rd_net/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProposePrice(goCtx context.Context, msg *types.MsgProposePrice) (*types.MsgProposePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProposePriceResponse{}, nil
}
