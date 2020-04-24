package cmd

import (
	"fmt"
	"os"

	"github.com/gosuri/uitable"
	"github.com/longguikeji/arkfbp-cli/dotarkfbp"
	"github.com/longguikeji/arkfbp-cli/dotarkfbp/database"
	"github.com/spf13/cobra"
)

var dbListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the databases",
	Long:  `List all the databases.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		home, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		app, err := dotarkfbp.LoadAppInfo(home)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		root, err := dotarkfbp.GetDatabasesRoot(app)

		databases, err := database.List(root)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		table := uitable.New()
		table.AddRow("Name", "Engine", "Description")
		for _, database := range databases {
			table.AddRow(database.Name, database.Engine, database.Description)
		}
		fmt.Fprintln(os.Stdout, table)
	},
}

func init() {
	dbCmd.AddCommand(dbListCmd)
}
