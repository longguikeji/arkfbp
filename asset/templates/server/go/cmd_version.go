package main

import (
	"fmt"

	"{{ .PackageName }}/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of arkfbp-cli",
	Long:  `Print the version of ark.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("", version.GetVersion())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

