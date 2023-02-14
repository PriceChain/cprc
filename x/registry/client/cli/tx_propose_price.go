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
		Use:   "propose-price [registry-id] [store-name] [store-addr] [purchase-time] [prod-name] [price] [receipt-code] [reserved]",
		Short: "Broadcast message propose-price",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRegistryId := args[0]
			argStoreName := args[1]
			argStoreAddr := args[2]
			argPurchaseTime := args[3]
			argProdName := args[4]
			argPrice := args[5]
			argReceiptCode := args[6]
			argReserved := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProposePrice(
				clientCtx.GetFromAddress().String(),
				argRegistryId,
				argStoreName,
				argStoreAddr,
				argPurchaseTime,
				argProdName,
				argPrice,
				argReceiptCode,
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
