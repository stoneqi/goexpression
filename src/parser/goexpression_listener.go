// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // goexpression
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// goexpressionListener is a complete listener for a parse tree produced by goexpressionParser.
type goexpressionListener interface {
	antlr.ParseTreeListener

	// EnterExpressionStmt is called when entering the expressionStmt production.
	EnterExpressionStmt(c *ExpressionStmtContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterOperandName is called when entering the operandName production.
	EnterOperandName(c *OperandNameContext)

	// EnterSlice_ is called when entering the slice_ production.
	EnterSlice_(c *Slice_Context)

	// EnterExpressionColon is called when entering the expressionColon production.
	EnterExpressionColon(c *ExpressionColonContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterBasicLit is called when entering the basicLit production.
	EnterBasicLit(c *BasicLitContext)

	// EnterNil_lit is called when entering the nil_lit production.
	EnterNil_lit(c *Nil_litContext)

	// EnterEn_bool is called when entering the en_bool production.
	EnterEn_bool(c *En_boolContext)

	// EnterFloat_lit is called when entering the float_lit production.
	EnterFloat_lit(c *Float_litContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// ExitExpressionStmt is called when exiting the expressionStmt production.
	ExitExpressionStmt(c *ExpressionStmtContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitOperandName is called when exiting the operandName production.
	ExitOperandName(c *OperandNameContext)

	// ExitSlice_ is called when exiting the slice_ production.
	ExitSlice_(c *Slice_Context)

	// ExitExpressionColon is called when exiting the expressionColon production.
	ExitExpressionColon(c *ExpressionColonContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitBasicLit is called when exiting the basicLit production.
	ExitBasicLit(c *BasicLitContext)

	// ExitNil_lit is called when exiting the nil_lit production.
	ExitNil_lit(c *Nil_litContext)

	// ExitEn_bool is called when exiting the en_bool production.
	ExitEn_bool(c *En_boolContext)

	// ExitFloat_lit is called when exiting the float_lit production.
	ExitFloat_lit(c *Float_litContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)
}
