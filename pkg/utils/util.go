package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Abraxas-365/gpto/pkg/funcnode"
	"github.com/Abraxas-365/gpto/pkg/funcvisitor"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
	"gonum.org/v1/gonum/graph/simple"
)

func NewFunctionIndexer(directory string, summaryFuntion func(string, []string) string) ([]funcnode.FuncNode, error) {
	parser := sitter.NewParser()
	parser.SetLanguage(golang.GetLanguage())

	g := simple.NewDirectedGraph()
	nodes := make(map[string]*funcnode.FuncNode)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		tree := parser.Parse(nil, content)
		rootNode := tree.RootNode()

		visitor := &funcvisitor.FuncVisitor{
			Graph: g,
			Nodes: nodes,
		}

		visitor.ParseFile(content, rootNode)
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return nil, err
	}

	nodeList := []funcnode.FuncNode{}

	for _, node := range nodes {
		if len(node.Body) == 0 {
			continue
		}
		if node.Summary == "" {
			getNodeSummary(node, g, summaryFuntion)
		}
		nodeList = append(nodeList, *node)
		fmt.Printf("Function: %s\nCombined Body:\n%s\n\n", node.Name, node.Body)
	}

	return nodeList, nil
}

func getNodeSummary(node *funcnode.FuncNode, g *simple.DirectedGraph, summaryFuntion func(string, []string) string) string {
	if node.Summary != "" {
		return node.Summary
	}

	calledSummaries := []string{}

	fromEdges := g.From(node.ID())
	for fromEdges.Next() {
		target := fromEdges.Node().(*funcnode.FuncNode)
		if len(target.Body) > 0 {
			calledSummaries = append(calledSummaries, getNodeSummary(target, g, summaryFuntion))
		}

	}

	node.Summary = summaryFuntion(node.Body, calledSummaries)
	return node.Summary
}
