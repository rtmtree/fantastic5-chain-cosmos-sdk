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

func CmdAnnounceAndCreateNextMw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "announce-and-create-next-mw [mw-id] [player-perf]",
		Short: "Broadcast message announceAndCreateNextMw",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMwId := args[0]
			argPlayerPerf := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAnnounceAndCreateNextMw(
				clientCtx.GetFromAddress().String(),
				argMwId,
				argPlayerPerf,
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
