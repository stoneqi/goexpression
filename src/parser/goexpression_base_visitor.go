// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // goexpression
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BasegoexpressionVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasegoexpressionVisitor) VisitExpressionStmt(ctx *ExpressionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitOperandName(ctx *OperandNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitSlice_(ctx *Slice_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitExpressionColon(ctx *ExpressionColonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitBasicLit(ctx *BasicLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitNil_lit(ctx *Nil_litContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitEn_bool(ctx *En_boolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitFloat_lit(ctx *Float_litContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegoexpressionVisitor) VisitString_(ctx *String_Context) interface{} {
	return v.VisitChildren(ctx)
}
