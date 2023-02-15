package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/PriceChain/cprc/testutil/keeper"
	"github.com/PriceChain/cprc/testutil/nullify"
	"github.com/PriceChain/cprc/x/prcibc/types"
)

func TestIbcMsgQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PrcibcKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNIbcMsg(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetIbcMsgRequest
		response *types.QueryGetIbcMsgResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetIbcMsgRequest{Id: msgs[0].Id},
			response: &types.QueryGetIbcMsgResponse{IbcMsg: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetIbcMsgRequest{Id: msgs[1].Id},
			response: &types.QueryGetIbcMsgResponse{IbcMsg: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetIbcMsgRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.IbcMsg(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestIbcMsgQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PrcibcKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNIbcMsg(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllIbcMsgRequest {
		return &types.QueryAllIbcMsgRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.IbcMsgAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.IbcMsg), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.IbcMsg),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.IbcMsgAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.IbcMsg), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.IbcMsg),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.IbcMsgAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.IbcMsg),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.IbcMsgAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
