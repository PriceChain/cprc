package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/PriceChain/cprc/x/prcibc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IbcMsgAll(c context.Context, req *types.QueryAllIbcMsgRequest) (*types.QueryAllIbcMsgResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ibcMsgs []types.IbcMsg
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ibcMsgStore := prefix.NewStore(store, types.KeyPrefix(types.IbcMsgKey))

	pageRes, err := query.Paginate(ibcMsgStore, req.Pagination, func(key []byte, value []byte) error {
		var ibcMsg types.IbcMsg
		if err := k.cdc.Unmarshal(value, &ibcMsg); err != nil {
			return err
		}

		ibcMsgs = append(ibcMsgs, ibcMsg)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIbcMsgResponse{IbcMsg: ibcMsgs, Pagination: pageRes}, nil
}

func (k Keeper) IbcMsg(c context.Context, req *types.QueryGetIbcMsgRequest) (*types.QueryGetIbcMsgResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	ibcMsg, found := k.GetIbcMsg(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetIbcMsgResponse{IbcMsg: ibcMsg}, nil
}
