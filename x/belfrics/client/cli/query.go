package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"belfrics/x/belfrics/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group belfrics queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListKyc())
	cmd.AddCommand(CmdShowKyc())
	cmd.AddCommand(CmdListKyc2())
	cmd.AddCommand(CmdShowKyc2())
	cmd.AddCommand(CmdListKyc3())
	cmd.AddCommand(CmdShowKyc3())
	// this line is used by starport scaffolding # 1

	return cmd
}
