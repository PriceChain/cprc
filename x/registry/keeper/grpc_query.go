package keeper

import (
	"github.com/PriceChain/cprc/x/registry/types"
)

var _ types.QueryServer = Keeper{}
