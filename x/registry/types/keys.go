package types

const (
	// ModuleName defines the module name
	ModuleName = "registry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_registry"

	// RegistryStakeCollector the root string for the registry stakers account address
	RegistryStakeCollectorName = "registry_stake"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	RegistryKey      = "Registry-value-"
	RegistryCountKey = "Registry-count-"
)

const (
	RegistryOwnerKey      = "RegistryOwner-value-"
	RegistryOwnerCountKey = "RegistryOwner-count-"
)

const (
	RegistryMemberKey      = "RegistryMember-value-"
	RegistryMemberCountKey = "RegistryMember-count-"
)

const (
	PriceConsensusKey      = "PriceConsensus-value-"
	PriceConsensusCountKey = "PriceConsensus-count-"
)
