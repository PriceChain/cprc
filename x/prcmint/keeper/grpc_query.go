package keeper

import (
	"github.com/PriceChain/cprc/x/prcmint/types"
)

var _ types.QueryServer = Keeper{}
