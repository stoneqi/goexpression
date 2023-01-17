package parserSecond

import (
	"errors"
	"fmt"
)

const shortCircuitHolder int = -1

type EvaluableExpression struct {
	stage       *evaluationNode
	ChecksTypes bool
	recordStep  []string
	IsDebug     bool
}

func (ee *EvaluableExpression) EvalString(expression string, parameters Parameters) (interface{}, error) {
	ee.stage, _ = VisitorParserString(expression)
	return ee.evaluateStage(ee.stage, parameters)
}

func (ee *EvaluableExpression) evaluateStage(stage *evaluationNode, parameters Parameters) (interface{}, error) {

	var left, right interface{}
	var err error

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
		right, err = ee.evaluateStage(stage.RightOperator, parameters)
		if err != nil {
			return nil, err
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
		value, err := stage.Operator(left, right, parameters)
		if err != nil {
			ee.recordStep = append(ee.recordStep, stage.RawString)
		} else {
			ee.recordStep = append(ee.recordStep, fmt.Sprintf("%s:%+v", stage.RawString, value))
		}
		return value, err
	} else {
		return stage.Operator(left, right, parameters)
	}
}

func typeCheck(check StageTypeCheck, value interface{}, symbol OperatorSymbol, format string) error {

	if check == nil {
		return nil
	}

	if check(value) {
		return nil
	}

	errorMsg := fmt.Sprintf("single check: %+v, value:%v, string:%v", format, value, symbol.String())
	return errors.New(errorMsg)
}
