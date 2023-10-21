package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"fantasfive/x/fantasfive/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group fantasfive queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListSystemInfo())
	cmd.AddCommand(CmdShowSystemInfo())
	cmd.AddCommand(CmdListStoredMW())
	cmd.AddCommand(CmdShowStoredMW())
	cmd.AddCommand(CmdListStoredTeam())
	cmd.AddCommand(CmdShowStoredTeam())
	cmd.AddCommand(CmdShowMwInfo())
	cmd.AddCommand(CmdShowTeamInfo())
	// this line is used by starport scaffolding # 1

	return cmd
}
