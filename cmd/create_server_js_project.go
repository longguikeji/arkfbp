package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func createServerJSFlow(home string, flow string) {
	var data = make(map[string]interface{})

	flowDir := path.Join("src", "flows")

	secs := strings.Split(flow, ".")

	for _, sec := range secs[0 : len(secs)-1] {
		flowDir = path.Join(flowDir, sec)
	}

	flowName := secs[len(secs)-1]

	flowDir = path.Join(flowDir, flowName)
	if _, err := os.Stat(flowDir); err != nil {
		os.MkdirAll(path.Join(home, flowDir), os.ModePerm)
	}

	nodesDir := path.Join(flowDir, "nodes")
	if _, err := os.Stat(nodesDir); err != nil {
		os.MkdirAll(path.Join(home, nodesDir), os.ModePerm)
	}

	writeFile(path.Join(home, flowDir, "index.js"), path.Join("asset/templates/server/js/", flowDir, "index.js"), data)
	writeFile(path.Join(home, flowDir, "nodes", "node1.js"), path.Join("asset/templates/server/js/", flowDir, "nodes", "node1.js"), data)
}

func createServerJSFlow2(home string, flow string) {
	var data = make(map[string]interface{})

	flowDir := path.Join("src", "flows")

	secs := strings.Split(flow, ".")

	for _, sec := range secs[0 : len(secs)-1] {
		flowDir = path.Join(flowDir, sec)
	}

	flowName := secs[len(secs)-1]

	flowDir = path.Join(flowDir, flowName)
	if _, err := os.Stat(flowDir); err != nil {
		os.MkdirAll(path.Join(home, flowDir), os.ModePerm)
	}

	nodesDir := path.Join(flowDir, "nodes")
	if _, err := os.Stat(nodesDir); err != nil {
		os.MkdirAll(path.Join(home, nodesDir), os.ModePerm)
	}

	writeFile(path.Join(home, flowDir, "index.js"), path.Join("asset/templates/server/js/src/flows/helloworld/index.js"), data)
	writeFile(path.Join(home, flowDir, "nodes", "node1.js"), path.Join("asset/templates/server/js/src/flows/helloworld/nodes", "node1.js"), data)
}

func createServerJSFlowNode(home, flowName, id, className, baseClassName string) {
	var data = make(map[string]interface{})

	baseClassName1 := baseClassName

	switch baseClassName {
	case "APINode":
		baseClassName1 = "apiNode"
	case "StartNode":
		baseClassName1 = "startNode"
	case "StopNode":
		baseClassName1 = "stopNode"
	case "NopNode":
		baseClassName1 = "nopNode"
	case "FunctionNode":
		baseClassName1 = "functionNode"
	case "IFNode":
		baseClassName1 = "ifNode"
	case "TestNode":
		baseClassName1 = "testNode"
	}

	baseClassName1 = strings.ToLower(baseClassName1[0:1]) + baseClassName1[1:]

	fileName := className
	fileName = strings.ToLower(fileName[0:1]) + fileName[1:] + ".js"

	data["BaseClassName"] = baseClassName
	data["BaseClassName1"] = baseClassName1
	data["ClassName"] = className

	flowDir := path.Join("src", "flows")

	secs := strings.Split(flowName, ".")

	for _, sec := range secs {
		flowDir = path.Join(flowDir, sec)
	}

	if _, err := os.Stat(flowDir); err != nil {
		os.MkdirAll(path.Join(home, flowDir), os.ModePerm)
	}

	nodesDir := path.Join(flowDir, "nodes")
	if _, err := os.Stat(nodesDir); err != nil {
		os.MkdirAll(path.Join(home, nodesDir), os.ModePerm)
	}

	writeFile(path.Join(home, flowDir, "nodes", fileName), path.Join("asset/templates/server/js/src/flows/helloworld/nodes", "node.tpl"), data)
}

func createServerJSProject(home string) {
	var data = make(map[string]interface{})

	os.Mkdir(home, os.ModePerm)

	writeFile(path.Join(home, ".babelrc"), "asset/templates/server/js/.babelrc", data)

	writeFile(path.Join(home, "package.json"), "asset/templates/server/js/package.json", data)
	writeFile(path.Join(home, "webpack.config.js"), "asset/templates/server/js/webpack.config.js", data)
	writeFile(path.Join(home, ".gitignore"), "asset/templates/server/js/.gitignore", data)
	writeFile(path.Join(home, "README.md"), "asset/templates/server/js/README.md", data)

	os.Mkdir(path.Join(home, "src"), os.ModePerm)
	writeFile(path.Join(home, "src", "cli.js"), "asset/templates/server/js/src/cli.js", data)
	writeFile(path.Join(home, "src", "router.js"), "asset/templates/server/js/src/router.js", data)
	writeFile(path.Join(home, "src", "server.js"), "asset/templates/server/js/src/server.js", data)

	os.Mkdir(path.Join(home, "src", "routes"), os.ModePerm)
	writeFile(path.Join(home, "src", "routes", "index.js"), "asset/templates/server/js/src/routes/index.js", data)

	createServerJSFlow(home, "hooks.app.beforeStart")
	createServerJSFlow(home, "hooks.app.started")
	createServerJSFlow(home, "hooks.flow.beforeCreate")
	createServerJSFlow(home, "hooks.flow.created")
	createServerJSFlow(home, "hooks.flow.executed")

	// createFlow(home, "helloworld")

	cmd := exec.Command("npm", "install")
	cmd.Dir = home
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Println(string(out))
}
