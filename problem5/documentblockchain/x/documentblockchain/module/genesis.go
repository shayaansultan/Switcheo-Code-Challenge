package documentblockchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"documentblockchain/x/documentblockchain/keeper"
	"documentblockchain/x/documentblockchain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the document
	for _, elem := range genState.DocumentList {
		k.SetDocument(ctx, elem)
	}

	// Set document count
	k.SetDocumentCount(ctx, genState.DocumentCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DocumentList = k.GetAllDocument(ctx)
	genesis.DocumentCount = k.GetDocumentCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
