package cli

import (
	"strconv"

	"fantasfive/x/fantasfive/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateTeam() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-team [player-0] [player-1] [player-2] [player-3] [player-4]",
		Short: "Broadcast message createTeam",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPlayer0 := args[0]
			argPlayer1 := args[1]
			argPlayer2 := args[2]
			argPlayer3 := args[3]
			argPlayer4 := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTeam(
				clientCtx.GetFromAddress().String(),
				argPlayer0,
				argPlayer1,
				argPlayer2,
				argPlayer3,
				argPlayer4,
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
