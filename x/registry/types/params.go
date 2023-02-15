package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// Parameter store keys
var (
	KeyMinimumStake = []byte("MinimumStake")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(minimumStake string) Params {
	return Params{MinimumStake: minimumStake}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		MinimumStake: "100000000ucprc",
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinimumStake, &p.MinimumStake, validateMinimumStake),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMinimumStake(p.MinimumStake); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateMinimumStake(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// Parse amount of CPRC token
	collateral, err := sdk.ParseCoinsNormalized(v)
	if err != nil {
		return fmt.Errorf("invalid token: %s", v)
	}

	// Get coin denomination
	denom := collateral.GetDenomByIndex(0)
	if denom != "ucprc" {
		return fmt.Errorf("invalid denom: %s", denom)
	}

	// Parse token amount
	amt := collateral.AmountOf(denom)
	if amt.Equal(sdk.ZeroInt()) {
		return fmt.Errorf("minimum stake amount must be positive: %s", v)
	}

	return nil
}
