package cmd

import (
	"fmt"

	"github.com/rockl2e/arkfbp-cli/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of arkfbp-cli",
	Long:  `Print the version of arkfbp-cli.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("arkfbp version:", version.GetVersion())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
