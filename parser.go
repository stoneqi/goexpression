// Package main provides ...
package goexpression

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser2 "github.com/stoneqi/goexpression/parser"
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

func (goExpressionListen) EnterExpressionStmt(c *parser2.ExpressionStmtContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterExpression(c *parser2.ExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterPrimaryExpr(c *parser2.PrimaryExprContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterOperand(c *parser2.OperandContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterOperandName(c *parser2.OperandNameContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterSlice_(c *parser2.Slice_Context) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterIndex(c *parser2.IndexContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterBasicLit(c *parser2.BasicLitContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterNil_lit(c *parser2.Nil_litContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterEn_bool(c *parser2.En_boolContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterFloat_lit(c *parser2.Float_litContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterInteger(c *parser2.IntegerContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterExpressionList(c *parser2.ExpressionListContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterArguments(c *parser2.ArgumentsContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterEos(c *parser2.EosContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterIdentifier(c *parser2.IdentifierContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) EnterString_(c *parser2.String_Context) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitExpressionStmt(c *parser2.ExpressionStmtContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitExpression(c *parser2.ExpressionContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitPrimaryExpr(c *parser2.PrimaryExprContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitOperand(c *parser2.OperandContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitOperandName(c *parser2.OperandNameContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitSlice_(c *parser2.Slice_Context) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitIndex(c *parser2.IndexContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitBasicLit(c *parser2.BasicLitContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitNil_lit(c *parser2.Nil_litContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitEn_bool(c *parser2.En_boolContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitFloat_lit(c *parser2.Float_litContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitInteger(c *parser2.IntegerContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitExpressionList(c *parser2.ExpressionListContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitArguments(c *parser2.ArgumentsContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitEos(c *parser2.EosContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitIdentifier(c *parser2.IdentifierContext) {
	//TODO implement me
	panic("implement me")
}

func (goExpressionListen) ExitString_(c *parser2.String_Context) {
	//TODO implement me
	panic("implement me")
}

func ParserString(expression string) (string, error) {
	lexer := parser2.NewgoexpressionLexer(antlr.NewInputStream(expression))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser2.NewgoexpressionParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.ExpressionStmt()
	return tree.GetText(), nil
}
