package cmd

import (
	"fmt"
	"os"

	"github.com/longguikeji/arkfbp-cli/dotarkfbp"
	"github.com/longguikeji/arkfbp-cli/dotarkfbp/database"
	"github.com/spf13/cobra"
)

var (
	dbCreateTableParamDatabase    string
	dbCreateTableParamName        string
	dbCreateTableParamType        string
	dbCreateTableParamDescription string
)

var dbCreateTableCmd = &cobra.Command{
	Use:   "createtable",
	Short: "Create a new table",
	Long:  `Create a new table.`,
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
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		db, err := database.GetByName(root, dbCreateTableParamDatabase)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		if dbCreateTableParamType != "abstract" && dbCreateTableParamType != "standard" {
			fmt.Fprintln(os.Stderr, "table type must be abstrct or standard")
			os.Exit(-1)
		}

		err = database.CreateTable(db, dbCreateTableParamName, dbCreateTableParamType, dbCreateParamDescription)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	dbCreateTableCmd.Flags().StringVarP(&dbCreateTableParamDatabase, "database", "", "", "database name")
	dbCreateTableCmd.Flags().StringVarP(&dbCreateTableParamName, "name", "", "", "table name")
	dbCreateTableCmd.Flags().StringVarP(&dbCreateTableParamType, "type", "", "standard", "table type")
	dbCreateTableCmd.Flags().StringVarP(&dbCreateParamDescription, "description", "", "", "table description")

	dbCreateTableCmd.MarkFlagRequired("database")
	dbCreateTableCmd.MarkFlagRequired("name")

	dbCmd.AddCommand(dbCreateTableCmd)
}
