package flow1

import (
	"fmt"

	"git.intra.longguikeji.com/longguikeji/arkfbp-go/node"
)

// Node1 ...
type Node1 struct {
	node.StartNode
}

// ID ...
func (n Node1) ID() string {
	return "Node1"
}

// Next ...
func (n Node1) Next() string {
	return "Node2"
}

// Run ...
func (n Node1) Run() interface{} {
	fmt.Println("Node1 Run...")
	return "node1ret"
}
