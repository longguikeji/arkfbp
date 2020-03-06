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
			fmt.Fprintf(os.Stderr, "")
			os.Exit(-1)
		}

		projectHome = dotarkfbp.GetProjectAbsPath(".")
		if !dotarkfbp.IsArkFbpProject(projectHome) {
			fmt.Fprintln(os.Stderr, "not an arkfbp project")
			os.Exit(-1)
		}

		// load the project information from .arkfbp directory
		metaInfo, err := dotarkfbp.LoadMetaInfo(projectHome)
		if err != nil {
			fmt.Fprintln(os.Stderr, "the .arkfbp directory is broken, please do the manual check")
			os.Exit(-1)
		}

		if metaInfo.Type == "server" && metaInfo.Language == "go" {
			CreateServerGoFlow(projectHome, metaInfo.Package, flowName)
		}
	},
}

func init() {
	RootCmd.AddCommand(createFlowCmd)
}
