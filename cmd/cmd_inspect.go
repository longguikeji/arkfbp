package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect the info of the arkfbp project",
	Long:  `Inspect the info of the arkfbp project.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		var p string

		if len(args) > 0 {
			p = args[0]
		} else {
			p = "."
		}

		if !path.IsAbs(p) {
			dir, _ := os.Getwd()
			p = path.Join(dir, p)
		}

		dotFbpDir := path.Join(p, ".arkfbp")
		dotFbpConfigFile := path.Join(dotFbpDir, "config.yml")

		if _, err := os.Stat(dotFbpDir); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "not an arkfbp project")
			os.Exit(-1)
		}

		if _, err := os.Stat(dotFbpConfigFile); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "not an arkfbp project")
			os.Exit(-1)
		}

		data, err := ioutil.ReadFile(dotFbpConfigFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Please check the data under the .arkfbp")
			os.Exit(-1)
		}

		var ret map[string]interface{}

		err = yaml.Unmarshal(data, &ret)
		if err != nil {
			fmt.Fprintln(os.Stderr, "the .arkfbp directory is broken, please do the manual check")
			os.Exit(-1)
		}

		fmt.Fprintf(os.Stdout, "name: %s\n", ret["name"].(string))
		fmt.Fprintf(os.Stdout, "arkfbp-cli version: %s\n", ret["arkfbpCliVersion"].(string))
		fmt.Fprintf(os.Stdout, "arkfbp-spec version: %s\n", ret["arkfbpSpecVersion"].(string))
		fmt.Fprintf(os.Stdout, "type: %s\n", ret["type"].(string))
		fmt.Fprintf(os.Stdout, "language: %s\n", ret["language"].(string))

		if v, ok := ret["package"].(string); ok {
			fmt.Fprintf(os.Stdout, "package: %s\n", v)
		}

		fmt.Fprintf(os.Stdout, "generated at: %s\n", ret["created"].(string))
	},
}

func init() {
	RootCmd.AddCommand(inspectCmd)
}
