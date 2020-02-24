package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func createWebTSProject(home string) {
	var data = make(map[string]interface{})
	_ = data

	os.Mkdir(home, os.ModePerm)

	writeFile(path.Join(home, ".browserslistrc"), "asset/templates/web/ts/.browserslistrc", data)
	writeFile(path.Join(home, ".eslintrc.js"), "asset/templates/web/ts/.eslintrc.js", data)
	writeFile(path.Join(home, ".gitignore"), "asset/templates/web/ts/.gitignore", data)
	writeFile(path.Join(home, "babel.config.js"), "asset/templates/web/ts/babel.config.js", data)
	writeFile(path.Join(home, "package.json"), "asset/templates/web/ts/package.json", data)
	writeFile(path.Join(home, "postcss.config.js"), "asset/templates/web/ts/postcss.config.js", data)
	writeFile(path.Join(home, "README.md"), "asset/templates/web/ts/README.md", data)
	writeFile(path.Join(home, "tsconfig.json"), "asset/templates/web/ts/tsconfig.json", data)
	writeFile(path.Join(home, "vue.config.js"), "asset/templates/web/ts/vue.config.js", data)

	os.Mkdir(path.Join(home, "public"), os.ModePerm)
	writeFile(path.Join(home, "public", "favicon.ico"), "asset/templates/web/ts/public/favicon.ico", data)
	writeFile(path.Join(home, "public", "index.html"), "asset/templates/web/ts/public/index.html", data)

	os.Mkdir(path.Join(home, "src"), os.ModePerm)
	writeFile(path.Join(home, "src", "App.vue"), "asset/templates/web/ts/src/App.vue", data)
	writeFile(path.Join(home, "src", "main.ts"), "asset/templates/web/ts/src/main.ts", data)
	writeFile(path.Join(home, "src", "shims-tsx.d.ts"), "asset/templates/web/ts/src/shims-tsx.d.ts", data)
	writeFile(path.Join(home, "src", "shims-vue.d.ts"), "asset/templates/web/ts/src/shims-vue.d.ts", data)

	os.Mkdir(path.Join(home, "src", "assets"), os.ModePerm)
	writeFile(path.Join(home, "src", "assets", "logo.png"), "asset/templates/web/ts/src/assets/logo.png", data)

	os.Mkdir(path.Join(home, "src", "components"), os.ModePerm)
	writeFile(path.Join(home, "src", "components", "HelloWorld.vue"), "asset/templates/web/ts/src/components/HelloWorld.vue", data)

	os.Mkdir(path.Join(home, "src", "router"), os.ModePerm)
	writeFile(path.Join(home, "src", "router", "index.ts"), "asset/templates/web/ts/src/router/index.ts", data)

	os.Mkdir(path.Join(home, "src", "store"), os.ModePerm)
	writeFile(path.Join(home, "src", "store", "index.ts"), "asset/templates/web/ts/src/store/index.ts", data)

	os.Mkdir(path.Join(home, "src", "views"), os.ModePerm)
	writeFile(path.Join(home, "src", "views", "About.vue"), "asset/templates/web/ts/src/views/About.vue", data)
	writeFile(path.Join(home, "src", "views", "Home.vue"), "asset/templates/web/ts/src/views/Home.vue", data)

	os.Mkdir(path.Join(home, "src", "flows"), os.ModePerm)
	os.Mkdir(path.Join(home, "src", "flows", "helloworld"), os.ModePerm)
	os.Mkdir(path.Join(home, "src", "flows", "helloworld", "nodes"), os.ModePerm)

	writeFile(path.Join(home, "src", "flows", "helloworld", "index.ts"), "asset/templates/web/ts/src/flows/helloworld/index.ts", data)
	writeFile(path.Join(home, "src", "flows", "helloworld", "nodes", "sayGoodbye.ts"), "asset/templates/web/ts/src/flows/helloworld/nodes/sayGoodbye.ts", data)
	writeFile(path.Join(home, "src", "flows", "helloworld", "nodes", "sayHi.ts"), "asset/templates/web/ts/src/flows/helloworld/nodes/sayHi.ts", data)

	cmd := exec.Command("npm", "install")
	cmd.Dir = home
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Println(string(out))
}
