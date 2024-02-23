package keeper

import (
	"documentblockchain/x/documentblockchain/types"
)

var _ types.QueryServer = Keeper{}
