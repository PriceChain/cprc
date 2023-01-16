package cli

import (
	"strconv"

	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVotePrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-price [registry-id] [answer] [reserved]",
		Short: "Broadcast message vote-price",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRegistryId := args[0]
			argAnswer := args[1]
			argReserved := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVotePrice(
				clientCtx.GetFromAddress().String(),
				argRegistryId,
				argAnswer,
				argReserved,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
