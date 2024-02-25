package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "docchain/testutil/keeper"
	"docchain/testutil/nullify"
	"docchain/x/docchain/types"
)

func TestDocumentQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DocchainKeeper(t)
	msgs := createNDocument(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetDocumentRequest
		response *types.QueryGetDocumentResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetDocumentRequest{Id: msgs[0].Id},
			response: &types.QueryGetDocumentResponse{Document: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetDocumentRequest{Id: msgs[1].Id},
			response: &types.QueryGetDocumentResponse{Document: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetDocumentRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Document(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestDocumentQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DocchainKeeper(t)
	msgs := createNDocument(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDocumentRequest {
		return &types.QueryAllDocumentRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DocumentAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Document), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Document),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DocumentAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Document), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Document),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DocumentAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Document),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DocumentAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
