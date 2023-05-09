package schema

import "gonum.org/v1/gonum/graph"

type Node struct {
	graph.Node
	Name    string
	Body    string
	Summary string
}

func (n *Node) ID() int64 { return n.Node.ID() }
