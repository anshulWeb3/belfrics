package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "belfrics/testutil/keeper"
	"belfrics/x/belfrics/keeper"
	"belfrics/x/belfrics/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestKyc3MsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BelfricsKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateKyc3{Creator: creator,
			Address: strconv.Itoa(i),
		}
		_, err := srv.CreateKyc3(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetKyc3(ctx,
			expected.Address,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestKyc3MsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateKyc3
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateKyc3{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateKyc3{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateKyc3{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BelfricsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateKyc3{Creator: creator,
				Address: strconv.Itoa(0),
			}
			_, err := srv.CreateKyc3(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateKyc3(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetKyc3(ctx,
					expected.Address,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestKyc3MsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteKyc3
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteKyc3{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteKyc3{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteKyc3{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BelfricsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateKyc3(wctx, &types.MsgCreateKyc3{Creator: creator,
				Address: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteKyc3(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetKyc3(ctx,
					tc.request.Address,
				)
				require.False(t, found)
			}
		})
	}
}
