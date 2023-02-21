package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Join as a member
func (k msgServer) JoinRegistryMember(goCtx context.Context, msg *types.MsgJoinRegistryMember) (*types.MsgJoinRegistryMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get timestamp
	timestamp := ctx.BlockTime().UTC().String()

	// Parse registry Id
	registryId, _ := strconv.ParseUint(msg.RegistryId, 10, 64)

	// Find registry
	registry, bFound := k.GetRegistry(ctx, registryId)

	// There is no registry
	if !bFound {
		return &types.MsgJoinRegistryMemberResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "There is no registry created with the id!")
	}

	// Validation check before process
	creator, amount, collateral, err := k.PreProcess(ctx, msg.Creator, msg.StakeAmount)

	// if invalid input
	if err != nil {
		return &types.MsgJoinRegistryMemberResponse{}, err
	}

	// Check if it is already validator
	for _, validator := range registry.Validators {
		if validator == creator.String() {
			return &types.MsgJoinRegistryMemberResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You are already a validator of this registry!")
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
	registry.Validators = append(registry.Validators, creator.String())
	k.SetRegistry(ctx, registry)

	// Get all registry memebers - price validators
	registryMembers := k.GetAllRegistryMember(ctx)

	// Cosntruct an item of registry memeber
	registryMember := types.RegistryMember{
		Id:           uint64(len(registryMembers) + 1),
		Wallet:       creator.String(),
		RegistryId:   msg.RegistryId,
		StakedAmount: msg.StakeAmount,
		Address:      "",
		Status:       types.STATUS_OPEN,
		PopCount:     "0",
		Level:        "0",
		Reputations:  []string{},
		Scores:       []string{},
		Reserved:     "",
	}

	// Check if he is already joined as a memeber
	for i, r := range registryMembers {
		// wallet address and registry Id, one user can join several registries
		if r.Wallet == creator.String() && r.RegistryId == msg.RegistryId {
			registryMember = registryMembers[i]

			ramt, _ := strconv.ParseUint(registryMember.StakedAmount, 10, 64)
			amt, _ = strconv.ParseUint(amount.String(), 10, 64)

			ramt += amt

			registryMember.StakedAmount = fmt.Sprintf("%d", ramt)
		}
	}

	// Set registry member
	k.SetRegistryMember(ctx, registryMember)

	// Update staked amount per wallet
	k.UpdateStakedAmountPerWallet(ctx, creator.String(), msg.StakeAmount)

	// Update total staked amount
	k.UpdateTotalStakedAmount(ctx, msg.StakeAmount)

	return &types.MsgJoinRegistryMemberResponse{}, nil
}
