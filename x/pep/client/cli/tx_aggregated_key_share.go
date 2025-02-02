package cli

import (
	"fairyring/x/pep/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateAggregatedKeyShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-aggregated-key-share [height] [data]",
		Short: "Submit a new aggregated keyshare into a destination chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexHeight, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// Get value arguments
			argData := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAggregatedKeyShare(
				clientCtx.GetFromAddress().String(),
				indexHeight,
				argData,
			)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
