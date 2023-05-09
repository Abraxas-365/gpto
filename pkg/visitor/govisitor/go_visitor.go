package govisitor

import (
	"fmt"

	"github.com/Abraxas-365/gpto/pkg/schema"
	"github.com/Abraxas-365/gpto/pkg/visitor"
	sitter "github.com/smacker/go-tree-sitter"
	"gonum.org/v1/gonum/graph/simple"
)

type GoVisitor struct {
	PackageName string
	FuncFound   func(functionName string)
}

var _ visitor.FileVisitor = (*GoVisitor)(nil)

func (v *GoVisitor) ParseFile(content []byte, rootNode *sitter.Node, graph *simple.DirectedGraph, nodes map[string]*schema.Node) {
	// Extract and process function data
	v.PackageName = extractPackageName(rootNode, content)
	extractFunctions(rootNode, content, v, graph, nodes)
}

func extractFunctions(tsnode *sitter.Node, content []byte, v *GoVisitor, graph *simple.DirectedGraph, nodes map[string]*schema.Node) {
	if tsnode.Type() == "function_declaration" || tsnode.Type() == "method_declaration" {
		nameNode := tsnode.ChildByFieldName("name")
		if tsnode.Type() == "method_declaration" {
			nameNode = tsnode.ChildByFieldName("name")
		}

		name := string(content[nameNode.StartByte():nameNode.EndByte()])
		qualifiedFuncName := v.PackageName + "." + name
		fmt.Println(qualifiedFuncName)

		functionBody := string(content[tsnode.StartByte():tsnode.EndByte()])
		fmt.Println("Body: ", functionBody)

		if _, ok := nodes[qualifiedFuncName]; !ok {
			newNode := graph.NewNode()
			nodes[qualifiedFuncName] = &schema.Node{Node: newNode, Name: qualifiedFuncName, Body: functionBody}
			graph.AddNode(nodes[qualifiedFuncName])
		}

		// Call the FuncFound function if it's set
		if v.FuncFound != nil {
			v.FuncFound(qualifiedFuncName)
		}

		functionCalls := extractFunctionCalls(tsnode, content)
		for _, calledFunc := range functionCalls {
			calledFuncName := calledFunc
			if calledNode, ok := nodes[calledFuncName]; ok {
				graph.SetEdge(simple.Edge{F: nodes[qualifiedFuncName], T: calledNode})
			}
		}
	}

	// Recurse into child nodes
	for i := 0; i < int(tsnode.ChildCount()); i++ {
		child := tsnode.Child(i)
		extractFunctions(child, content, v, graph, nodes)
	}
}

func extractFunctionCalls(node *sitter.Node, content []byte) []string {
	var calls []string

	if node.Type() == "call_expression" {
		for i := 0; i < int(node.ChildCount()); i++ {
			child := node.Child(i)
			if child.Type() == "selector_expression" {
				callName := string(content[child.StartByte():child.EndByte()])
				calls = append(calls, callName)
			}
		}
	}

	// Recurse into child nodes
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		childCalls := extractFunctionCalls(child, content)
		calls = append(calls, childCalls...)
	}

	return calls
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
