package keeper

import (
	"beacon/x/beacon/types"
)

var _ types.QueryServer = Keeper{}
