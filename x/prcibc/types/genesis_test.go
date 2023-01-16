package types_test

import (
	"testing"

	"github.com/PriceChain/cprc/x/prcibc/types"
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
				PortId: types.PortID,
				IbcMsgList: []types.IbcMsg{
	{
		Id: 0,
	},
	{
		Id: 1,
	},
},
IbcMsgCount: 2,
// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
	desc:     "duplicated ibcMsg",
	genState: &types.GenesisState{
		IbcMsgList: []types.IbcMsg{
			{
				Id: 0,
			},
			{
				Id: 0,
			},
		},
	},
	valid:    false,
},
{
	desc:     "invalid ibcMsg count",
	genState: &types.GenesisState{
		IbcMsgList: []types.IbcMsg{
			{
				Id: 1,
			},
		},
		IbcMsgCount: 0,
	},
	valid:    false,
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
