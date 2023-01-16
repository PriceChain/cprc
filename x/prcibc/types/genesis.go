package types

import (
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
		IbcMsgList: []IbcMsg{},
// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in ibcMsg
ibcMsgIdMap := make(map[uint64]bool)
ibcMsgCount := gs.GetIbcMsgCount()
for _, elem := range gs.IbcMsgList {
	if _, ok := ibcMsgIdMap[elem.Id]; ok {
		return fmt.Errorf("duplicated id for ibcMsg")
	}
	if elem.Id >= ibcMsgCount {
		return fmt.Errorf("ibcMsg id should be lower or equal than the last id")
	}
	ibcMsgIdMap[elem.Id] = true
}
// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
