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

		DocumentList: []types.Document{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		DocumentCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DocumentblockchainKeeper(t)
	documentblockchain.InitGenesis(ctx, k, genesisState)
	got := documentblockchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DocumentList, got.DocumentList)
	require.Equal(t, genesisState.DocumentCount, got.DocumentCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
