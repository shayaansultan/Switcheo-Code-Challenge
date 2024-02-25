package keeper

import (
	"docchain/x/docchain/types"
)

var _ types.QueryServer = Keeper{}
