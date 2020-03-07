package flow1

import "github.com/longguikeji/arkfbp-go/node"

// Node3 ...
type Node3 struct {
	node.StopNode
}

// ID ...
func (n *Node3) ID() string {
	return "Node3"
}
