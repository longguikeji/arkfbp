package flow1

import (
	"fmt"

	"git.intra.longguikeji.com/longguikeji/arkfbp-go/node"
)

// Node2 ...
type Node2 struct {
	node.NopNode
}

// ID ...
func (n Node2) ID() string {
	return "Node2"
}

// Next ...
func (n Node2) Next() string {
	return "Node3"
}

// Run ...
func (n Node2) Run() interface{} {
	fmt.Println("Node2 Run...")
	return nil
}
