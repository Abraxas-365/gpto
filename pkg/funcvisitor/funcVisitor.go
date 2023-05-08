package funcvisitor

import (
	"fmt"

	"github.com/Abraxas-365/gpto/pkg/funcnode"
	sitter "github.com/smacker/go-tree-sitter"
	"gonum.org/v1/gonum/graph/simple"
)

type FuncVisitor struct {
	Graph       *simple.DirectedGraph
	Nodes       map[string]*funcnode.FuncNode
	PackageName string
	FuncFound   func(functionName string)
}

func extractFunctions(node *sitter.Node, content []byte, v *FuncVisitor) {
	if node.Type() == "function_declaration" {
		nameNode := node.ChildByFieldName("name")
		name := string(content[nameNode.StartByte():nameNode.EndByte()])
		qualifiedFuncName := v.PackageName + "." + name
		fmt.Println(qualifiedFuncName)

		functionBody := string(content[node.StartByte():node.EndByte()])

		if _, ok := v.Nodes[qualifiedFuncName]; !ok {
			newNode := v.Graph.NewNode()
			v.Nodes[qualifiedFuncName] = &funcnode.FuncNode{Node: newNode, Name: qualifiedFuncName, Body: functionBody}
			v.Graph.AddNode(v.Nodes[qualifiedFuncName])
		}

		// Call the FuncFound function if it's set
		if v.FuncFound != nil {
			v.FuncFound(qualifiedFuncName)
		}
	}

	// Recurse into child nodes
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		extractFunctions(child, content, v)
	}
}
func (v *FuncVisitor) ParseGoFile(content []byte, rootNode *sitter.Node) {
	// Extract and process function data
	v.PackageName = extractPackageName(rootNode, content)
	extractFunctions(rootNode, content, v)
}

func extractPackageName(node *sitter.Node, content []byte) string {
	if node.Type() == "package_identifier" {
		name := string(content[node.StartByte():node.EndByte()])
		return name
	}

	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		packageName := extractPackageName(child, content)
		if packageName != "" {
			return packageName
		}
	}

	return ""
}
