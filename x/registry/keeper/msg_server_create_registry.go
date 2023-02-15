package keeper

import (
	"context"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateRegistry(goCtx context.Context, msg *types.MsgCreateRegistry) (*types.MsgCreateRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Increase transfer request count
	n := k.GetRegistryCount(ctx)

	// Get timestamp
	timestamp := ctx.BlockTime().UTC().String()

	// Parse creator bec32 address
	registry_owner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgCreateRegistryResponse{}, err
	}

	// Parse amount of CPRC token
	collateral, err := sdk.ParseCoinsNormalized(msg.StakeAmount)
	if err != nil {
		return &types.MsgCreateRegistryResponse{}, err
	}

	//
	amt := collateral.AmountOf("ucprc")

	// If equals to 0
	if amt.Equal(sdk.ZeroInt()) {
		return &types.MsgCreateRegistryResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid token amount")
	}

	// Collect fund from user's wallet to stake
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, registry_owner, types.RegistryStakeCollectorName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

	var registry = types.Registry{
		Id:           n + 1,
		Name:         msg.Name,
		StakedAmount: amt.String(),
		Status:       "Open",
		Description:  msg.Description,
		ImageUrl:     msg.ImageUrl,
		PriceCount:   "0",
		ReviewCount:  "0",
		Timestamp:    timestamp,
		Reserved:     "",
		Creator:      msg.Creator,
	}

	// Append registry data
	count := k.AppendRegistry(ctx, registry)

	// Update registry count
	k.SetRegistryCount(ctx, count)

	return &types.MsgCreateRegistryResponse{}, nil
}
