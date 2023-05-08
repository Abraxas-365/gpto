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

		functionCalls := extractFunctionCalls(node, content)
		for _, calledFunc := range functionCalls {
			calledFuncName := calledFunc
			fmt.Println("qualifiedFuncName:", qualifiedFuncName, "call:", calledFuncName)
			if calledNode, ok := v.Nodes[calledFuncName]; ok {
				v.Graph.SetEdge(simple.Edge{F: v.Nodes[qualifiedFuncName], T: calledNode})
			}
		}
	}

	// Recurse into child nodes
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		extractFunctions(child, content, v)
	}
}

func extractFunctionCalls(node *sitter.Node, content []byte) []string {
	var calls []string

	if node.Type() == "call_expression" {
		for i := 0; i < int(node.ChildCount()); i++ {
			child := node.Child(i)
			fmt.Println("child", i, child.Type())
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

func (v *FuncVisitor) ParseFile(content []byte, rootNode *sitter.Node) {
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
