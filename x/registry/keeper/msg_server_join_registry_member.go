package keeper

import (
	"context"

	"github.com/PriceChain/rd_net/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) JoinRegistryMember(goCtx context.Context, msg *types.MsgJoinRegistryMember) (*types.MsgJoinRegistryMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgJoinRegistryMemberResponse{}, nil
}
