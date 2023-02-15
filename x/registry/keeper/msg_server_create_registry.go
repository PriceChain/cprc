package keeper

import (
	"context"

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
	amt := collateral.AmountOf(denom)

	// If equals to 0
	if amt.Equal(sdk.ZeroInt()) {
		return &types.MsgCreateRegistryResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid token amount")
	}

	// fetch stored params
	params := k.GetParams(ctx)

	// Parse amount of Minstake
	minStakeCoin, err := sdk.ParseCoinsNormalized(params.MinimumStake)
	if err != nil {
		return &types.MsgCreateRegistryResponse{}, err
	}

	// Get denomination
	denomMinStake := minStakeCoin.GetDenomByIndex(0)

	// Get Int amount
	amtMinStake := minStakeCoin.AmountOf(denomMinStake)

	// if it is below than minimum stake amount
	if denom != denomMinStake || amt.LT(amtMinStake) {
		return &types.MsgCreateRegistryResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "It shouldn't be below than minimum staking amount.")
	}

	// Collect fund from user's wallet to stake
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.RegistryStakeCollectorName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

	var registry = types.Registry{
		Id:           n + 1,
		Name:         msg.Name,
		StakedAmount: amt.String(),
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

	// Update registry count
	k.SetRegistryCount(ctx, count)

	return &types.MsgCreateRegistryResponse{}, nil
}
