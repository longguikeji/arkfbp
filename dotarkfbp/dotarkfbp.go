package dotarkfbp

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/longguikeji/arkfbp-cli/constants"
	"gopkg.in/yaml.v2"
)

// DotArkFbpDir is the meta folder of the arkfbp project
const DotArkFbpDir = ".arkfbp"

// App is the structure describing the arkfbp meta info
type App struct {
	Name        string
	Description string
	CliVersion  string `yaml:"cliVersion,omitempty"`
	SpecVersion string `yaml:"specVersion,omitempty"`
	Type        string
	Language    string
	Package     string `yaml:"package,omitempty"`
	Created     string
}

// GetDotArkFbpConfigFile ...
func GetDotArkFbpConfigFile(home string) string {
	dir := path.Join(home, DotArkFbpDir)
	configFile := path.Join(dir, "config.yml")
	return configFile
}

// IsApp checks whether the project is an arkfbp project or not
func IsApp(home string) bool {
	dir := path.Join(home, DotArkFbpDir)
	configFile := path.Join(dir, "config.yml")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return false
	}

	return true
}

// LoadAppInfo ...
func LoadAppInfo(home string) (*App, error) {
	if !IsApp(home) {
		return nil, errors.New("not an arkfbp project")
	}

	configFile := GetDotArkFbpConfigFile(home)

	data, err := ioutil.ReadFile(configFile)

	var ret = App{}

	err = yaml.Unmarshal(data, &ret)
	if err != nil {
		return nil, errors.New("failed to parse .arkfbp configuration file")
	}

	return &ret, nil
}

// GetAppAbsPath ...
func GetAppAbsPath(p string) string {
	if p == "" {
		p = "."
	}

	if !path.IsAbs(p) {
		dir, _ := os.Getwd()
		p = path.Join(dir, p)
	}

	return p
}

// GetAppRoot ...
func GetAppRoot() string {
	return GetAppAbsPath(".")
}

// GetDatabasesRoot ...
func GetDatabasesRoot(app *App) (string, error) {
	t, err := constants.UnifyAppType(app.Type)
	if err != nil {
		return "", err
	}

	l, err := constants.UnifyLanguageType(app.Language)
	if err != nil {
		return "", err
	}

	if t == constants.Server {
		switch l {
		case constants.Javascript, constants.Typescript:
			return path.Join(GetAppRoot(), "src", "databases"), nil
		case constants.Go:
			return path.Join(GetAppRoot(), "databases"), nil
		case constants.Python:
			return path.Join(GetAppRoot(), "databases"), nil
		}
	}

	return "", errors.New("failed to get the database root")
}
