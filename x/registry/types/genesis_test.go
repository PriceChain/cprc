package types_test

import (
	"testing"

	"github.com/PriceChain/rd_net/x/registry/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				RegistryList: []types.Registry{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				RegistryCount: 2,
				RegistryOwnerList: []types.RegistryOwner{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				RegistryOwnerCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated registry",
			genState: &types.GenesisState{
				RegistryList: []types.Registry{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid registry count",
			genState: &types.GenesisState{
				RegistryList: []types.Registry{
					{
						Id: 1,
					},
				},
				RegistryCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated registryOwner",
			genState: &types.GenesisState{
				RegistryOwnerList: []types.RegistryOwner{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid registryOwner count",
			genState: &types.GenesisState{
				RegistryOwnerList: []types.RegistryOwner{
					{
						Id: 1,
					},
				},
				RegistryOwnerCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
