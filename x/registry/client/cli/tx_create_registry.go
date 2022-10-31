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

func CmdCreateRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-registry [name] [stake-amount] [quorum] [consensus-expiring-time]",
		Short: "Broadcast message create-registry",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argStakeAmount := args[1]
			argQuorum := args[2]
			argConsensusExpiringTime := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRegistry(
				clientCtx.GetFromAddress().String(),
				argName,
				argStakeAmount,
				argQuorum,
				argConsensusExpiringTime,
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
