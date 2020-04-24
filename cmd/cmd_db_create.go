package cmd

import (
	"fmt"
	"os"

	"github.com/longguikeji/arkfbp-cli/dotarkfbp"
	"github.com/longguikeji/arkfbp-cli/dotarkfbp/database"
	"github.com/spf13/cobra"
)

var (
	dbCreateParamName        string
	dbCreateParamDescription string
	dbCreateParamEngine      string
)

var dbCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new ArkFBP DB",
	Long:  `Create a new ArkFBP DB.`,
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

		if dbCreateParamEngine != "sqlite3" && dbCreateParamEngine != "mysql" && dbCreateParamEngine != "postgres" {
			fmt.Fprintln(os.Stderr, "database engines must be one of sqlite3, mysql, postgres")
			os.Exit(-1)
		}

		err = database.Create(root, dbCreateParamName, dbCreateParamEngine, dbCreateParamDescription)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	dbCreateCmd.Flags().StringVarP(&dbCreateParamName, "name", "", "", "database name")
	dbCreateCmd.Flags().StringVarP(&dbCreateParamDescription, "description", "", "", "descripton of the database")
	dbCreateCmd.Flags().StringVarP(&dbCreateParamEngine, "engine", "", "sqlite3", "the engine of the database, choices: sqlite3, mysql, postgres")

	dbCreateCmd.MarkFlagRequired("name")

	dbCmd.AddCommand(dbCreateCmd)
}
