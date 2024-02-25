package docchain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"docchain/testutil/sample"
	docchainsimulation "docchain/x/docchain/simulation"
	"docchain/x/docchain/types"
)

// avoid unused import issue
var (
	_ = docchainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateDocument = "op_weight_msg_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDocument int = 100

	opWeightMsgUpdateDocument = "op_weight_msg_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDocument int = 100

	opWeightMsgDeleteDocument = "op_weight_msg_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDocument int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	docchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DocumentList: []types.Document{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		DocumentCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&docchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDocument int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateDocument, &weightMsgCreateDocument, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDocument = defaultWeightMsgCreateDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDocument,
		docchainsimulation.SimulateMsgCreateDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDocument int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateDocument, &weightMsgUpdateDocument, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDocument = defaultWeightMsgUpdateDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDocument,
		docchainsimulation.SimulateMsgUpdateDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDocument int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteDocument, &weightMsgDeleteDocument, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDocument = defaultWeightMsgDeleteDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDocument,
		docchainsimulation.SimulateMsgDeleteDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateDocument,
			defaultWeightMsgCreateDocument,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				docchainsimulation.SimulateMsgCreateDocument(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateDocument,
			defaultWeightMsgUpdateDocument,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				docchainsimulation.SimulateMsgUpdateDocument(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteDocument,
			defaultWeightMsgDeleteDocument,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				docchainsimulation.SimulateMsgDeleteDocument(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
