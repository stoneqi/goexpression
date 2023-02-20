package src

import (
	"errors"
	"fmt"
	"sync"
)

const shortCircuitHolder int = -1

type EvaluableExpressionContext struct {
	singleExpr   *evaluationNode
	singleString string
	ChecksTypes  bool
	recordStep   []string
	IsDebug      bool
	expr         sync.Map
	exprString   sync.Map
}

func NewEvaluableExpression() *EvaluableExpressionContext {
	return &EvaluableExpressionContext{}
}

func (ee *EvaluableExpressionContext) AddExpr(key any, expr string) (err error) {
	if key == nil {
		return errors.New("key is nil")
	}
	return ee.AddExprWithParameters(key, expr, MapParameters(map[string]any{}))
}
func (ee *EvaluableExpressionContext) DeleteExpr(key any) {
	if key == nil {
		return
	}
	ee.exprString.Delete(key)
	ee.expr.Delete(key)
	return
}
func (ee *EvaluableExpressionContext) DeleteAll() {
	ee.singleExpr = nil
	ee.singleString = ""
	ee.recordStep = nil
	ee.expr = sync.Map{}
	ee.exprString = sync.Map{}
	return
}

func (ee *EvaluableExpressionContext) AddExprWithMap(key any, expr string, parameters map[string]any) (err error) {
	if key == nil {
		return errors.New("key is nil")
	}
	return ee.AddExprWithParameters(key, expr, MapParameters(parameters))
}

func (ee *EvaluableExpressionContext) AddExprWithParameters(key any, expr string, parameters Parameters) (err error) {
	if key == nil {
		return errors.New("key is nil")
	}

	stageTemp, err := VisitorParserString(&ParserStringContext{
		parameters: parameters,
	}, expr)
	if err != nil {
		return err
	}

	stageTemp = elideLiterals(stageTemp)
	ee.expr.Store(key, stageTemp)
	ee.exprString.Store(key, expr)
	return nil
}

func (ee *EvaluableExpressionContext) EvalAllExprWithParameters(parameters Parameters) (map[any]any, map[any]error) {
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

func (ee *EvaluableExpressionContext) EvalAllExpr(parameters map[string]any) (map[any]any, map[any]error) {
	return ee.EvalAllExprWithParameters(MapParameters(parameters))
}

func (ee *EvaluableExpressionContext) EvalExprByKey(key any, parameters map[string]any) (any, error) {
	return ee.EvalExprByKeyWithParameters(key, MapParameters(parameters))
}

func (ee *EvaluableExpressionContext) EvalExprByKeyWithParameters(key any, parameters Parameters) (any, error) {
	if key == nil {
		return nil, errors.New("key is nil")
	}
	if expr, ok := ee.expr.Load(key); ok {
		tempExpr := expr.(*evaluationNode)
		return ee.evaluateStage(tempExpr, parameters)
	} else {
		return nil, errors.New("no found expr")
	}
}

func (ee *EvaluableExpressionContext) AddSingleExpr(expr string) (err error) {

	ee.singleExpr, err = VisitorParserString(&ParserStringContext{}, expr)
	if err != nil {
		return err
	}
	ee.singleExpr = elideLiterals(ee.singleExpr)
	return nil
}

func (ee *EvaluableExpressionContext) EvalSingleStringWithParameters(parameters Parameters) (any, error) {
	return ee.evaluateStage(ee.singleExpr, parameters)
}
func (ee *EvaluableExpressionContext) EvalSingleString(parameters map[string]any) (any, error) {
	return ee.EvalSingleStringWithParameters(MapParameters(parameters))
}
func (ee *EvaluableExpressionContext) String(key any) (string, error) {
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

func (ee *EvaluableExpressionContext) AllString() map[any]string {
	ret := make(map[any]string, 0)
	ee.exprString.Range(func(key, value any) bool {
		ret[key] = value.(string)
		return true
	})
	return ret
}

func (ee *EvaluableExpressionContext) evalStringTest(expression string, parameters Parameters) (any, error) {
	ee.singleExpr, _ = VisitorParserString(&ParserStringContext{}, expression)
	ee.singleExpr = elideLiterals(ee.singleExpr)
	return ee.evaluateStage(ee.singleExpr, parameters)
}
func (ee *EvaluableExpressionContext) evalStringTest2(expression string, parameters Parameters, parameters2 Parameters) (any, error) {
	ee.singleExpr, _ = VisitorParserString(&ParserStringContext{
		parameters: parameters,
	}, expression)
	ee.singleExpr = elideLiterals(ee.singleExpr)
	return ee.evaluateStage(ee.singleExpr, parameters2)
}

func (ee *EvaluableExpressionContext) evaluateStage(stage *evaluationNode, parameters Parameters) (ret any, err error) {

	defer func() {
		if errRec := recover(); errRec != nil {
			err = errors.New(fmt.Sprintf("panic: %+v", errRec))
			return
		}
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
