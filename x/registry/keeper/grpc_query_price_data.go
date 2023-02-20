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
