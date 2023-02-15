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

func (k Keeper) RegistryStakedAmountAll(c context.Context, req *types.QueryAllRegistryStakedAmountRequest) (*types.QueryAllRegistryStakedAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var registryStakedAmounts []types.RegistryStakedAmount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	registryStakedAmountStore := prefix.NewStore(store, types.KeyPrefix(types.RegistryStakedAmountKeyPrefix))

	pageRes, err := query.Paginate(registryStakedAmountStore, req.Pagination, func(key []byte, value []byte) error {
		var registryStakedAmount types.RegistryStakedAmount
		if err := k.cdc.Unmarshal(value, &registryStakedAmount); err != nil {
			return err
		}

		registryStakedAmounts = append(registryStakedAmounts, registryStakedAmount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRegistryStakedAmountResponse{RegistryStakedAmount: registryStakedAmounts, Pagination: pageRes}, nil
}

func (k Keeper) RegistryStakedAmount(c context.Context, req *types.QueryGetRegistryStakedAmountRequest) (*types.QueryGetRegistryStakedAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRegistryStakedAmount(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRegistryStakedAmountResponse{RegistryStakedAmount: val}, nil
}
