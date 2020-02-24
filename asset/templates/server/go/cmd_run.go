package main

import (
	"encoding/json"
	"fmt"
	"os"

	"git.intra.longguikeji.com/longguikeji/arkfbp-go/intr"
	"github.com/spf13/cobra"
)

var (
	runCmdParamFlowName string
	runCmdParamInputs   string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the flow",
	Long:  `Run the flow.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		var (
			handler intr.IFlow  = nil
			inputs  interface{} = nil
		)

		if len(runCmdParamInputs) > 0 {
			if err := json.Unmarshal([]byte(runCmdParamInputs), &inputs); err != nil {
				fmt.Println("inputs cannot be converted  to a JSON object, please check your inputs setting")
				os.Exit(1)
			}
		}

		routes := Routes()

		for _, route := range routes {
			if route.Name == runCmdParamFlowName {
				handler = route.Handler
				break
			}
		}

		_ = handler

		if handler == nil {
			fmt.Println("no matched flow found")
			os.Exit(1)
		}

		handler.Run(inputs)
	},
}

func init() {
	runCmd.Flags().StringVarP(&runCmdParamFlowName, "name", "", "127.0.0.1", "flow to execute")
	runCmd.Flags().StringVarP(&runCmdParamInputs, "inputs", "", "5000", "inputs data")

	RootCmd.AddCommand(runCmd)
}
