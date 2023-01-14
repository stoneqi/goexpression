// Package main provides ...
package parserSecond

import (
	"github.com/stoneqi/goexpression/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func ParserString(expression string) (string, error) {
	lexer := parser.NewgoexpressionLexer(antlr.NewInputStream(expression))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewgoexpressionParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.ExpressionStmt()
	return tree.GetText(), nil
}
