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

func createNKyc2(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Kyc2 {
	items := make([]types.Kyc2, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetKyc2(ctx, items[i])
	}
	return items
}

func TestKyc2Get(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	items := createNKyc2(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetKyc2(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestKyc2Remove(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	items := createNKyc2(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKyc2(ctx,
			item.Address,
		)
		_, found := keeper.GetKyc2(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestKyc2GetAll(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	items := createNKyc2(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKyc2(ctx)),
	)
}
