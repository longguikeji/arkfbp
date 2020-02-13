package flow1

import (
	"fmt"

	"git.intra.longguikeji.com/longguikeji/arkfbp-go/node"
)

// Node4 ...
type Node4 struct {
	node.IFNode
}

// ID ...
func (n Node4) ID() string {
	return "Node4"
}

// Positive ...
func (n Node4) Positive() interface{} {
	fmt.Println("Node4 IF Postive")
	return nil
}

// Negative ...
func (n Node4) Negative() interface{} {
	fmt.Println("Node4 IF Negative")
	return nil
}
