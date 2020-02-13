package main

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
	Use:   "app",
	Short: "app is a project built with arkfbp framework",
	Long:  `app is a project built with arkfbp framework`,
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
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}
