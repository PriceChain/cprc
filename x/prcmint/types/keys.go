package types

// MinterKey is the key to use for the keeper store.
var MinterKey = []byte{0x00}

const (
	// ModuleName defines the module name
	ModuleName = "prcmint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_prcmint"

	// Query endpoints supported by the minting querier
	QueryParameters       = "parameters"
	QueryInflation        = "inflation"
	QueryAnnualProvisions = "annual_provisions"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
