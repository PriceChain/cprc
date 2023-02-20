package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) JoinRegistryCoOperator(goCtx context.Context, msg *types.MsgJoinRegistryCoOperator) (*types.MsgJoinRegistryCoOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get timestamp
	timestamp := ctx.BlockTime().UTC().String()

	// Parse registry Id
	registryId, _ := strconv.ParseUint(msg.RegistryId, 10, 64)

	// Find registry
	registry, bFound := k.GetRegistry(ctx, registryId)

	// There is no registry
	if !bFound {
		return &types.MsgJoinRegistryCoOperatorResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "There is no registry created with the id!")
	}

	// Validation check before process
	creator, amount, collateral, err := k.PreProcess(ctx, msg.Creator, msg.StakeAmount)

	// if invalid input
	if err != nil {
		return &types.MsgJoinRegistryCoOperatorResponse{}, err
	}

	// Check if it is already owner
	for _, owner := range registry.Owners {
		if owner == creator.String() {
			return &types.MsgJoinRegistryCoOperatorResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You are already an owner of this registry!")
		}
	}

	// Collect fund from user's wallet to stake
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.RegistryStakeCollectorName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

	amt, _ := strconv.ParseUint(amount.String(), 10, 64)
	stakedAmt, _ := strconv.ParseUint(registry.StakedAmount, 10, 64)
	amt += stakedAmt

	// Update registry
	registry.StakedAmount = fmt.Sprintf("%d", amt)
	registry.Timestamp = timestamp
	registry.Owners = append(registry.Owners, creator.String())
	k.SetRegistry(ctx, registry)

	// Update staked amount per wallet
	k.UpdateStakedAmountPerWallet(ctx, creator.String(), msg.StakeAmount)

	// Update total staked amount
	k.UpdateTotalStakedAmount(ctx, msg.StakeAmount)

	return &types.MsgJoinRegistryCoOperatorResponse{}, nil
}
