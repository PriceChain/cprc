package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RegistryList:       []Registry{},
		RegistryOwnerList:  []RegistryOwner{},
		RegistryMemberList: []RegistryMember{},
		PriceConsensusList: []PriceConsensus{},
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
	// Check for duplicated ID in registryOwner
	registryOwnerIdMap := make(map[uint64]bool)
	registryOwnerCount := gs.GetRegistryOwnerCount()
	for _, elem := range gs.RegistryOwnerList {
		if _, ok := registryOwnerIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for registryOwner")
		}
		if elem.Id >= registryOwnerCount {
			return fmt.Errorf("registryOwner id should be lower or equal than the last id")
		}
		registryOwnerIdMap[elem.Id] = true
	}
	// Check for duplicated ID in registryMember
	registryMemberIdMap := make(map[uint64]bool)
	registryMemberCount := gs.GetRegistryMemberCount()
	for _, elem := range gs.RegistryMemberList {
		if _, ok := registryMemberIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for registryMember")
		}
		if elem.Id >= registryMemberCount {
			return fmt.Errorf("registryMember id should be lower or equal than the last id")
		}
		registryMemberIdMap[elem.Id] = true
	}
	// Check for duplicated ID in priceConsensus
	priceConsensusIdMap := make(map[uint64]bool)
	priceConsensusCount := gs.GetPriceConsensusCount()
	for _, elem := range gs.PriceConsensusList {
		if _, ok := priceConsensusIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for priceConsensus")
		}
		if elem.Id >= priceConsensusCount {
			return fmt.Errorf("priceConsensus id should be lower or equal than the last id")
		}
		priceConsensusIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
