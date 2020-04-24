package cmd

import (
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db [subcommand]",
	Short: "DB stuff",
	Long:  `Operate on ArkFBP DB.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		RootCmd.PersistentPreRun(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(dbCmd)
}
