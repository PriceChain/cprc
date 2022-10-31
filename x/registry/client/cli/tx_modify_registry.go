package cli

import (
	"strconv"

	"github.com/PriceChain/rd_net/x/registry/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdModifyRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "modify-registry [registry-id] [stake-amount] [name] [quorum] [consensus-expring-time] [reason]",
		Short: "Broadcast message modify-registry",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRegistryId := args[0]
			argStakeAmount := args[1]
			argName := args[2]
			argQuorum := args[3]
			argConsensusExpringTime := args[4]
			argReason := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgModifyRegistry(
				clientCtx.GetFromAddress().String(),
				argRegistryId,
				argStakeAmount,
				argName,
				argQuorum,
				argConsensusExpringTime,
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
