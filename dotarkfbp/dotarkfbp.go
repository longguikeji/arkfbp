package dotarkfbp

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// DotArkFbpDir is the meta folder of the arkfbp project
const DotArkFbpDir = ".arkfbp"

// ArkFbpProject is the structure describing the arkfbp meta info
type ArkFbpProject struct {
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

// IsArkFbpProject checks whether the project is an arkfbp project or not
func IsArkFbpProject(home string) bool {
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

// LoadMetaInfo ...
func LoadMetaInfo(home string) (*ArkFbpProject, error) {
	if !IsArkFbpProject(home) {
		return nil, errors.New("not an arkfbp project")
	}

	configFile := GetDotArkFbpConfigFile(home)

	data, err := ioutil.ReadFile(configFile)

	var ret = ArkFbpProject{}

	err = yaml.Unmarshal(data, &ret)
	if err != nil {
		return nil, errors.New("failed to parse .arkfbp configuration file")
	}

	return &ret, nil
}

// GetProjectAbsPath ...
func GetProjectAbsPath(p string) string {
	if p == "" {
		p = "."
	}

	if !path.IsAbs(p) {
		dir, _ := os.Getwd()
		p = path.Join(dir, p)
	}

	return p
}
