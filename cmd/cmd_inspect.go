package cmd

import (
	"fmt"
	"os"

	"github.com/longguikeji/arkfbp-cli/dotarkfbp"
	"github.com/spf13/cobra"
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
		}

		p = dotarkfbp.GetAppAbsPath(p)

		if !dotarkfbp.IsApp(p) {
			fmt.Fprintln(os.Stderr, "not an arkfbp project")
			os.Exit(-1)
		}

		metaInfo, err := dotarkfbp.LoadAppInfo(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, "the .arkfbp directory is broken, please do the manual check")
			os.Exit(-1)
		}

		fmt.Fprintf(os.Stdout, "Name: %s\n", metaInfo.Name)
		fmt.Fprintf(os.Stdout, "CLI version: %s\n", metaInfo.CliVersion)
		fmt.Fprintf(os.Stdout, "Spec version: %s\n", metaInfo.SpecVersion)
		fmt.Fprintf(os.Stdout, "Type: %s\n", metaInfo.Type)
		fmt.Fprintf(os.Stdout, "Language: %s\n", metaInfo.Language)
		fmt.Fprintf(os.Stdout, "Package: %s\n", metaInfo.Package)
		fmt.Fprintf(os.Stdout, "Generated at: %s\n", metaInfo.Created)
	},
}

func init() {
	RootCmd.AddCommand(inspectCmd)
}
