package keeper

import (
	"context"
	"fmt"
	"strconv"

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
	}

	// Append registry data
	k.SetRegistry(ctx, registry)

	// Fetch previous total staked amount
	prevStakedAmount := (uint64)(0)
	stakedAmount, bFound := k.GetRegistryStakedAmount(ctx, "total")
	if bFound {
		amt, _ := strconv.ParseUint(stakedAmount.Amount, 10, 64)
		prevStakedAmount = amt
	}

	// Initalize a new total staked amount item
	rsa := types.RegistryStakedAmount{
		Index:  "total",
		Amount: fmt.Sprintf("%d", prevStakedAmount+amount.Uint64()),
	}

	// Update total staked amount
	k.SetRegistryStakedAmount(ctx, rsa)

	// Update registry count
	k.SetRegistryCount(ctx, n+1)

	return &types.MsgCreateRegistryResponse{}, nil
}
