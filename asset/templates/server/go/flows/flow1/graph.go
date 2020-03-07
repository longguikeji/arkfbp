package flow1

import (
	"github.com/longguikeji/arkfbp-go/graph"
)

func createGraph() *graph.Graph {
	g := graph.New()
	g.Add(&Node1{})
	g.Add(&Node2{})
	g.Add(&Node3{})

	return g
}
