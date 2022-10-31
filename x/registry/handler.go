package registry

import (
	"fmt"

	"github.com/PriceChain/rd_net/x/registry/keeper"
	"github.com/PriceChain/rd_net/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateRegistry:
			res, err := msgServer.CreateRegistry(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgJoinRegistryCoOperator:
			res, err := msgServer.JoinRegistryCoOperator(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgJoinRegistryMember:
			res, err := msgServer.JoinRegistryMember(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUnbondRegistry:
			res, err := msgServer.UnbondRegistry(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgModifyRegistry:
			res, err := msgServer.ModifyRegistry(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
