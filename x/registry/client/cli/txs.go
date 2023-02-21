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

func CmdCreateRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-registry [name] [stake-amount] [description] [imageUrl]",
		Short: "Broadcast message create-registry",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argStakeAmount := args[1]
			argDescription := args[2]
			argImageUrl := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRegistry(
				clientCtx.GetFromAddress().String(),
				argName,
				argStakeAmount,
				argDescription,
				argImageUrl,
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

func CmdJoinRegistryCoOperator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join-registry-co-operator [registry-id] [stake-amount]",
		Short: "Broadcast message join-registry-co-operator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRegistryId := args[0]
			argStakeAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgJoinRegistryCoOperator(
				clientCtx.GetFromAddress().String(),
				argRegistryId,
				argStakeAmount,
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

func CmdJoinRegistryMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join-registry-member [registry-id] [stake-amount]",
		Short: "Broadcast message join-registry-member",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRegistryId := args[0]
			argStakeAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgJoinRegistryMember(
				clientCtx.GetFromAddress().String(),
				argRegistryId,
				argStakeAmount,
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

func CmdWithdrawRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-rewards [amount] [registryId]",
		Short: "Broadcast message withdrawRewards",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argRegistryId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawRewards(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argRegistryId,
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
