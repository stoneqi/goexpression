package parserSecond

import (
	"errors"
	"fmt"
	"github.com/stoneqi/goexpression"
	"sync"
)

const shortCircuitHolder int = -1

type EvaluableExpression struct {
	singleExpr   *evaluationNode
	singleString string
	ChecksTypes  bool
	recordStep   []string
	IsDebug      bool
	expr         sync.Map
	exprString   sync.Map
}

func NewEvaluableExpression(stage *evaluationNode) *EvaluableExpression {
	return &EvaluableExpression{
		singleExpr: stage}
}

func (ee *EvaluableExpression) AddExpr(key any, expr string) (err error) {
	defer func() {
		if errRec := recover(); errRec != nil {
			err = errors.New(fmt.Sprintf("panic: %+v", errRec))
			return
		}
		err = errors.New("panic: no err info")
		return
	}()
	stageTemp, err := VisitorParserString(expr)
	if err != nil {
		return err
	}
	stageTemp = elideLiterals(stageTemp)
	ee.expr.Store(key, stageTemp)
	ee.exprString.Store(key, expr)
	return nil
}

func (ee *EvaluableExpression) EvalAllExpr(parameters goexpression.Parameters) (map[any]any, map[any]error) {
	ret := make(map[any]any, 0)
	errs := make(map[any]error, 0)
	ee.expr.Range(func(key, value any) bool {
		stage, err := ee.evaluateStage(value.(*evaluationNode), parameters)
		if err != nil {
			errs[key] = err
		} else {
			ret[key] = stage
		}
		return true
	})
	return ret, errs
}

func (ee *EvaluableExpression) EvaluateAllExpr(parameters map[string]any) (map[any]any, map[any]error) {
	return ee.EvalAllExpr(MapParameters(parameters))
}

func (ee *EvaluableExpression) EvaluateExprByKey(key any, parameters map[string]any) (any, error) {
	return ee.EvalExprByKey(key, MapParameters(parameters))
}

func (ee *EvaluableExpression) EvalExprByKey(key any, parameters goexpression.Parameters) (any, error) {
	if expr, ok := ee.expr.Load(key); ok {
		tempExpr := expr.(*evaluationNode)
		return ee.evaluateStage(tempExpr, parameters)
	} else {
		return nil, errors.New("no found expr")
	}
}

func (ee *EvaluableExpression) AddSingleExpr(expr string) (err error) {
	defer func() {
		if errRec := recover(); errRec != nil {
			err = errors.New(fmt.Sprintf("panic: %+v", errRec))
			return
		}
		err = errors.New("panic: no err info")
		return
	}()
	ee.singleExpr, err = VisitorParserString(expr)
	if err != nil {
		return err
	}
	ee.singleExpr = elideLiterals(ee.singleExpr)
	return nil
}

func (ee *EvaluableExpression) EvalSingleString(parameters goexpression.Parameters) (any, error) {
	return ee.evaluateStage(ee.singleExpr, parameters)
}
func (ee *EvaluableExpression) EvaluateSingleString(parameters map[string]any) (any, error) {
	return ee.evaluateStage(ee.singleExpr, MapParameters(parameters))
}
func (ee *EvaluableExpression) String(key any) (string, error) {
	if key != nil {
		if value, ok := ee.exprString.Load(key); ok {
			return value.(string), nil
		}
		return "", errors.New("no found key")
	}
	if ee.singleString != "" {
		return "", errors.New("no found key")
	}
	return ee.singleString, nil
}

func (ee *EvaluableExpression) evalString(expression string, parameters goexpression.Parameters) (any, error) {
	ee.singleExpr, _ = VisitorParserString(expression)
	ee.singleExpr = elideLiterals(ee.singleExpr)
	return ee.evaluateStage(ee.singleExpr, parameters)
}

func (ee *EvaluableExpression) evaluateStage(stage *evaluationNode, parameters goexpression.Parameters) (ret any, err error) {

	defer func() {
		if errRec := recover(); errRec != nil {
			err = errors.New(fmt.Sprintf("panic: %+v", errRec))
			return
		}
		err = errors.New("panic: no err info")
		return
	}()
	var left, right any

	if stage.LeftOperator != nil {
		left, err = ee.evaluateStage(stage.LeftOperator, parameters)
		if err != nil {
			return nil, err
		}
	}

	if stage.isShortCircuitable() {
		switch stage.Symbol {
		case LOGICAL_AND:
			if left == false {
				return false, nil
			}
		case LOGICAL_OR:
			if left == true {
				return true, nil
			}
			//case COALESCE:
			//	if left != nil {
			//		return left, nil
			//	}

			//case TERNARY_TRUE:
			//	if left == false {
			//		right = shortCircuitHolder
			//	}
			//case TERNARY_FALSE:
			//	if left != nil {
			//		right = shortCircuitHolder
			//	}
			//}
		}
	}

	if right != shortCircuitHolder && stage.RightOperator != nil {
		var rightValue []interface{}
		for _, node := range stage.RightOperator {
			rightTemp, err := ee.evaluateStage(node, parameters)
			if err != nil {
				return nil, err
			}
			rightValue = append(rightValue, rightTemp)
		}
		if len(rightValue) == 1 {
			right = rightValue[0]
		} else {
			right = rightValue
		}
	}

	if ee.ChecksTypes {
		if stage.TypeCheck == nil {

			err = typeCheck(stage.LeftTypeCheck, left, stage.Symbol, stage.TypeErrorFormat)
			if err != nil {
				return nil, err
			}

			err = typeCheck(stage.RightTypeCheck, right, stage.Symbol, stage.TypeErrorFormat)
			if err != nil {
				return nil, err
			}
		} else {
			// special case where the type check needs to know both sides to determine if the Operator can handle it
			if !stage.TypeCheck(left, right) {
				errorMsg := fmt.Sprintf("double check: %+v, value:%v, string:%v", stage.TypeErrorFormat, left, stage.Symbol.String())
				return nil, errors.New(errorMsg)
			}
		}
	}

	if ee.IsDebug {
		ret, err := stage.Operator(left, right, parameters)
		if err != nil {
			ee.recordStep = append(ee.recordStep, stage.RawString)
		} else {
			ee.recordStep = append(ee.recordStep, fmt.Sprintf("%s:%+v", stage.RawString, ret))
		}
		return ret, err
	} else {
		return stage.Operator(left, right, parameters)
	}
}

func typeCheck(check StageTypeCheck, value any, symbol OperatorSymbol, format string) error {

	if check == nil {
		return nil
	}

	if check(value) {
		return nil
	}

	errorMsg := fmt.Sprintf("single check: %+v, value:%v, string:%v", format, value, symbol.String())
	return errors.New(errorMsg)
}
