package keeper

import (
	"context"
	"strings"

	"docchain/x/docchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SearchTitle(goCtx context.Context, req *types.QuerySearchTitleRequest) (*types.QuerySearchTitleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	var documents []types.Document
	allDocuments := k.GetAllDocument(ctx)
	for _, document := range allDocuments {
		if strings.Contains(strings.ToLower(document.Title), strings.ToLower(req.Word)) {
			documents = append(documents, document)
		}
	}

	// Convert documents of type []types.Document to type []*types.Document
	// This is necessary because the QuerySearchTitleResponse struct has a field of type []*types.Document
	// and the documents variable is of type []types.Document
	var documentsPtr []*types.Document
	for i := range documents {
		documentsPtr = append(documentsPtr, &documents[i])
	}

	return &types.QuerySearchTitleResponse{Document: documentsPtr}, nil
}
