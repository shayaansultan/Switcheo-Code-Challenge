package keeper_test

import (
	"context"
	"testing"

	keepertest "docchain/testutil/keeper"
	"docchain/testutil/nullify"
	"docchain/x/docchain/keeper"
	"docchain/x/docchain/types"

	"github.com/stretchr/testify/require"
)

func createNDocument(keeper keeper.Keeper, ctx context.Context, n int) []types.Document {
	items := make([]types.Document, n)
	for i := range items {
		items[i].Id = keeper.AppendDocument(ctx, items[i])
	}
	return items
}

func TestDocumentGet(t *testing.T) {
	keeper, ctx := keepertest.DocchainKeeper(t)
	items := createNDocument(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetDocument(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestDocumentRemove(t *testing.T) {
	keeper, ctx := keepertest.DocchainKeeper(t)
	items := createNDocument(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDocument(ctx, item.Id)
		_, found := keeper.GetDocument(ctx, item.Id)
		require.False(t, found)
	}
}

func TestDocumentGetAll(t *testing.T) {
	keeper, ctx := keepertest.DocchainKeeper(t)
	items := createNDocument(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDocument(ctx)),
	)
}

func TestDocumentCount(t *testing.T) {
	keeper, ctx := keepertest.DocchainKeeper(t)
	items := createNDocument(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetDocumentCount(ctx))
}
