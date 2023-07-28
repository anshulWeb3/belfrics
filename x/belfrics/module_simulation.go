package belfrics

import (
	"math/rand"

	"belfrics/testutil/sample"
	belfricssimulation "belfrics/x/belfrics/simulation"
	"belfrics/x/belfrics/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = belfricssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateKyc2 = "op_weight_msg_kyc_2"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKyc2 int = 100

	opWeightMsgUpdateKyc2 = "op_weight_msg_kyc_2"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKyc2 int = 100

	opWeightMsgDeleteKyc2 = "op_weight_msg_kyc_2"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKyc2 int = 100

	opWeightMsgCreateKyc3 = "op_weight_msg_kyc_3"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKyc3 int = 100

	opWeightMsgUpdateKyc3 = "op_weight_msg_kyc_3"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKyc3 int = 100

	opWeightMsgDeleteKyc3 = "op_weight_msg_kyc_3"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKyc3 int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	belfricsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		Kyc2List: []types.Kyc2{
			{
				Creator: sample.AccAddress(),
				Address: "0",
			},
			{
				Creator: sample.AccAddress(),
				Address: "1",
			},
		},
		Kyc3List: []types.Kyc3{
			{
				Creator: sample.AccAddress(),
				Address: "0",
			},
			{
				Creator: sample.AccAddress(),
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&belfricsGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateKyc2 int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKyc2, &weightMsgCreateKyc2, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKyc2 = defaultWeightMsgCreateKyc2
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKyc2,
		belfricssimulation.SimulateMsgCreateKyc2(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateKyc2 int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateKyc2, &weightMsgUpdateKyc2, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateKyc2 = defaultWeightMsgUpdateKyc2
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateKyc2,
		belfricssimulation.SimulateMsgUpdateKyc2(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteKyc2 int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteKyc2, &weightMsgDeleteKyc2, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteKyc2 = defaultWeightMsgDeleteKyc2
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteKyc2,
		belfricssimulation.SimulateMsgDeleteKyc2(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateKyc3 int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKyc3, &weightMsgCreateKyc3, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKyc3 = defaultWeightMsgCreateKyc3
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKyc3,
		belfricssimulation.SimulateMsgCreateKyc3(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateKyc3 int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateKyc3, &weightMsgUpdateKyc3, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateKyc3 = defaultWeightMsgUpdateKyc3
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateKyc3,
		belfricssimulation.SimulateMsgUpdateKyc3(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteKyc3 int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteKyc3, &weightMsgDeleteKyc3, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteKyc3 = defaultWeightMsgDeleteKyc3
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteKyc3,
		belfricssimulation.SimulateMsgDeleteKyc3(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
