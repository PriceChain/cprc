package keeper

import (
	"context"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateRegistry(goCtx context.Context, msg *types.MsgCreateRegistry) (*types.MsgCreateRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get timestamp
	timestamp := ctx.BlockTime().UTC().String()

	// Validation check before process
	creator, amount, collateral, err := k.PreProcess(ctx, msg.Creator, msg.StakeAmount)

	// if invalid input
	if err != nil {
		return &types.MsgCreateRegistryResponse{}, err
	}

	// Check if it is already registered registry name
	allRegistries := k.GetAllRegistry(ctx)
	for _, r := range allRegistries {
		if r.Name == msg.Name {
			return &types.MsgCreateRegistryResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "There is already registry created with the same name!")
		}
	}

	// Collect fund from user's wallet to stake
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.RegistryStakeCollectorName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

	// Get registry count
	n := k.GetRegistryCount(ctx)

	// Create an registry item
	registry := types.Registry{
		Id:           n + 1,
		Name:         msg.Name,
		StakedAmount: amount.String(),
		Status:       types.STATUS_OPEN,
		Description:  msg.Description,
		ImageUrl:     msg.ImageUrl,
		PriceCount:   "0",
		ReviewCount:  "0",
		Timestamp:    timestamp,
		Reserved:     "",
		Owners:       []string{creator.String()},
		Validators:   []string{},
	}

	// Append registry data
	k.SetRegistry(ctx, registry)

	// Update registry count
	k.SetRegistryCount(ctx, n+1)

	// Update staked amount per wallet
	k.UpdateStakedAmountPerWallet(ctx, creator.String(), msg.StakeAmount)

	// Update total staked amount
	k.UpdateTotalStakedAmount(ctx, msg.StakeAmount)

	// Return result
	return &types.MsgCreateRegistryResponse{}, nil
}
