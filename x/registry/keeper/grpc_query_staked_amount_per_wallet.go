package keeper

import (
	"context"

	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakedAmountPerWalletAll(c context.Context, req *types.QueryAllStakedAmountPerWalletRequest) (*types.QueryAllStakedAmountPerWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stakedAmountPerWallets []types.StakedAmountPerWallet
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	stakedAmountPerWalletStore := prefix.NewStore(store, types.KeyPrefix(types.StakedAmountPerWalletKeyPrefix))

	pageRes, err := query.Paginate(stakedAmountPerWalletStore, req.Pagination, func(key []byte, value []byte) error {
		var stakedAmountPerWallet types.StakedAmountPerWallet
		if err := k.cdc.Unmarshal(value, &stakedAmountPerWallet); err != nil {
			return err
		}

		stakedAmountPerWallets = append(stakedAmountPerWallets, stakedAmountPerWallet)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStakedAmountPerWalletResponse{StakedAmountPerWallet: stakedAmountPerWallets, Pagination: pageRes}, nil
}

func (k Keeper) StakedAmountPerWallet(c context.Context, req *types.QueryGetStakedAmountPerWalletRequest) (*types.QueryGetStakedAmountPerWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStakedAmountPerWallet(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStakedAmountPerWalletResponse{StakedAmountPerWallet: val}, nil
}
