package cmd

import (
	goflag "flag"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	paramVerbose      string
	paramOutputFormat string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "arkfbp",
	Short: "arkfbp is the standard tooling tools of arkfbp.",
	Long:  `arkfbp is the standard tooling tools of arkfbp.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&paramVerbose, "verbose", "", "5", "the verbose level")
	RootCmd.PersistentFlags().StringVarP(&paramOutputFormat, "output-format", "", "", "the output format")

	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}
