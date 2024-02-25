package types

const (
	// ModuleName defines the module name
	ModuleName = "documentblockchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_documentblockchain"
)

var (
	ParamsKey = []byte("p_documentblockchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DocumentKey      = "Document/value/"
	DocumentCountKey = "Document/count/"
)
