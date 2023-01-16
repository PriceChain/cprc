package keeper

import (
	"context"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UnbondRegistry(goCtx context.Context, msg *types.MsgUnbondRegistry) (*types.MsgUnbondRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnbondRegistryResponse{}, nil
}
