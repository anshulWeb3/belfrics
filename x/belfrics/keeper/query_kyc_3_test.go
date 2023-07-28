package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "belfrics/testutil/keeper"
	"belfrics/testutil/nullify"
	"belfrics/x/belfrics/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestKyc3QuerySingle(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNKyc3(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetKyc3Request
		response *types.QueryGetKyc3Response
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetKyc3Request{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetKyc3Response{Kyc3: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetKyc3Request{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetKyc3Response{Kyc3: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetKyc3Request{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Kyc3(wctx, tc.request)
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

func TestKyc3QueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNKyc3(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllKyc3Request {
		return &types.QueryAllKyc3Request{
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
			resp, err := keeper.Kyc3All(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Kyc3), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Kyc3),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.Kyc3All(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Kyc3), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Kyc3),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.Kyc3All(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Kyc3),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.Kyc3All(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
