package keeper

import (
	"github.com/PriceChain/rd_net/x/registry/types"
)

var _ types.QueryServer = Keeper{}
