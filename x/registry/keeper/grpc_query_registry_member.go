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

func (k Keeper) RegistryMemberAll(c context.Context, req *types.QueryAllRegistryMemberRequest) (*types.QueryAllRegistryMemberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var registryMembers []types.RegistryMember
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	registryMemberStore := prefix.NewStore(store, types.KeyPrefix(types.RegistryMemberKey))

	pageRes, err := query.Paginate(registryMemberStore, req.Pagination, func(key []byte, value []byte) error {
		var registryMember types.RegistryMember
		if err := k.cdc.Unmarshal(value, &registryMember); err != nil {
			return err
		}

		registryMembers = append(registryMembers, registryMember)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRegistryMemberResponse{RegistryMember: registryMembers, Pagination: pageRes}, nil
}

func (k Keeper) RegistryMember(c context.Context, req *types.QueryGetRegistryMemberRequest) (*types.QueryGetRegistryMemberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	registryMember, found := k.GetRegistryMember(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRegistryMemberResponse{RegistryMember: registryMember}, nil
}
