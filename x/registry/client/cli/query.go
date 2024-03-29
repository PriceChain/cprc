package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/PriceChain/cprc/x/registry/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group registry queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListRegistry())
	cmd.AddCommand(CmdShowRegistry())
	cmd.AddCommand(CmdListRegistryOwner())
	cmd.AddCommand(CmdShowRegistryOwner())
	cmd.AddCommand(CmdListRegistryMember())
	cmd.AddCommand(CmdShowRegistryMember())
	cmd.AddCommand(CmdListRegistryStakedAmount())
	cmd.AddCommand(CmdShowRegistryStakedAmount())
	cmd.AddCommand(CmdListStakedAmountPerWallet())
	cmd.AddCommand(CmdShowStakedAmountPerWallet())
	cmd.AddCommand(CmdListPriceData())
	cmd.AddCommand(CmdShowPriceData())
	// this line is used by starport scaffolding # 1

	return cmd
}
