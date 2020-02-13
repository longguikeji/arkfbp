package flow1

import (
	"git.intra.longguikeji.com/longguikeji/arkfbp-go/flow"
)

// MyFlow ...
type MyFlow struct {
	flow.Flow
}

// New ...
func New() *MyFlow {
	f := MyFlow{
		Flow: flow.Flow{
			CreateGraph: createGraph,
		},
	}

	return &f
}
