package registry

import (
	"math/rand"

	"github.com/PriceChain/rd_net/testutil/sample"
	registrysimulation "github.com/PriceChain/rd_net/x/registry/simulation"
	"github.com/PriceChain/rd_net/x/registry/types"
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
	_ = registrysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgUnbondRegistry = "op_weight_msg_unbond_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnbondRegistry int = 100

	opWeightMsgModifyRegistry = "op_weight_msg_modify_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgModifyRegistry int = 100

	opWeightMsgProposePrice = "op_weight_msg_propose_price"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProposePrice int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	registryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&registryGenesis)
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

	var weightMsgUnbondRegistry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnbondRegistry, &weightMsgUnbondRegistry, nil,
		func(_ *rand.Rand) {
			weightMsgUnbondRegistry = defaultWeightMsgUnbondRegistry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnbondRegistry,
		registrysimulation.SimulateMsgUnbondRegistry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgModifyRegistry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgModifyRegistry, &weightMsgModifyRegistry, nil,
		func(_ *rand.Rand) {
			weightMsgModifyRegistry = defaultWeightMsgModifyRegistry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgModifyRegistry,
		registrysimulation.SimulateMsgModifyRegistry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProposePrice int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposePrice, &weightMsgProposePrice, nil,
		func(_ *rand.Rand) {
			weightMsgProposePrice = defaultWeightMsgProposePrice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposePrice,
		registrysimulation.SimulateMsgProposePrice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
