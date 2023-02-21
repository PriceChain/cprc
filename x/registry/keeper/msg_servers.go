package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

//------------------------------------------------------------------
//--------------------------Create Registry-------------------------
//------------------------------------------------------------------
// 1. Check min staking amount
// 2. Check if already exists registry, it shouldn't already exists
// 3. Send coins from user wallet to module account
// 4. Increase registry index and set it to keeper
// 5. Update staked amount per wallet
// 6. Updaste total staked amount
//------------------------------------------------------------------
//------------------------------------------------------------------
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

//-------------------------------------------------------------------
//-------------------Join a registry as a co owner-------------------
//-------------------------------------------------------------------
// 1. Check min staking amount
// 2. Check if already exists registry, it should already exists
// 4. Check if he is already in the owners list
// 5. Send coins from user wallet to module account
// 6. Update the registry staked amount, add itself to owners list
// 7. Update staked amount per wallet
// 8. Updaste total staked amount
//-------------------------------------------------------------------
//-------------------------------------------------------------------
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

//-------------------------------------------------------------------------
//---------------------------Join as a member------------------------------
//-------------------------------------------------------------------------
// 1. Check min staking amount
// 2. Check if already exists registry, it should already exists
// 4. Check if he is already in the validators list
// 5. Send coins from user wallet to module account
// 6. Update the registry staked amount, and add itself to validator list
// 7. Add itself to registry member list with zero rewards
// 8. Update staked amount per wallet
// 9. Updaste total staked amount
//--------------------------------------------------------------------------
//--------------------------------------------------------------------------
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
		Id:             uint64(len(registryMembers) + 1),
		Wallet:         creator.String(),
		RegistryId:     msg.RegistryId,
		StakedAmount:   msg.StakeAmount,
		Address:        "",
		Status:         types.STATUS_OPEN,
		PopCount:       "0",
		Level:          "0",
		Reputations:    []string{},
		Scores:         []string{},
		RewardSum:      "0",
		RewardPaid:     "0",
		RewardPaidDate: "",
		Reserved:       "",
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

//-----------------------------------------------------------------
// ---------Modify an existing registry----------------------------
// 1. This should be handled through governance proposal by owners
//-----------------------------------------------------------------
//-----------------------------------------------------------------
func (k msgServer) ModifyRegistry(goCtx context.Context, msg *types.MsgModifyRegistry) (*types.MsgModifyRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgModifyRegistryResponse{}, nil
}

//-------------------------------------------------------------------------------------------------------
// ---------Propose a price data-------------------------------------------------------------------------
// 1. Validation check for input parameters - price, registryId, creator address, valid proposer
// 2. Add a new price data to keeper (on chain)
// 3. Add PoP count of the chosen proposer(validator), level up if he proposed more than 10 price data.
//-------------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------------
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

//-----------------------------------------------------------------
//-----------------------------------------------------------------
// ------------Unbond registry-------------------------------------
// 1. Remove tokens from registry by owners
// 2. Should change registry status according to his staked amount
// 3. There is unbonding period etc.
//-----------------------------------------------------------------
//-----------------------------------------------------------------
//-----------------------------------------------------------------
func (k msgServer) UnbondRegistry(goCtx context.Context, msg *types.MsgUnbondRegistry) (*types.MsgUnbondRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnbondRegistryResponse{}, nil
}
