package transaction

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Command about transactions",
	RunE: func(cmd *cobra.Command, args []string)error{
		return cmd.Usage()
	},
}