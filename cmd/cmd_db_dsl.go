package cmd

import (
	"github.com/spf13/cobra"
)

var dbDSLCmd = &cobra.Command{
	Use:   "dsl",
	Short: "Generate the DSL in relevant languages",
	Long:  `Generate the DSL codes in relevant languages/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
	},
}

func init() {
	dbCmd.AddCommand(dbDSLCmd)
}
