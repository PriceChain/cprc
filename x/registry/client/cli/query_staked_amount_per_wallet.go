package cli

import (
	"context"

	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListStakedAmountPerWallet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-staked-amount-per-wallet",
		Short: "list all stakedAmountPerWallet",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllStakedAmountPerWalletRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.StakedAmountPerWalletAll(context.Background(), params)
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

func CmdShowStakedAmountPerWallet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-staked-amount-per-wallet [index]",
		Short: "shows a stakedAmountPerWallet",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetStakedAmountPerWalletRequest{
				Index: argIndex,
			}

			res, err := queryClient.StakedAmountPerWallet(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
