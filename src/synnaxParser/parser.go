// Package main provides ...
package parserSecond

import (
	"github.com/stoneqi/goexpression/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type goExpressionListen struct {
}

func (goExpressionListen) VisitTerminal(node antlr.TerminalNode) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) VisitErrorNode(node antlr.ErrorNode) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterExpressionStmt(c *parser.ExpressionStmtContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterExpression(c *parser.ExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterPrimaryExpr(c *parser.PrimaryExprContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterOperand(c *parser.OperandContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterOperandName(c *parser.OperandNameContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterSlice_(c *parser.Slice_Context) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterIndex(c *parser.IndexContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterBasicLit(c *parser.BasicLitContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterNil_lit(c *parser.Nil_litContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterEn_bool(c *parser.En_boolContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterFloat_lit(c *parser.Float_litContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterInteger(c *parser.IntegerContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterExpressionList(c *parser.ExpressionListContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterArguments(c *parser.ArgumentsContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterEos(c *parser.EosContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterIdentifier(c *parser.IdentifierContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterString_(c *parser.String_Context) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitExpressionStmt(c *parser.ExpressionStmtContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitExpression(c *parser.ExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitPrimaryExpr(c *parser.PrimaryExprContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitOperand(c *parser.OperandContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitOperandName(c *parser.OperandNameContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitSlice_(c *parser.Slice_Context) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitIndex(c *parser.IndexContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitBasicLit(c *parser.BasicLitContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitNil_lit(c *parser.Nil_litContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitEn_bool(c *parser.En_boolContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitFloat_lit(c *parser.Float_litContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitInteger(c *parser.IntegerContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitExpressionList(c *parser.ExpressionListContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitArguments(c *parser.ArgumentsContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitEos(c *parser.EosContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitIdentifier(c *parser.IdentifierContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitString_(c *parser.String_Context) {
	//TODO implement me
	panic("implement me")
}

func ParserString(expression string) (string, error) {
	lexer := parser.NewgoexpressionLexer(antlr.NewInputStream(expression))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewgoexpressionParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.ExpressionStmt()
	return tree.GetText(), nil
}
