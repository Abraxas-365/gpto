package funcnode

import "gonum.org/v1/gonum/graph"

type FuncNode struct {
	graph.Node
	Name    string
	Body    string
	Summary string
}

func (n *FuncNode) ID() int64 { return n.Node.ID() }
