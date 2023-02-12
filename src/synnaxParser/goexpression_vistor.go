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

func (ge *goExpreesionVisitor) Visit(tree antlr.ParseTree) any {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitChildren(node antlr.RuleNode) any {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitTerminal(node antlr.TerminalNode) any {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) any {
	return ctx.Expression().Accept(ge)
}

func (ge *goExpreesionVisitor) VisitExpression(ctx *parser.ExpressionContext) any {

	if ctx.PrimaryExpr() != nil {
		return ctx.PrimaryExpr().Accept(ge).(*evaluationNode)
	}
	node := newEvaluationNode()
	node.RawString = ctx.GetText()

	if ctx.GetChildCount() == 2 {
		if ctx.PLUS() != nil {
			node.Symbol = PLUS
			node.Operator = unAryAddOperator
			node.RightTypeCheck = isNumber
		}

		if ctx.MINUS() != nil {
			node.Symbol = MINUS
			node.Operator = negateOperator
			node.RightTypeCheck = isNumber
		}

		if ctx.EXCLAMATION() != nil {
			node.Symbol = EXCLAMATION
			node.Operator = invertOperator
			//node.RightTypeCheck = isBool
		}

		if ctx.EN_NOT() != nil {
			node.Symbol = EN_NOT
			node.Operator = invertOperator
			//node.RightTypeCheck = isBool
		}

		if ctx.CARET() != nil {
			node.Symbol = CARET
			node.Operator = bitwiseNotOperator
			//node.RightTypeCheck = isBool
		}

		rightNode := ctx.GetChildOfType(1, nil)
		node.RightOperator = []*evaluationNode{rightNode.Accept(ge).(*evaluationNode)}

	}

	if ctx.GetChildCount() == 3 {
		//oper, _ := ctx.GetChild(0).GetChild(1).(antlr.TerminalNode)
		leftNode := ctx.GetChildOfType(0, nil)
		node.LeftOperator = leftNode.Accept(ge).(*evaluationNode)

		rightNode := ctx.GetChildOfType(2, nil)
		node.RightOperator = []*evaluationNode{rightNode.Accept(ge).(*evaluationNode)}

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
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}
		if ctx.DIV() != nil {
			node.Symbol = DIV
			node.Operator = divideOperator
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}
		if ctx.MOD() != nil {
			node.Symbol = MOD
			node.Operator = modulusOperator
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}

		if ctx.LSHIFT() != nil {
			node.Symbol = LSHIFT
			node.Operator = leftShiftOperator
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}

		if ctx.RSHIFT() != nil {
			node.Symbol = RSHIFT
			node.Operator = rightShiftOperator
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}

		if ctx.AMPERSAND() != nil {
			node.Symbol = AMPERSAND
			node.Operator = bitwiseAndOperator
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}

		if ctx.BIT_CLEAR() != nil {
			node.Symbol = BIT_CLEAR
			node.Operator = bitwiseAndNotOperator
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}

		if ctx.MINUS() != nil {
			node.Symbol = MINUS
			node.Operator = subtractOperator
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}
		if ctx.OR() != nil {
			node.Symbol = OR
			node.Operator = bitwiseOrOperator
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}
		if ctx.CARET() != nil {
			node.Symbol = CARET
			node.Operator = bitwiseXOROperator
			node.LeftTypeCheck = isNumber
			node.RightTypeCheck = isNumber
			node.TypeCheck = nil
		}

		if ctx.EQUALS() != nil {
			node.Symbol = EQUALS
			node.Operator = equalOperator
			node.LeftTypeCheck = nil
			node.RightTypeCheck = nil
			node.TypeCheck = nil
		}

		if ctx.NOT_EQUALS() != nil {
			node.Symbol = NOT_EQUALS
			node.Operator = notEqualOperator
			node.LeftTypeCheck = nil
			node.RightTypeCheck = nil
			node.TypeCheck = nil
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
			//node.LeftTypeCheck = isBool
			//node.RightTypeCheck = isBool
			node.TypeCheck = nil
		}

		if ctx.LOGICAL_OR() != nil || ctx.EN_OR() != nil {
			node.Symbol = LOGICAL_OR
			node.Operator = orOperator
			//node.LeftTypeCheck = isBool
			//node.RightTypeCheck = isBool
			node.TypeCheck = nil
		}
		if ctx.QUESTION() != nil {
			node.Symbol = QUESTION
			node.Operator = conditionalOperator
			//node.LeftTypeCheck = isBool
			node.TypeErrorFormat = leftNode.GetText() + " value is not boolean "
		}
		if ctx.EN_IN() != nil {
			node.Symbol = LOGICAL_IN
			node.Operator = inOperator
			//node.LeftTypeCheck = isBool
			node.TypeErrorFormat = leftNode.GetText() + " value is not boolean "
		}

	}

	return node
}

func (ge *goExpreesionVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) any {
	if ctx.GetChildCount() == 1 {
		return ctx.GetChildOfType(0, nil).Accept(ge).(*evaluationNode)
	}
	if ctx.GetChildCount() == 2 {
		node := newEvaluationNode()
		node.RawString = ctx.GetText()

		node.LeftOperator = ctx.PrimaryExpr().Accept(ge).(*evaluationNode)
		//node.LeftOperator = ctx.PrimaryExpr().Accept(ge).(*evaluationNode)
		if ctx.Index() != nil {
			node.Symbol = INDEX
			node.RightOperator = []*evaluationNode{ctx.Index().Accept(ge).(*evaluationNode)}
			node.Operator = indexOperator
			node.LeftTypeCheck = nil
		}

		if ctx.Slice_() != nil {
			node.Symbol = SLICE
			node.RightOperator = []*evaluationNode{ctx.Slice_().Accept(ge).(*evaluationNode)}
			node.Operator = makeSliceOperator
		}
		if ctx.Arguments() != nil {
			node.Symbol = FUNCCALL
			node.RightOperator = []*evaluationNode{ctx.Arguments().Accept(ge).(*evaluationNode)}
			node.Operator = makeFunction2Operator
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

func (ge *goExpreesionVisitor) VisitOperand(ctx *parser.OperandContext) any {
	if ctx.BasicLit() != nil {
		return ctx.BasicLit().Accept(ge).(*evaluationNode)
	}
	if ctx.OperandName() != nil {
		return ctx.OperandName().Accept(ge).(*evaluationNode)
	}
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(ge).(*evaluationNode)
	}
	if ctx.L_CURLY() != nil && ctx.R_CURLY() != nil {
		if ctx.ExpressionList() != nil {
			return ctx.ExpressionList().Accept(ge).(*evaluationNode)
		} else {
			node := newEvaluationNode()
			node.RawString = ctx.GetText()
			node.Symbol = NIL
			node.Operator = makeLiteralOperator([]interface{}{})
			return node
		}
	}

	return nil
}

func (ge *goExpreesionVisitor) VisitOperandName(ctx *parser.OperandNameContext) any {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = VALUE
	node.RawString = ctx.Identifier().GetText()
	node.Operator = makeParameterOperator(ctx.Identifier().GetText())
	return node
}

func (ge *goExpreesionVisitor) VisitSlice_(ctx *parser.Slice_Context) any {
	return ctx.ExpressionColon().Accept(ge).(*evaluationNode)
}

func (ge *goExpreesionVisitor) VisitExpressionColon(ctx *parser.ExpressionColonContext) any {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()

	allExp := ctx.AllExpression()
	expList := make([]*evaluationNode, 0, len(allExp))
	for _, context := range allExp {
		expList = append(expList, context.Accept(ge).(*evaluationNode))
	}
	node.Symbol = EXPRESSIONCOLON
	node.RightOperator = expList
	node.Operator = rightOperator
	return node
}

func (ge *goExpreesionVisitor) VisitIndex(ctx *parser.IndexContext) any {
	return ctx.Expression().Accept(ge).(*evaluationNode)
}

func (ge *goExpreesionVisitor) VisitBasicLit(ctx *parser.BasicLitContext) any {
	leftNode := ctx.GetChildOfType(0, nil)
	return leftNode.Accept(ge).(*evaluationNode)
}

func (ge *goExpreesionVisitor) VisitNil_lit(ctx *parser.Nil_litContext) any {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = NIL
	node.Operator = makeLiteralOperator(nil)
	return node
}

func (ge *goExpreesionVisitor) VisitEn_bool(ctx *parser.En_boolContext) any {
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

func (ge *goExpreesionVisitor) VisitFloat_lit(ctx *parser.Float_litContext) any {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = FLOATE64
	if ctx.FLOAT_LIT() != nil {
		num, _ := strconv.ParseFloat(ctx.FLOAT_LIT().GetText(), 64)
		node.Operator = makeLiteralOperator(num)
	}

	return node
}

func (ge *goExpreesionVisitor) VisitInteger(ctx *parser.IntegerContext) any {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = FLOATE64
	node.RawString = ctx.GetText()
	numStr := strings.Replace(ctx.GetText(), "_", "", -1)
	if ctx.OCTAL_LIT() != nil {
		num, _ := strconv.ParseInt(numStr[2:len(numStr)], 8, 64)
		node.Operator = makeLiteralOperator(num)
	}
	if ctx.DECIMAL_LIT() != nil {
		num, _ := strconv.ParseInt(numStr, 10, 64)
		node.Operator = makeLiteralOperator(num)
	}
	if ctx.HEX_LIT() != nil {
		num, _ := strconv.ParseInt(numStr[2:len(numStr)], 16, 64)
		node.Operator = makeLiteralOperator(num)
	}
	if ctx.BINARY_LIT() != nil {
		num, _ := strconv.ParseInt(numStr[2:len(numStr)], 2, 64)
		node.Operator = makeLiteralOperator(num)
	}

	return node
}

func (ge *goExpreesionVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) any {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()

	allExp := ctx.AllExpression()
	expList := make([]*evaluationNode, 0, len(allExp))
	for _, context := range allExp {
		expList = append(expList, context.Accept(ge).(*evaluationNode))
	}

	node.Symbol = EXPRESSIONLIST
	node.RightOperator = expList
	node.Operator = rightOperator
	return node
}

func (ge *goExpreesionVisitor) VisitArguments(ctx *parser.ArgumentsContext) any {
	if ctx.ExpressionList() == nil {
		node := newEvaluationNode()
		node.RawString = ctx.GetText()
		node.Symbol = NIL
		node.Operator = makeLiteralOperator([]interface{}{})
		return node
	}
	return ctx.ExpressionList().Accept(ge).(*evaluationNode)
}

func (ge *goExpreesionVisitor) VisitEos(ctx *parser.EosContext) any {
	//TODO implement me
	panic("implement me")
}

func (ge *goExpreesionVisitor) VisitIdentifier(ctx *parser.IdentifierContext) any {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = LITERAL
	if ctx.IDENTIFIER() != nil {
		node.Operator = makeParameterOperator(ctx.IDENTIFIER().GetText())
	}
	return node
}

func (ge *goExpreesionVisitor) VisitString_(ctx *parser.String_Context) any {
	node := newEvaluationNode()
	node.RawString = ctx.GetText()
	node.Symbol = STRING
	if ctx.INTERPRETED_STRING_LIT() != nil {
		unquote, err := strconv.Unquote(ctx.INTERPRETED_STRING_LIT().GetText())
		if err != nil {
			node.Operator = makeLiteralOperator(ctx.INTERPRETED_STRING_LIT().GetText()[1 : len(ctx.INTERPRETED_STRING_LIT().GetText())-1])
		} else {
			node.Operator = makeLiteralOperator(unquote)
		}
		//node.Operator = makeLiteralOperator(ctx.INTERPRETED_STRING_LIT().GetText()[1 : len(ctx.INTERPRETED_STRING_LIT().GetText())-1])
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
