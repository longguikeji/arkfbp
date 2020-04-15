package constants

import (
	"errors"
	"strings"
)

// AppType represents the ArkFBP project type
type AppType uint8

func (t AppType) String() string {
	switch t {
	case Server:
		return "server"
	case Web:
		return "web"
	case Script:
		return "script"
	}

	return ""
}

const (
	Server AppType = iota
	Web
	Script
)

// LanguageType presents a language type
type LanguageType uint8

func (t LanguageType) String() string {
	switch t {
	case Python:
		return "python"
	case Javascript:
		return "javascript"
	case Typescript:
		return "typescript"
	case Go:
		return "go"
	case Java:
		return "java"
	}

	return ""
}

const (
	Python LanguageType = iota
	Javascript
	Typescript
	Go
	Java
)

// UnifyLanguageType parses a language description in string format, and return the unified definition in arkfbp
func UnifyLanguageType(language string) (LanguageType, error) {
	language = strings.ToLower(language)
	switch language {
	case "python":
		return Python, nil
	case "javascript", "nodejs":
		return Javascript, nil
	case "typescript":
		return Typescript, nil
	case "go", "golang":
		return Go, nil
	case "java":
		return Java, nil
	}

	return 255, errors.New("unsupported language")
}

// UnifyAppType parses a type description in string format and return the unified defition
func UnifyAppType(t string) (AppType, error) {
	t = strings.ToLower(t)
	switch t {
	case "server":
		return Server, nil
	case "web":
		return Web, nil
	case "script":
		return Script, nil
	}

	return 255, errors.New("unsupported app type")
}
