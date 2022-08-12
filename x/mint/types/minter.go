package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMinter returns a new Minter object with the given inflation and annual
// provisions values.
func NewMinter(inflation, annualProvisions sdk.Int, phase, startPhaseBlock uint64) Minter {
	return Minter{
		Inflation:        inflation,
		AnnualProvisions: annualProvisions,
		Phase:            phase,
		StartPhaseBlock:  startPhaseBlock,
	}
}

// InitialMinter returns an initial Minter object with a given inflation value.
func InitialMinter(inflation sdk.Int) Minter {
	return NewMinter(
		inflation,
		sdk.NewInt(0),
		0,
		0,
	)
}

// DefaultInitialMinter returns a default initial Minter object for a new chain
// which uses an inflation rate of 13%.
func DefaultInitialMinter() Minter {
	return InitialMinter(
		sdk.NewInt(0),
	)
}

// validate minter
func ValidateMinter(minter Minter) error {
	if minter.Inflation.IsNegative() {
		return fmt.Errorf("mint parameter Inflation should be positive, is %s",
			minter.Inflation.String())
	}
	return nil
}

// InflationcalculationFn returns the creators additional coin amounts by phase.
func (m Minter) InflationcalculationFn(phase uint64) sdk.Int {
	InflationAmt := sdk.NewInt(Inflation_Amount_Per_Year)

	if phase > 30 {
		InflationAmt = sdk.NewInt(0)
	}

	//-------------- Phase 2--------------
	// switch {
	// case phase > 5:
	// 	return InflationAmt

	// case phase == 1:
	// 	return InflationAmt.Add(sdk.NewInt(13333333))

	// case phase == 2:
	// 	return InflationAmt.Add(sdk.NewInt(4870130))

	// case phase == 3:
	// 	return InflationAmt.Add(sdk.NewInt(3893871))

	// case phase == 4:
	// 	return InflationAmt.Add(sdk.NewInt(1971160))

	// case phase == 5:
	// 	return InflationAmt.Add(sdk.NewInt(1742798))
	// default:
	// 	return sdk.NewInt(0)
	// }

	return InflationAmt
}

// NextPhase returns the new phase.
func (m Minter) NextPhase(params Params, currentBlock uint64) uint64 {
	nonePhase := m.Phase == 0
	if nonePhase {
		return 1
	}

	blockNewPhase := m.StartPhaseBlock + params.BlocksPerYear
	if blockNewPhase > currentBlock {
		return m.Phase
	}

	return m.Phase + 1
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(params Params) sdk.Coin {
	provisionAmt := m.AnnualProvisions.Quo(sdk.NewInt(int64(params.BlocksPerYear)))
	return sdk.NewCoin(params.MintDenom, sdk.NewInt(provisionAmt.Int64()))
}
