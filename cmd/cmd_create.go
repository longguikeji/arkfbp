package cmd

import (
	"io/ioutil"
	"os"
	"path"

	"time"

	"github.com/longguikeji/arkfbp-cli/version"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	createParamName        string
	createParamType        string
	createParamLanguage    string
	createParamPackageName string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an arkfbp project",
	Long:  `Create an arkfbp project.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		pName := path.Base(createParamName)

		switch createParamType {
		case "server":
			{
				switch createParamLanguage {
				case "javascript", "js":
					createServerJSProject(pName)
				case "typescript", "ts":
					createServerTSProject(pName)
				case "go":
					if createParamPackageName == "" {
						panic("for go project must give the package name")
					}
					createServerGoProject(pName, createParamPackageName)
				default:
					panic("server side, we support: node, go, python, java")
				}
			}

		case "web":
			{
				switch createParamLanguage {
				case "javascript", "js":
					createWebJSProject(pName)
				case "typescript", "ts":
					createWebTSProject(pName)
				default:
					panic("web side, we support: javascript & typescript")
				}
			}

		default:
			panic("only server & web support")
		}

		// Generate .arkfbp directory
		dotFbpDir := path.Join(createParamName, ".arkfbp")
		os.Mkdir(dotFbpDir, os.ModePerm)
		dotFbpConfigFile := path.Join(dotFbpDir, "config.yml")

		data := make(map[string]interface{})

		data["name"] = pName
		data["arkfbpCliVersion"] = version.GetVersion()
		data["arkfbpSpecVersion"] = version.GetArkFBPSpecVersion()
		data["type"] = createParamType
		data["language"] = createParamLanguage
		if createParamPackageName != "" {
			data["package"] = createParamPackageName
		}
		data["created"] = time.Now().String()

		out, _ := yaml.Marshal(data)
		ioutil.WriteFile(dotFbpConfigFile, out, 0644)
	},
}

func init() {
	createCmd.Flags().StringVarP(&createParamName, "name", "", "", "project name")
	createCmd.Flags().StringVarP(&createParamType, "type", "t", "", "project type, web | server")
	createCmd.Flags().StringVarP(&createParamLanguage, "language", "l", "", "project language, javascript | typescript | python | go | java")
	createCmd.Flags().StringVarP(&createParamPackageName, "package", "p", "", "package name: go | java")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("type")
	createCmd.MarkFlagRequired("language")

	RootCmd.AddCommand(createCmd)
}
