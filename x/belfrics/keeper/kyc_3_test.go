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

func createNKyc3(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Kyc3 {
	items := make([]types.Kyc3, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetKyc3(ctx, items[i])
	}
	return items
}

func TestKyc3Get(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	items := createNKyc3(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetKyc3(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestKyc3Remove(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	items := createNKyc3(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKyc3(ctx,
			item.Address,
		)
		_, found := keeper.GetKyc3(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestKyc3GetAll(t *testing.T) {
	keeper, ctx := keepertest.BelfricsKeeper(t)
	items := createNKyc3(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKyc3(ctx)),
	)
}
