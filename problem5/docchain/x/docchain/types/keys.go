package types

const (
	// ModuleName defines the module name
	ModuleName = "docchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_docchain"
)

var (
	ParamsKey = []byte("p_docchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DocumentKey      = "Document/value/"
	DocumentCountKey = "Document/count/"
)
