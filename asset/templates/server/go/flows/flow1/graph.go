package flow1

import (
	"fmt"

	"git.intra.longguikeji.com/longguikeji/arkfbp-go/graph"
	"git.intra.longguikeji.com/longguikeji/arkfbp-go/node"
)

func createGraph() *graph.Graph {
	g := graph.New()
	g.Add(&Node2{})
	g.Add(&Node1{})

	nn := &Node3{
		APINode: node.APINode{
			Mode:   "direct",
			Method: "GET",
			URL:    "https://git.intra.longguikeji.com/api/v1/repos/search",
		},
	}
	g.Add(nn)

	g.Add(&Node4{
		IFNode: node.IFNode{
			Expression: func() bool {
				fmt.Println("Node4.Expression")
				return true
			},
			PositiveNext: "Node5",
			NegativeNext: "Node6",
		},
	})

	g.Add(&Node5{})
	g.Add(&Node6{})

	return g
}