package cmd

import (
	"fmt"
	"os"

	"github.com/gosuri/uitable"
	"github.com/longguikeji/arkfbp-cli/dotarkfbp"
	"github.com/longguikeji/arkfbp-cli/dotarkfbp/database"
	"github.com/spf13/cobra"
)

var (
	dbListTableParamDatabase string
)

var dbListTableCmd = &cobra.Command{
	Use:   "listtable",
	Short: "List all the tables in the database",
	Long:  `List all the tables in the database.`,
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

		db, err := database.GetByName(root, dbListTableParamDatabase)
		if err != nil {
			fmt.Fprintln(os.Stderr, "database not found")
			os.Exit(-1)
		}

		tables, err := database.ListTables(db)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to get all tables")
			os.Exit(-1)
		}

		table := uitable.New()
		table.AddRow("Name", "Table Type", "Description")
		for _, t := range tables {
			table.AddRow(t.Name, t.Type, t.Description)
		}
		fmt.Fprintln(os.Stdout, table)
	},
}

func init() {
	dbListTableCmd.Flags().StringVarP(&dbListTableParamDatabase, "database", "", "", "database name")
	dbListTableCmd.MarkFlagRequired("database")

	dbCmd.AddCommand(dbListTableCmd)
}
