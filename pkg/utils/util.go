package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Abraxas-365/gpto/pkg/schema"
	"github.com/Abraxas-365/gpto/pkg/visitor"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
	"gonum.org/v1/gonum/graph/simple"
)

func NewIndexer(directory string, v visitor.Visitor, summaryFuntion func(string, []string) string) ([]schema.Node, error) {
	parser := sitter.NewParser()
	parser.SetLanguage(golang.GetLanguage())

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

		v.Visitor.ParseFile(content, rootNode, v.Graph, v.Nodes)
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return nil, err
	}

	nodeList := []schema.Node{}

	for _, node := range v.Nodes {
		if len(node.Body) == 0 {
			continue
		}
		if node.Summary == "" {
			getNodeSummary(node, v.Graph, summaryFuntion)
		}
		nodeList = append(nodeList, *node)
	}

	return nodeList, nil
}

func getNodeSummary(node *schema.Node, g *simple.DirectedGraph, summaryFuntion func(string, []string) string) string {
	if node.Summary != "" {
		return node.Summary
	}

	calledSummaries := []string{}

	fromEdges := g.From(node.ID())
	for fromEdges.Next() {
		target := fromEdges.Node().(*schema.Node)
		if len(target.Body) > 0 {
			calledSummaries = append(calledSummaries, getNodeSummary(target, g, summaryFuntion))
		}

	}

	node.Summary = summaryFuntion(node.Body, calledSummaries)
	return node.Summary
}
