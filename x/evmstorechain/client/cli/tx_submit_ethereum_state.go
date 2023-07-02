package cli

import (
	"strconv"

	"EVMStoreChain/x/evmstorechain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitEthereumState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-ethereum-state [blocknumber] [state]",
		Short: "Broadcast message SubmitEthereumState",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBlocknumber, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argState, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitEthereumState(
				clientCtx.GetFromAddress().String(),
				argBlocknumber,
				argState,
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
