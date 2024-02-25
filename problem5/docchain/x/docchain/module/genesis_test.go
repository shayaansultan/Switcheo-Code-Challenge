package docchain_test

import (
	"testing"

	keepertest "docchain/testutil/keeper"
	"docchain/testutil/nullify"
	docchain "docchain/x/docchain/module"
	"docchain/x/docchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DocchainKeeper(t)
	docchain.InitGenesis(ctx, k, genesisState)
	got := docchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
