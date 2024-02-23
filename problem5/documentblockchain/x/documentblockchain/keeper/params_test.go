package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "documentblockchain/testutil/keeper"
	"documentblockchain/x/documentblockchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DocumentblockchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
