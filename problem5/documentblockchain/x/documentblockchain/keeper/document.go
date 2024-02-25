package keeper

import (
	"context"
	"encoding/binary"

	"documentblockchain/x/documentblockchain/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetDocumentCount get the total number of document
func (k Keeper) GetDocumentCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.DocumentCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetDocumentCount set the total number of document
func (k Keeper) SetDocumentCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.DocumentCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendDocument appends a document in the store with a new id and update the count
func (k Keeper) AppendDocument(
	ctx context.Context,
	document types.Document,
) uint64 {
	// Create the document
	count := k.GetDocumentCount(ctx)

	// Set the ID of the appended value
	document.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DocumentKey))
	appendedValue := k.cdc.MustMarshal(&document)
	store.Set(GetDocumentIDBytes(document.Id), appendedValue)

	// Update document count
	k.SetDocumentCount(ctx, count+1)

	return count
}

// SetDocument set a specific document in the store
func (k Keeper) SetDocument(ctx context.Context, document types.Document) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DocumentKey))
	b := k.cdc.MustMarshal(&document)
	store.Set(GetDocumentIDBytes(document.Id), b)
}

// GetDocument returns a document from its id
func (k Keeper) GetDocument(ctx context.Context, id uint64) (val types.Document, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DocumentKey))
	b := store.Get(GetDocumentIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDocument removes a document from the store
func (k Keeper) RemoveDocument(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DocumentKey))
	store.Delete(GetDocumentIDBytes(id))
}

// GetAllDocument returns all document
func (k Keeper) GetAllDocument(ctx context.Context) (list []types.Document) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DocumentKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Document
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDocumentIDBytes returns the byte representation of the ID
func GetDocumentIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.DocumentKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
