package belfrics

import (
	"belfrics/x/belfrics/keeper"
	"belfrics/x/belfrics/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the kyc
	for _, elem := range genState.KycList {
		k.SetKyc(ctx, elem)
	}
	// Set all the kyc2
	for _, elem := range genState.Kyc2List {
		k.SetKyc2(ctx, elem)
	}
	// Set all the kyc3
	for _, elem := range genState.Kyc3List {
		k.SetKyc3(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.KycList = k.GetAllKyc(ctx)
	genesis.Kyc2List = k.GetAllKyc2(ctx)
	genesis.Kyc3List = k.GetAllKyc3(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
