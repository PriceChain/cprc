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

func CmdUnbondRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unbond-registry [registry-id] [unbond-amount] [reason]",
		Short: "Broadcast message unbond-registry",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRegistryId := args[0]
			argUnbondAmount := args[1]
			argReason := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUnbondRegistry(
				clientCtx.GetFromAddress().String(),
				argRegistryId,
				argUnbondAmount,
				argReason,
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
