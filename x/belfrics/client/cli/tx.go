package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"belfrics/x/belfrics/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateKyc())
	cmd.AddCommand(CmdUpdateKyc())
	cmd.AddCommand(CmdDeleteKyc())
	cmd.AddCommand(CmdCreateKyc2())
	cmd.AddCommand(CmdUpdateKyc2())
	cmd.AddCommand(CmdDeleteKyc2())
	cmd.AddCommand(CmdCreateKyc3())
	cmd.AddCommand(CmdUpdateKyc3())
	cmd.AddCommand(CmdDeleteKyc3())
	// this line is used by starport scaffolding # 1

	return cmd
}
