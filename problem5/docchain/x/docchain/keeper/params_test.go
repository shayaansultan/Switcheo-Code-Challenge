package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "docchain/testutil/keeper"
	"docchain/x/docchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DocchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
