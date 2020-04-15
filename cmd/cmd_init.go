package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"time"

	"github.com/longguikeji/arkfbp-cli/constants"
	"github.com/longguikeji/arkfbp-cli/version"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	initParamName        string
	initParamType        string
	initParamLanguage    string
	initParamPackageName string
	initParamFramework   string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init an arkfbp project in the exsiting folder",
	Long:  `Init an arkfbp project in the exsting folder.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		pName := path.Base(initParamName)

		appType, err := constants.UnifyAppType(initParamType)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		languageType, err := constants.UnifyLanguageType(initParamLanguage)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		switch appType {
		case constants.Server:
			{
				switch languageType {
				case constants.Javascript, constants.Go, constants.Java, constants.Python, constants.Typescript:
				default:
					panic("server app, we support: javascript, typescript, go, python, java")
				}
			}

		case constants.Web:
			{
				switch languageType {
				case constants.Javascript, constants.Typescript:
				default:
					panic("web app, we support: javascript & typescript")
				}
			}
		case constants.Script:
			{
				switch languageType {
				case constants.Javascript, constants.Go, constants.Java, constants.Python, constants.Typescript:
				default:
					panic("script app, we support: javascript, typescript, go, python, java")
				}
			}

		default:
			panic("only server & web & script app support")
		}

		// Generate .arkfbp directory
		dotFbpDir := path.Join(createParamName, ".arkfbp")
		os.Mkdir(dotFbpDir, os.ModePerm)
		dotFbpConfigFile := path.Join(dotFbpDir, "config.yml")

		data := make(map[string]interface{})

		data["name"] = pName
		data["arkfbpCliVersion"] = version.GetVersion()
		data["arkfbpSpecVersion"] = version.GetArkFBPSpecVersion()
		data["type"] = appType.String()
		data["language"] = languageType.String()
		if initParamPackageName != "" {
			data["package"] = initParamPackageName
		}
		if initParamFramework != "" {
			data["framework"] = initParamFramework
		}
		data["created"] = time.Now().String()

		out, _ := yaml.Marshal(data)
		ioutil.WriteFile(dotFbpConfigFile, out, 0644)
	},
}

func init() {
	initCmd.Flags().StringVarP(&initParamName, "name", "", "", "project name")
	initCmd.Flags().StringVarP(&initParamType, "type", "t", "", "project type, web | server | script")
	initCmd.Flags().StringVarP(&initParamLanguage, "language", "l", "", "project language, javascript | typescript | python | go | java")
	initCmd.Flags().StringVarP(&initParamPackageName, "package", "p", "", "package name: go | java")
	initCmd.Flags().StringVarP(&initParamFramework, "framework", "", "", "python: django | flask")

	initCmd.MarkFlagRequired("name")
	initCmd.MarkFlagRequired("type")
	initCmd.MarkFlagRequired("language")

	RootCmd.AddCommand(initCmd)
}
