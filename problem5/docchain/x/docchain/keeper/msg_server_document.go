package keeper

import (
	"context"
	"fmt"

	"docchain/x/docchain/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDocument(goCtx context.Context, msg *types.MsgCreateDocument) (*types.MsgCreateDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var document = types.Document{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
		Category: msg.Category,
	}

	id := k.AppendDocument(
		ctx,
		document,
	)

	return &types.MsgCreateDocumentResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateDocument(goCtx context.Context, msg *types.MsgUpdateDocument) (*types.MsgUpdateDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var document = types.Document{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
		Category: msg.Category,
	}

	// Checks that the element exists
	val, found := k.GetDocument(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetDocument(ctx, document)

	return &types.MsgUpdateDocumentResponse{}, nil
}

func (k msgServer) DeleteDocument(goCtx context.Context, msg *types.MsgDeleteDocument) (*types.MsgDeleteDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetDocument(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDocument(ctx, msg.Id)

	return &types.MsgDeleteDocumentResponse{}, nil
}
