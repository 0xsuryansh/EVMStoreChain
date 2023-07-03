package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"EVMStoreChain/x/evmstorechain/types"
)

func CmdListBlockstoragestate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-blockstoragestate",
		Short: "list all blockstoragestate",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllBlockstoragestateRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.BlockstoragestateAll(cmd.Context(), params)
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

func CmdShowBlockstoragestate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-blockstoragestate [blocknumber]",
		Short: "shows a blockstoragestate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argBlocknumber := args[0]

			params := &types.QueryGetBlockstoragestateRequest{
				Blocknumber: argBlocknumber,
			}

			res, err := queryClient.Blockstoragestate(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
