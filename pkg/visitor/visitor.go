package visitor

import (
	"github.com/Abraxas-365/gpto/pkg/schema"
	sitter "github.com/smacker/go-tree-sitter"
	"gonum.org/v1/gonum/graph/simple"
)

type FileVisitor interface {
	ParseFile(content []byte, rootNode *sitter.Node, graph *simple.DirectedGraph, nodes map[string]*schema.Node)
}

type Visitor struct {
	Visitor FileVisitor
	Graph   *simple.DirectedGraph
	Nodes   map[string]*schema.Node
}

func New(v FileVisitor) *Visitor {
	g := simple.NewDirectedGraph()
	nodes := make(map[string]*schema.Node)
	return &Visitor{
		Visitor: v,
		Graph:   g,
		Nodes:   nodes,
	}
}
