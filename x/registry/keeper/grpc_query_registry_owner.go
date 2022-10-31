package keeper

import (
	"context"

	"github.com/PriceChain/rd_net/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RegistryOwnerAll(c context.Context, req *types.QueryAllRegistryOwnerRequest) (*types.QueryAllRegistryOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var registryOwners []types.RegistryOwner
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	registryOwnerStore := prefix.NewStore(store, types.KeyPrefix(types.RegistryOwnerKey))

	pageRes, err := query.Paginate(registryOwnerStore, req.Pagination, func(key []byte, value []byte) error {
		var registryOwner types.RegistryOwner
		if err := k.cdc.Unmarshal(value, &registryOwner); err != nil {
			return err
		}

		registryOwners = append(registryOwners, registryOwner)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRegistryOwnerResponse{RegistryOwner: registryOwners, Pagination: pageRes}, nil
}

func (k Keeper) RegistryOwner(c context.Context, req *types.QueryGetRegistryOwnerRequest) (*types.QueryGetRegistryOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	registryOwner, found := k.GetRegistryOwner(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRegistryOwnerResponse{RegistryOwner: registryOwner}, nil
}
