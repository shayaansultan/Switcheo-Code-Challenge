package documentblockchain_test

import (
	"testing"

	keepertest "documentblockchain/testutil/keeper"
	"documentblockchain/testutil/nullify"
	documentblockchain "documentblockchain/x/documentblockchain/module"
	"documentblockchain/x/documentblockchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DocumentblockchainKeeper(t)
	documentblockchain.InitGenesis(ctx, k, genesisState)
	got := documentblockchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
