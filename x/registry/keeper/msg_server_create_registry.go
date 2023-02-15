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

	// Get registry count
	n := k.GetRegistryCount(ctx)

	// Get timestamp
	timestamp := ctx.BlockTime().UTC().String()

	// Parse creator bech32 address
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgCreateRegistryResponse{}, err
	}

	// Parse amount of CPRC token
	collateral, err := sdk.ParseCoinsNormalized(msg.StakeAmount)
	if err != nil {
		return &types.MsgCreateRegistryResponse{}, err
	}

	// Get denomination
	denom := collateral.GetDenomByIndex(0)

	// Get Int amount
	amount := collateral.AmountOf(denom)

	// If equals to 0
	if amount.Equal(sdk.ZeroInt()) {
		return &types.MsgCreateRegistryResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid token amount")
	}

	// fetch stored params
	params := k.GetParams(ctx)

	// Parse amount of Minstake
	minStakeCoin, err := sdk.ParseCoinsNormalized(params.MinimumStakeAmount)
	if err != nil {
		return &types.MsgCreateRegistryResponse{}, err
	}

	// Get denomination
	denomMinStake := minStakeCoin.GetDenomByIndex(0)

	// Get Int amount
	amtMinStake := minStakeCoin.AmountOf(denomMinStake)

	// if it is below than minimum stake amount
	if denom != denomMinStake || amount.LT(amtMinStake) {
		return &types.MsgCreateRegistryResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "It shouldn't be below than minimum staking amount.")
	}

	// Check if it is already registered registry name
	allRegistries := k.GetAllRegistry(ctx)
	for _, r := range allRegistries {
		if r.Name == msg.Name {
			return &types.MsgCreateRegistryResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "It shouldn't be below than minimum staking amount.")
		}
	}

	// Collect fund from user's wallet to stake
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.RegistryStakeCollectorName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

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
		Creator:      msg.Creator,
	}

	// Append registry data
	count := k.AppendRegistry(ctx, registry)

	// Fetch previous total staked amount
	prevStakedAmount := (uint64)(0)
	stakedAmount, bFound := k.GetRegistryStakedAmount(ctx, "total")
	if bFound {
		amt, _ := strconv.ParseUint(stakedAmount.Amount, 10, 64)
		prevStakedAmount = amt
	}
	
	// Initalize a new total staked amount item
	rsa := types.RegistryStakedAmount {
		Index: "total",
		Amount: fmt.Sprintf("%d", prevStakedAmount + amount.Uint64()),
	}

	// Update total staked amount
	k.SetRegistryStakedAmount(ctx, rsa)

	// Update registry count
	k.SetRegistryCount(ctx, count)

	return &types.MsgCreateRegistryResponse{}, nil
}
