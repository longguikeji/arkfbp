package cmd

import (
	"fmt"
	"os"

	"github.com/longguikeji/arkfbp-cli/dotarkfbp"
	"github.com/spf13/cobra"
)

var (
	createNodeParamFlowName      string
	createNodeParamID            string
	createNodeParamClassName     string
	createNodeParamBaseClassName string
)

var createNodeCmd = &cobra.Command{
	Use:   "createnode",
	Short: "Create node file of one flow",
	Long:  `Create node file of one flow.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		var (
			projectHome string
		)

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
			CreateServerGoFlowNode(
				projectHome,
				metaInfo.Package,
				createNodeParamFlowName,
				createNodeParamID,
				createNodeParamClassName,
				createNodeParamBaseClassName,
			)
		}

		if metaInfo.Type == "server" && (metaInfo.Language == "javascript" || metaInfo.Language == "js") {
			createServerJSFlowNode(
				projectHome,
				createNodeParamFlowName,
				createNodeParamID,
				createNodeParamClassName,
				createNodeParamBaseClassName,
			)
		}
	},
}

func init() {
	createNodeCmd.Flags().StringVarP(&createNodeParamFlowName, "flow", "", "", "flow name")
	createNodeCmd.Flags().StringVarP(&createNodeParamID, "id", "", "", "node id")
	createNodeCmd.Flags().StringVarP(&createNodeParamClassName, "class", "", "", "node class")
	createNodeCmd.Flags().StringVarP(&createNodeParamBaseClassName, "base", "", "FunctionNode", "node base class")

	createNodeCmd.MarkFlagRequired("flow")
	createNodeCmd.MarkFlagRequired("id")
	createNodeCmd.MarkFlagRequired("class")

	RootCmd.AddCommand(createNodeCmd)
}
