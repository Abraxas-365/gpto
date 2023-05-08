package funcvisitor

import (
	"github.com/Abraxas-365/gpto/pkg/funcnode"
	sitter "github.com/smacker/go-tree-sitter"
	"gonum.org/v1/gonum/graph/simple"
)

type FuncVisitor interface {
	ParseFile(content []byte, rootNode *sitter.Node, graph *simple.DirectedGraph, nodes map[string]*funcnode.FuncNode)
}

type Visitor struct {
	Visitor FuncVisitor
	Graph   *simple.DirectedGraph
	Nodes   map[string]*funcnode.FuncNode
}

func New(v FuncVisitor) *Visitor {
	g := simple.NewDirectedGraph()
	nodes := make(map[string]*funcnode.FuncNode)
	return &Visitor{
		Visitor: v,
		Graph:   g,
		Nodes:   nodes,
	}
}
