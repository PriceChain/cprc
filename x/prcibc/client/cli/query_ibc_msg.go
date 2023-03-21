package cli

import (
	"context"
	"strconv"

	"github.com/PriceChain/cprc/x/prcibc/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListIbcMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-ibc-msg",
		Short: "list all ibcMsg",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllIbcMsgRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.IbcMsgAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowIbcMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-ibc-msg [id]",
		Short: "shows a ibcMsg",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetIbcMsgRequest{
				Id: id,
			}

			res, err := queryClient.IbcMsg(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
