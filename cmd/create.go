package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/rockl2e/arkfbp-cli/tpl"
	"github.com/rockl2e/arkfbp-cli/version"
	"github.com/spf13/cobra"
)

var (
	createParamName     string
	createParamType     string
	createParamLanguage string
)

func createFlow(home string, flow string) {
	flowDir := path.Join(home, "src", "flows")

	secs := strings.Split(flow, ".")

	for _, sec := range secs[0 : len(secs)-1] {
		flowDir = path.Join(flowDir, sec)
	}

	flowName := secs[len(secs)-1]

	flowDir = path.Join(flowDir, flowName)
	if _, err := os.Stat(flowDir); err != nil {
		os.MkdirAll(flowDir, os.ModePerm)
	}

	nodesDir := path.Join(flowDir, "nodes")
	if _, err := os.Stat(nodesDir); err != nil {
		os.MkdirAll(nodesDir, os.ModePerm)
	}

	ioutil.WriteFile(path.Join(flowDir, "index.js"), []byte(tpl.Tpl11), 0644)
	ioutil.WriteFile(path.Join(flowDir, "nodes", "node1.js"), []byte(tpl.Tpl12), 0644)
}

func create(home string) {
	os.Mkdir(home, os.ModePerm)

	ioutil.WriteFile(path.Join(home, ".babelrc"), []byte(tpl.Tpl1), 0644)

	ioutil.WriteFile(path.Join(home, "package.json"), []byte(tpl.Tpl2), 0644)
	ioutil.WriteFile(path.Join(home, "package-lock.json"), []byte(tpl.Tpl3), 0644)
	ioutil.WriteFile(path.Join(home, "webpack.config.js"), []byte(tpl.Tpl4), 0644)
	ioutil.WriteFile(path.Join(home, ".gitignore"), []byte(tpl.Tpl5), 0644)
	ioutil.WriteFile(path.Join(home, "README.md"), []byte(tpl.Tpl6), 0644)

	os.Mkdir(path.Join(home, "src"), os.ModePerm)
	ioutil.WriteFile(path.Join(home, "src", "cli.js"), []byte(tpl.Tpl7), 0644)
	ioutil.WriteFile(path.Join(home, "src", "router.js"), []byte(tpl.Tpl8), 0644)
	ioutil.WriteFile(path.Join(home, "src", "server.js"), []byte(tpl.Tpl9), 0644)

	os.Mkdir(path.Join(home, "src", "routes"), os.ModePerm)
	ioutil.WriteFile(path.Join(home, "src", "routes", "index.js"), []byte(tpl.Tpl10), 0644)

	createFlow(home, "hooks.app.beforeStart")
	createFlow(home, "hooks.app.started")
	createFlow(home, "hooks.flow.beforeCreate")
	createFlow(home, "hooks.flow.created")
	createFlow(home, "hooks.flow.executed")

	createFlow(home, "helloworld")

	cmd := exec.Command("npm", "install")
	cmd.Dir = home
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Println(string(out))
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a arkfbp project",
	Long:  `Create a arkfbp project.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("arkfbp version:", version.GetVersion())

		create(createParamName)
	},
}

func init() {
	createCmd.Flags().StringVarP(&createParamName, "name", "", "", "whether need to login into the cluster")
	createCmd.Flags().StringVarP(&createParamType, "type", "", "", "auto confirm during the installation")
	createCmd.Flags().StringVarP(&createParamLanguage, "language", "", "", "The release name")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("type")
	createCmd.MarkFlagRequired("language")

	RootCmd.AddCommand(createCmd)
}
