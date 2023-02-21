package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ProposePrice(goCtx context.Context, msg *types.MsgProposePrice) (*types.MsgProposePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// price validation check
	price, err := strconv.ParseFloat(msg.Price, 64)
	// If invalid price
	if price <= 0.0 || err != nil {
		return nil, err
	}

	// registry id validation check
	registryId, err := strconv.ParseUint(msg.RegistryId, 10, 64)
	if err != nil {
		return nil, err
	}

	// registry validation check
	registry, bFound := k.GetRegistry(ctx, registryId)
	if !bFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "There isn't such registry exists!")
	}

	// signer validation check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid signer!")
	}

	bValidProposer := false
	// check if the sender has the rights to propose price
	for _, v := range registry.Validators {
		if v == creator.String() {
			bValidProposer = true
		}
	}

	// if he is not a valid proposer
	if !bValidProposer {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You are not allowed to propose price to this registry!")
	}

	prices := k.GetAllPriceData(ctx)
	index := fmt.Sprintf("%d", len(prices)+1)

	// Construct price data
	priceData := types.PriceData{
		Index:        index,
		RegistryId:   msg.RegistryId,
		Creator:      creator.String(),
		StoreName:    msg.StoreName,
		StoreAddr:    msg.StoreAddr,
		PurchaseTime: msg.PurchaseTime,
		ProdName:     msg.ProdName,
		Price:        msg.Price,
		ReceiptCode:  msg.ReceiptCode,
		Reserved:     msg.Reserved,
	}

	// Set price data to keeper
	k.SetPriceData(ctx, priceData)

	//------------------------------
	//------Update PoP count--------
	//------------------------------
	// Get all registry memebers - price validators
	registryMembers := k.GetAllRegistryMember(ctx)
	// Check if he is already joined as a memeber
	for i, r := range registryMembers {
		// wallet address and registry Id, one user can join several registries
		if r.Wallet == creator.String() && r.RegistryId == msg.RegistryId {
			rm := registryMembers[i]
			popCount, _ := strconv.ParseUint(rm.PopCount, 10, 64)
			level, _ := strconv.ParseUint(rm.Level, 10, 64)

			rm.PopCount = fmt.Sprintf("%d", popCount+1)

			// If he has attended more than 10 PoP, it would starts receiving rewards
			if popCount+1 > types.LEVEL_1_THRESH && level < 1 {
				rm.Level = "1"
			}

			// Update registry member
			k.SetRegistryMember(ctx, rm)
		}
	}

	// Returning
	return &types.MsgProposePriceResponse{}, nil
}
