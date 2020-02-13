package main

import (
	"{{ .PackageName }}/flows/flow1"
	"git.intra.longguikeji.com/longguikeji/arkfbp-go/flow"
)

// Routes ...
func Routes() map[string]flow.IFlow {
	handlers := make(map[string]flow.IFlow)
	handlers["/hello"] = flow1.New()
	return handlers
}
