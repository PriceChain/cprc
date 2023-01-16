package keeper

import (
	"context"

	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RegistryAll(c context.Context, req *types.QueryAllRegistryRequest) (*types.QueryAllRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var registrys []types.Registry
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	registryStore := prefix.NewStore(store, types.KeyPrefix(types.RegistryKey))

	pageRes, err := query.Paginate(registryStore, req.Pagination, func(key []byte, value []byte) error {
		var registry types.Registry
		if err := k.cdc.Unmarshal(value, &registry); err != nil {
			return err
		}

		registrys = append(registrys, registry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRegistryResponse{Registry: registrys, Pagination: pageRes}, nil
}

func (k Keeper) Registry(c context.Context, req *types.QueryGetRegistryRequest) (*types.QueryGetRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	registry, found := k.GetRegistry(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRegistryResponse{Registry: registry}, nil
}
