package flow1

import (
	"fmt"

	"git.intra.longguikeji.com/longguikeji/arkfbp-go/node"
)

// Node6 ...
type Node6 struct {
	node.FunctionNode
}

// ID ...
func (n *Node6) ID() string {
	return "Node6"
}

// Run ...
func (n *Node6) Run() interface{} {
	fmt.Println("Node6 Run...")
	return nil
}
