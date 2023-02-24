package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/PriceChain/cprc/testutil/keeper"
	testkeeper "github.com/PriceChain/cprc/testutil/keeper"
	"github.com/PriceChain/cprc/testutil/nullify"
	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPriceDataQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPriceData(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPriceDataRequest
		response *types.QueryGetPriceDataResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPriceDataRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetPriceDataResponse{PriceData: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPriceDataRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetPriceDataResponse{PriceData: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPriceDataRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PriceData(wctx, tc.request)
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

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

func TestPriceDataQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPriceData(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPriceDataRequest {
		return &types.QueryAllPriceDataRequest{
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
			resp, err := keeper.PriceDataAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PriceData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PriceData),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PriceDataAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PriceData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PriceData),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PriceDataAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PriceData),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PriceDataAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestRegistryMemberQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistryMember(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRegistryMemberRequest
		response *types.QueryGetRegistryMemberResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetRegistryMemberRequest{Id: msgs[0].Id},
			response: &types.QueryGetRegistryMemberResponse{RegistryMember: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetRegistryMemberRequest{Id: msgs[1].Id},
			response: &types.QueryGetRegistryMemberResponse{RegistryMember: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetRegistryMemberRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RegistryMember(wctx, tc.request)
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

func TestRegistryMemberQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistryMember(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRegistryMemberRequest {
		return &types.QueryAllRegistryMemberRequest{
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
			resp, err := keeper.RegistryMemberAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegistryMember), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegistryMember),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RegistryMemberAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegistryMember), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegistryMember),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RegistryMemberAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RegistryMember),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RegistryMemberAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestRegistryOwnerQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistryOwner(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRegistryOwnerRequest
		response *types.QueryGetRegistryOwnerResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetRegistryOwnerRequest{Id: msgs[0].Id},
			response: &types.QueryGetRegistryOwnerResponse{RegistryOwner: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetRegistryOwnerRequest{Id: msgs[1].Id},
			response: &types.QueryGetRegistryOwnerResponse{RegistryOwner: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetRegistryOwnerRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RegistryOwner(wctx, tc.request)
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

func TestRegistryOwnerQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistryOwner(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRegistryOwnerRequest {
		return &types.QueryAllRegistryOwnerRequest{
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
			resp, err := keeper.RegistryOwnerAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegistryOwner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegistryOwner),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RegistryOwnerAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegistryOwner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegistryOwner),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RegistryOwnerAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RegistryOwner),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RegistryOwnerAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestRegistryStakedAmountQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistryStakedAmount(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRegistryStakedAmountRequest
		response *types.QueryGetRegistryStakedAmountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRegistryStakedAmountRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetRegistryStakedAmountResponse{RegistryStakedAmount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRegistryStakedAmountRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetRegistryStakedAmountResponse{RegistryStakedAmount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRegistryStakedAmountRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RegistryStakedAmount(wctx, tc.request)
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

func TestRegistryStakedAmountQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistryStakedAmount(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRegistryStakedAmountRequest {
		return &types.QueryAllRegistryStakedAmountRequest{
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
			resp, err := keeper.RegistryStakedAmountAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegistryStakedAmount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegistryStakedAmount),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RegistryStakedAmountAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegistryStakedAmount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegistryStakedAmount),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RegistryStakedAmountAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RegistryStakedAmount),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RegistryStakedAmountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestStakedAmountPerWalletQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStakedAmountPerWallet(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRegistryStakedAmountPerWalletRequest
		response *types.QueryGetRegistryStakedAmountPerWalletResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRegistryStakedAmountPerWalletRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetRegistryStakedAmountPerWalletResponse{RegistryStakedAmountPerWallet: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRegistryStakedAmountPerWalletRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetRegistryStakedAmountPerWalletResponse{RegistryStakedAmountPerWallet: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRegistryStakedAmountPerWalletRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RegistryStakedAmountPerWallet(wctx, tc.request)
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

func TestStakedAmountPerWalletQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStakedAmountPerWallet(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRegistryStakedAmountPerWalletRequest {
		return &types.QueryAllRegistryStakedAmountPerWalletRequest{
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
			resp, err := keeper.RegistryStakedAmountPerWalletAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegistryStakedAmountPerWallet), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegistryStakedAmountPerWallet),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RegistryStakedAmountPerWalletAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RegistryStakedAmountPerWallet), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RegistryStakedAmountPerWallet),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RegistryStakedAmountPerWalletAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RegistryStakedAmountPerWallet),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RegistryStakedAmountPerWalletAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestRegistryQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistry(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRegistryRequest
		response *types.QueryGetRegistryResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetRegistryRequest{Id: msgs[0].Id},
			response: &types.QueryGetRegistryResponse{Registry: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetRegistryRequest{Id: msgs[1].Id},
			response: &types.QueryGetRegistryResponse{Registry: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetRegistryRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Registry(wctx, tc.request)
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

func TestRegistryQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistry(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRegistryRequest {
		return &types.QueryAllRegistryRequest{
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
			resp, err := keeper.RegistryAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Registry), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Registry),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RegistryAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Registry), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Registry),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RegistryAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Registry),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RegistryAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
