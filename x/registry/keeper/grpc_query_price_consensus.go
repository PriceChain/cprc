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

func (k Keeper) PriceConsensusAll(c context.Context, req *types.QueryAllPriceConsensusRequest) (*types.QueryAllPriceConsensusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var priceConsensuss []types.PriceConsensus
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	priceConsensusStore := prefix.NewStore(store, types.KeyPrefix(types.PriceConsensusKey))

	pageRes, err := query.Paginate(priceConsensusStore, req.Pagination, func(key []byte, value []byte) error {
		var priceConsensus types.PriceConsensus
		if err := k.cdc.Unmarshal(value, &priceConsensus); err != nil {
			return err
		}

		priceConsensuss = append(priceConsensuss, priceConsensus)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPriceConsensusResponse{PriceConsensus: priceConsensuss, Pagination: pageRes}, nil
}

func (k Keeper) PriceConsensus(c context.Context, req *types.QueryGetPriceConsensusRequest) (*types.QueryGetPriceConsensusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	priceConsensus, found := k.GetPriceConsensus(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPriceConsensusResponse{PriceConsensus: priceConsensus}, nil
}
