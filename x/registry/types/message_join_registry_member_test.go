package types

import (
	"testing"

	"github.com/PriceChain/rd_net/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgJoinRegistryMember_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgJoinRegistryMember
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgJoinRegistryMember{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgJoinRegistryMember{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
