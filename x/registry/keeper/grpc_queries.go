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

var _ types.QueryServer = Keeper{}

// Fetch Params
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

// Fetcha all price data
func (k Keeper) PriceDataAll(c context.Context, req *types.QueryAllPriceDataRequest) (*types.QueryAllPriceDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var priceDatas []types.PriceData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	priceDataStore := prefix.NewStore(store, types.KeyPrefix(types.PriceDataKeyPrefix))

	pageRes, err := query.Paginate(priceDataStore, req.Pagination, func(key []byte, value []byte) error {
		var priceData types.PriceData
		if err := k.cdc.Unmarshal(value, &priceData); err != nil {
			return err
		}

		priceDatas = append(priceDatas, priceData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPriceDataResponse{PriceData: priceDatas, Pagination: pageRes}, nil
}

// Fetch price data per index
func (k Keeper) PriceData(c context.Context, req *types.QueryGetPriceDataRequest) (*types.QueryGetPriceDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPriceData(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPriceDataResponse{PriceData: val}, nil
}

// Fetch all registry members
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

// Fetch registry member with index
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

// Fetch all registry owners
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

// Fetch registry owner
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

// Fetch all registry staked amount
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

// Fetch registry staked amount per index
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

// Fetch all registries
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

// Fetch a registry with an index
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

// Fetch all staked amount per wallet
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

// Fetch all staked amount per wallet of an index
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
