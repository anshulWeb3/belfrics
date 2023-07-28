package belfrics_test

import (
	"testing"

	keepertest "belfrics/testutil/keeper"
	"belfrics/testutil/nullify"
	"belfrics/x/belfrics"
	"belfrics/x/belfrics/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		KycList: []types.Kyc{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		Kyc2List: []types.Kyc2{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BelfricsKeeper(t)
	belfrics.InitGenesis(ctx, *k, genesisState)
	got := belfrics.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.KycList, got.KycList)
	require.ElementsMatch(t, genesisState.Kyc2List, got.Kyc2List)
	// this line is used by starport scaffolding # genesis/test/assert
}
