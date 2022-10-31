package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RegistryList: []Registry{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in registry
	registryIdMap := make(map[uint64]bool)
	registryCount := gs.GetRegistryCount()
	for _, elem := range gs.RegistryList {
		if _, ok := registryIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for registry")
		}
		if elem.Id >= registryCount {
			return fmt.Errorf("registry id should be lower or equal than the last id")
		}
		registryIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
