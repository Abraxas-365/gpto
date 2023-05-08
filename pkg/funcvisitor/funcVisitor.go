package funcvisitor

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type FuncVisitor interface {
	ParseFile(content []byte, rootNode *sitter.Node)
}
