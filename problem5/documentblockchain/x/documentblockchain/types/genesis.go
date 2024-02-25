package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DocumentList: []Document{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in document
	documentIdMap := make(map[uint64]bool)
	documentCount := gs.GetDocumentCount()
	for _, elem := range gs.DocumentList {
		if _, ok := documentIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for document")
		}
		if elem.Id >= documentCount {
			return fmt.Errorf("document id should be lower or equal than the last id")
		}
		documentIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
