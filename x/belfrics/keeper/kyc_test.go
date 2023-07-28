package keeper_test

import (
	"strconv"
	"testing"

	keepertest "belfrics/testutil/keeper"
	"belfrics/testutil/nullify"
	"belfrics/x/belfrics/keeper"
	"belfrics/x/belfrics/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNKyc(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Kyc {
	items := make([]types.Kyc, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetKyc(ctx, items[i])
	}
	return items
}

func TestKycGet(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	items := createNKyc(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetKyc(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestKycRemove(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	items := createNKyc(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKyc(ctx,
			item.Index,
		)
		_, found := keeper.GetKyc(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestKycGetAll(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	items := createNKyc(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKyc(ctx)),
	)
}
