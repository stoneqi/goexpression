// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // goexpression
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BasegoexpressionListener is a complete listener for a parse tree produced by goexpressionParser.
type BasegoexpressionListener struct{}

var _ goexpressionListener = &BasegoexpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegoexpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegoexpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegoexpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegoexpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpressionStmt is called when production expressionStmt is entered.
func (s *BasegoexpressionListener) EnterExpressionStmt(ctx *ExpressionStmtContext) {}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *BasegoexpressionListener) ExitExpressionStmt(ctx *ExpressionStmtContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasegoexpressionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasegoexpressionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BasegoexpressionListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BasegoexpressionListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterOperand is called when production operand is entered.
func (s *BasegoexpressionListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BasegoexpressionListener) ExitOperand(ctx *OperandContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BasegoexpressionListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BasegoexpressionListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterSlice_ is called when production slice_ is entered.
func (s *BasegoexpressionListener) EnterSlice_(ctx *Slice_Context) {}

// ExitSlice_ is called when production slice_ is exited.
func (s *BasegoexpressionListener) ExitSlice_(ctx *Slice_Context) {}

// EnterIndex is called when production index is entered.
func (s *BasegoexpressionListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BasegoexpressionListener) ExitIndex(ctx *IndexContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *BasegoexpressionListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *BasegoexpressionListener) ExitBasicLit(ctx *BasicLitContext) {}

// EnterNil_lit is called when production nil_lit is entered.
func (s *BasegoexpressionListener) EnterNil_lit(ctx *Nil_litContext) {}

// ExitNil_lit is called when production nil_lit is exited.
func (s *BasegoexpressionListener) ExitNil_lit(ctx *Nil_litContext) {}

// EnterEn_bool is called when production en_bool is entered.
func (s *BasegoexpressionListener) EnterEn_bool(ctx *En_boolContext) {}

// ExitEn_bool is called when production en_bool is exited.
func (s *BasegoexpressionListener) ExitEn_bool(ctx *En_boolContext) {}

// EnterFloat_lit is called when production float_lit is entered.
func (s *BasegoexpressionListener) EnterFloat_lit(ctx *Float_litContext) {}

// ExitFloat_lit is called when production float_lit is exited.
func (s *BasegoexpressionListener) ExitFloat_lit(ctx *Float_litContext) {}

// EnterInteger is called when production integer is entered.
func (s *BasegoexpressionListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasegoexpressionListener) ExitInteger(ctx *IntegerContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasegoexpressionListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasegoexpressionListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BasegoexpressionListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasegoexpressionListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterEos is called when production eos is entered.
func (s *BasegoexpressionListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BasegoexpressionListener) ExitEos(ctx *EosContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasegoexpressionListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasegoexpressionListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BasegoexpressionListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BasegoexpressionListener) ExitString_(ctx *String_Context) {}
