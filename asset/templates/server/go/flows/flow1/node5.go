package flow1

import (
	"fmt"

	"git.intra.longguikeji.com/longguikeji/arkfbp-go/node"
)

// Node5 ...
type Node5 struct {
	node.FunctionNode
}

// ID ...
func (n *Node5) ID() string {
	return "Node5"
}

// Run ...
func (n *Node5) Run() interface{} {
	fmt.Println("Node5 Run...")
	m := make(map[string]string)
	m["name"] = "rock"
	return m
}

