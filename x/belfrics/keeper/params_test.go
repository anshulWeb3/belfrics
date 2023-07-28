package keeper_test

import (
	"testing"

	testkeeper "belfrics/testutil/keeper"
	"belfrics/x/belfrics/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BelfricsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
