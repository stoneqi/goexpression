package parserSecond

import (
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/stoneqi/goexpression/src/parser"
)

type goExpreesionVisitor struct {
	root *evaluationNode
}

func newGoExpreesionVisitor() *goExpreesionVisitor {
	return &goExpreesionVisitor{}
}

func (ge *goExpreesionVisitor) Visit(tree antlr.ParseTree) interface{} {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	return ctx.Expression().Accept(ge)
}

func (ge *goExpreesionVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {

	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	if ctx.GetChildCount() == 1 {
		node.Symbol = OXXX
		node.Operator = leftOperator

		leftNode := ctx.GetChildOfType(0, nil)
		node.LeftOperator = leftNode.Accept(ge).(*evaluationNode)
	}

	if ctx.GetChildCount() == 2 {
		if ctx.PLUS() != nil {
			node.Symbol = PLUS
			node.Operator = unAryAddOperator
			node.RightTypeCheck = isFloat64
		}

		if ctx.MINUS() != nil {
			node.Symbol = MINUS
			node.Operator = negateOperator
			node.RightTypeCheck = isFloat64
		}

		if ctx.EXCLAMATION() != nil {
			node.Symbol = EXCLAMATION
			node.Operator = invertOperator
			node.RightTypeCheck = isBool
		}

		if ctx.EN_NOT() != nil {
			node.Symbol = EN_NOT
			node.Operator = invertOperator
			node.RightTypeCheck = isBool
		}

		if ctx.CARET() != nil {
			node.Symbol = CARET
			node.Operator = bitwiseNotOperator
			node.RightTypeCheck = isBool
		}

		rightNode := ctx.GetChildOfType(1, nil)
		node.RightOperator = rightNode.Accept(ge).(*evaluationNode)

	}

	if ctx.GetChildCount() == 3 {
		//oper, _ := ctx.GetChild(0).GetChild(1).(antlr.TerminalNode)
		leftNode := ctx.GetChildOfType(0, nil)
		node.LeftOperator = leftNode.Accept(ge).(*evaluationNode)
		rightNode := ctx.GetChildOfType(2, nil)
		node.RightOperator = rightNode.Accept(ge).(*evaluationNode)

		if ctx.PLUS() != nil {
			node.Symbol = PLUS
			node.Operator = addOperator
			node.LeftTypeCheck = nil
			node.RightTypeCheck = nil
			node.TypeCheck = additionTypeCheck
		}

		if ctx.STAR() != nil {
			node.Symbol = STAR
			node.Operator = multiplyOperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}
		if ctx.DIV() != nil {
			node.Symbol = DIV
			node.Operator = divideOperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}
		if ctx.MOD() != nil {
			node.Symbol = MOD
			node.Operator = modulusOperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}

		if ctx.LSHIFT() != nil {
			node.Symbol = LSHIFT
			node.Operator = leftShiftOperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}

		if ctx.RSHIFT() != nil {
			node.Symbol = RSHIFT
			node.Operator = rightShiftOperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}

		if ctx.AMPERSAND() != nil {
			node.Symbol = AMPERSAND
			node.Operator = bitwiseAndOperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}

		if ctx.BIT_CLEAR() != nil {
			node.Symbol = BIT_CLEAR
			node.Operator = bitwiseAndNotOperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}

		if ctx.MINUS() != nil {
			node.Symbol = MINUS
			node.Operator = subtractOperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}
		if ctx.OR() != nil {
			node.Symbol = OR
			node.Operator = bitwiseOrOperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}
		if ctx.CARET() != nil {
			node.Symbol = CARET
			node.Operator = bitwiseXOROperator
			node.LeftTypeCheck = isFloat64
			node.RightTypeCheck = isFloat64
			node.TypeCheck = nil
		}

		if ctx.EQUALS() != nil {
			node.Symbol = EQUALS
			node.Operator = equalOperator
			node.LeftTypeCheck = nil
			node.RightTypeCheck = nil
			node.TypeCheck = comparatorTypeCheck
		}

		if ctx.NOT_EQUALS() != nil {
			node.Symbol = NOT_EQUALS
			node.Operator = notEqualOperator
			node.LeftTypeCheck = nil
			node.RightTypeCheck = nil
			node.TypeCheck = comparatorTypeCheck
		}
		if ctx.LESS() != nil {
			node.Symbol = LESS
			node.Operator = ltOperator
			node.LeftTypeCheck = nil
			node.RightTypeCheck = nil
			node.TypeCheck = comparatorTypeCheck
		}
		if ctx.LESS_OR_EQUALS() != nil {
			node.Symbol = LESS_OR_EQUALS
			node.Operator = lteOperator
			node.LeftTypeCheck = nil
			node.RightTypeCheck = nil
			node.TypeCheck = comparatorTypeCheck
		}
		if ctx.GREATER() != nil {
			node.Symbol = GREATER
			node.Operator = gtOperator
			node.LeftTypeCheck = nil
			node.RightTypeCheck = nil
			node.TypeCheck = comparatorTypeCheck
		}
		if ctx.GREATER_OR_EQUALS() != nil {
			node.Symbol = GREATER_OR_EQUALS
			node.Operator = gteOperator
			node.LeftTypeCheck = nil
			node.RightTypeCheck = nil
			node.TypeCheck = comparatorTypeCheck
		}

		if ctx.LOGICAL_AND() != nil || ctx.EN_AND() != nil {
			node.Symbol = LOGICAL_AND
			node.Operator = andOperator
			node.LeftTypeCheck = isBool
			node.RightTypeCheck = isBool
			node.TypeCheck = nil
		}

		if ctx.LOGICAL_OR() != nil || ctx.EN_OR() != nil {
			node.Symbol = LOGICAL_OR
			node.Operator = orOperator
			node.LeftTypeCheck = isBool
			node.RightTypeCheck = isBool
			node.TypeCheck = nil
		}

	}

	return node
}

func (ge *goExpreesionVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	if ctx.GetChildCount() == 1 {
		return ctx.GetChildOfType(0, nil).Accept(ge).(*evaluationNode)
	}
	if ctx.GetChildCount() == 2 {
		node := newEvaluationNode()
		node.RawString = ctx.GetText()
		node.LeftOperator = ctx.PrimaryExpr().Accept(ge).(*evaluationNode)
		if ctx.Index() != nil {
			node.RightOperator = ctx.Index().Accept(ge).(*evaluationNode)
			node.Operator = indexOperator
			node.LeftTypeCheck = nil
		}

		if ctx.Slice_() != nil {
			node.RightOperator = ctx.Slice_().Accept(ge).(*evaluationNode)
		}
		if ctx.Arguments() != nil {
			node.RightOperator = ctx.Arguments().Accept(ge).(*evaluationNode)
			node.Operator = makeFunctionOperator
		}
		return node
	}
	if ctx.GetChildCount() == 3 {
		node := newEvaluationNode()
		node.RawString = ctx.GetText()
		node.LeftOperator = ctx.PrimaryExpr().Accept(ge).(*evaluationNode)
		node.Operator = makeAccessOperator(ctx.IDENTIFIER().GetText())
		return node
	}

	return nil
}

func (ge *goExpreesionVisitor) VisitOperand(ctx *parser.OperandContext) interface{} {
	if ctx.BasicLit() != nil {
		return ctx.BasicLit().Accept(ge).(*evaluationNode)
	}
	if ctx.OperandName() != nil {
		return ctx.OperandName().Accept(ge).(*evaluationNode)
	}
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(ge).(*evaluationNode)
	}
	return nil
}

func (ge *goExpreesionVisitor) VisitOperandName(ctx *parser.OperandNameContext) interface{} {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = LITERAL
	node.RawString = ctx.Identifier().GetText()
	node.Operator = makeParameterOperator(ctx.Identifier().GetText())
	return node
}

func (ge *goExpreesionVisitor) VisitSlice_(ctx *parser.Slice_Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitIndex(ctx *parser.IndexContext) interface{} {
	return ctx.Expression().Accept(ge).(*evaluationNode)
}

func (ge *goExpreesionVisitor) VisitBasicLit(ctx *parser.BasicLitContext) interface{} {
	leftNode := ctx.GetChildOfType(0, nil)
	return leftNode.Accept(ge).(*evaluationNode)
}

func (ge *goExpreesionVisitor) VisitNil_lit(ctx *parser.Nil_litContext) interface{} {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = NIL
	node.Operator = makeLiteralOperator(nil)
	return node
}

func (ge *goExpreesionVisitor) VisitEn_bool(ctx *parser.En_boolContext) interface{} {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = BOOL
	if ctx.EN_TRUE() != nil {
		node.Operator = makeLiteralOperator(true)
	}
	if ctx.EN_FALSE() != nil {
		node.Operator = makeLiteralOperator(false)
	}
	return node
}

func (ge *goExpreesionVisitor) VisitFloat_lit(ctx *parser.Float_litContext) interface{} {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = FLOATE64
	if ctx.FLOAT_LIT() != nil {
		num, _ := strconv.ParseFloat(ctx.FLOAT_LIT().GetText(), 64)
		node.Operator = makeLiteralOperator(num)
	}

	return node
}

func (ge *goExpreesionVisitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = FLOATE64
	node.RawString = ctx.GetText()
	numStr := strings.Replace(ctx.GetText(), "_", "", -1)
	if ctx.OCTAL_LIT() != nil {
		num, _ := strconv.ParseInt(numStr[2:len(numStr)], 8, 64)
		node.Operator = makeLiteralOperator(float64(num))
	}
	if ctx.DECIMAL_LIT() != nil {
		num, _ := strconv.ParseInt(numStr, 10, 64)
		node.Operator = makeLiteralOperator(float64(num))
	}
	if ctx.HEX_LIT() != nil {
		num, _ := strconv.ParseInt(numStr[2:len(numStr)], 16, 64)
		node.Operator = makeLiteralOperator(float64(num))
	}
	if ctx.BINARY_LIT() != nil {
		num, _ := strconv.ParseInt(numStr[2:len(numStr)], 2, 64)
		node.Operator = makeLiteralOperator(float64(num))
	}

	return node
}

func (ge *goExpreesionVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()

	allExp := ctx.AllExpression()
	expList := make([]*evaluationNode, 0, len(allExp))
	for _, context := range allExp {
		expList = append(expList, context.Accept(ge).(*evaluationNode))
	}

	node.Symbol = FUNCPARAMS
	node.Operator = makeExpressionListOperator(expList)
	return node
}

func (ge *goExpreesionVisitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	return ctx.ExpressionList().Accept(ge).(*evaluationNode)
}

func (ge *goExpreesionVisitor) VisitEos(ctx *parser.EosContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = LITERAL
	if ctx.IDENTIFIER() != nil {
		node.Operator = makeParameterOperator(ctx.IDENTIFIER().GetText())
	}
	return node
}

func (ge *goExpreesionVisitor) VisitString_(ctx *parser.String_Context) interface{} {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = STRING
	if ctx.INTERPRETED_STRING_LIT() != nil {
		node.Operator = makeLiteralOperator(ctx.INTERPRETED_STRING_LIT().GetText())
	}
	return node
}

func VisitorParserString(expression string) (*evaluationNode, error) {
	lexer := parser.NewgoexpressionLexer(antlr.NewInputStream(expression))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewgoexpressionParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	visitor := newGoExpreesionVisitor()
	tree := p.ExpressionStmt()
	res := tree.Accept(visitor)

	node := res.(*evaluationNode)
	return node, nil
}
