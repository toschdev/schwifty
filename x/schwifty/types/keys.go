package types

const (
	// ModuleName defines the module name
	ModuleName = "schwifty"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_schwifty"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	CollectionKey      = "Collection/value/"
	CollectionCountKey = "Collection/count/"
)

const (
	NftKey      = "Nft/value/"
	NftCountKey = "Nft/count/"
)
