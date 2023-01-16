package keeper

import (
	"context"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ModifyRegistry(goCtx context.Context, msg *types.MsgModifyRegistry) (*types.MsgModifyRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgModifyRegistryResponse{}, nil
}
