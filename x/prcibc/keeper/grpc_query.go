package keeper

import (
	"github.com/PriceChain/cprc/x/prcibc/types"
)

var _ types.QueryServer = Keeper{}
