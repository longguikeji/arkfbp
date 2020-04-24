package cmd

import (
	"fmt"
	"os"

	"github.com/longguikeji/arkfbp-cli/dotarkfbp"
	"github.com/spf13/cobra"
)

var ()

var createFlowCmd = &cobra.Command{
	Use:   "createflow",
	Short: "Create flow in the arkfbp project",
	Long:  `Create flow in arkfbp project.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		var (
			projectHome string
			flowName    string
		)

		if len(args) > 0 {
			flowName = args[0]
		} else {
			fmt.Fprintf(os.Stderr, "not enough arguments")
			os.Exit(-1)
		}

		projectHome = dotarkfbp.GetAppAbsPath(".")
		if !dotarkfbp.IsApp(projectHome) {
			fmt.Fprintln(os.Stderr, "not an arkfbp project")
			os.Exit(-1)
		}

		// load the project information from .arkfbp directory
		metaInfo, err := dotarkfbp.LoadAppInfo(projectHome)
		if err != nil {
			fmt.Fprintln(os.Stderr, "the .arkfbp directory is broken, please do the manual check")
			os.Exit(-1)
		}

		fmt.Fprintln(os.Stdout, metaInfo)

		if metaInfo.Type == "server" && metaInfo.Language == "go" {
			CreateServerGoFlow(projectHome, metaInfo.Package, flowName)
		}

		if metaInfo.Type == "server" && (metaInfo.Language == "javascript" || metaInfo.Language == "js") {
			createServerJSFlow2(projectHome, flowName)
		}
	},
}

func init() {
	RootCmd.AddCommand(createFlowCmd)
}
