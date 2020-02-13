package cmd

import (
	"fmt"

	"git.intra.longguikeji.com/longguikeji/arkfbp-cli/version"
	"github.com/spf13/cobra"
)

var (
	createParamName        string
	createParamType        string
	createParamLanguage    string
	createParamPackageName string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a arkfbp project",
	Long:  `Create a arkfbp project.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("arkfbp version:", version.GetVersion())

		switch createParamType {
		case "server":
			{
				switch createParamLanguage {
				case "javascript", "js":
					createServerJSProject(createParamName)
				case "typescript", "ts":
					createServerTSProject(createParamName)
				case "go":
					if createParamPackageName == "" {
						panic("for go project must give the package name")
					}
					createServerGoProject(createParamName, createParamPackageName)
				default:
					panic("server side, we support: node, go, python, java")
				}
			}

		case "web":
			{
				switch createParamLanguage {
				case "javascript", "js":
					createWebJSProject(createParamName)
				case "typescript", "ts":
					createWebTSProject(createParamName)
				default:
					panic("web side, we support: javascript & typescript")
				}
			}

		default:
			panic("only server & web support")
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&createParamName, "name", "", "", "project name")
	createCmd.Flags().StringVarP(&createParamType, "type", "", "", "project type, web | server")
	createCmd.Flags().StringVarP(&createParamLanguage, "language", "", "", "project language, javascript | typescript | python | go | java")
	createCmd.Flags().StringVarP(&createParamPackageName, "package-name", "", "", "package name: go | java")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("type")
	createCmd.MarkFlagRequired("language")

	RootCmd.AddCommand(createCmd)
}
