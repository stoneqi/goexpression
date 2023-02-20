// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // goexpression
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by goexpressionParser.
type goexpressionVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by goexpressionParser#expressionStmt.
	VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

	// Visit a parse tree produced by goexpressionParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by goexpressionParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by goexpressionParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by goexpressionParser#operandName.
	VisitOperandName(ctx *OperandNameContext) interface{}

	// Visit a parse tree produced by goexpressionParser#slice_.
	VisitSlice_(ctx *Slice_Context) interface{}

	// Visit a parse tree produced by goexpressionParser#expressionColon.
	VisitExpressionColon(ctx *ExpressionColonContext) interface{}

	// Visit a parse tree produced by goexpressionParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by goexpressionParser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}

	// Visit a parse tree produced by goexpressionParser#nil_lit.
	VisitNil_lit(ctx *Nil_litContext) interface{}

	// Visit a parse tree produced by goexpressionParser#en_bool.
	VisitEn_bool(ctx *En_boolContext) interface{}

	// Visit a parse tree produced by goexpressionParser#float_lit.
	VisitFloat_lit(ctx *Float_litContext) interface{}

	// Visit a parse tree produced by goexpressionParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by goexpressionParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by goexpressionParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by goexpressionParser#eos.
	VisitEos(ctx *EosContext) interface{}

	// Visit a parse tree produced by goexpressionParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by goexpressionParser#string_.
	VisitString_(ctx *String_Context) interface{}
}
