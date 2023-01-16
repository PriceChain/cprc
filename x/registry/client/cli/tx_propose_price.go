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

func CmdProposePrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-price [registry-id] [price] [prod-info] [reserved]",
		Short: "Broadcast message propose-price",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRegistryId := args[0]
			argPrice := args[1]
			argProdInfo := args[2]
			argReserved := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProposePrice(
				clientCtx.GetFromAddress().String(),
				argRegistryId,
				argPrice,
				argProdInfo,
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
