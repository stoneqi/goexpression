package parserSecond

import (
	"errors"
	"fmt"
)

const shortCircuitHolder int = -1

type EvaluableExpression struct {
	ChecksTypes bool
}

func (this EvaluableExpression) evaluateStage(stage *evaluationNode, parameters Parameters) (interface{}, error) {

	var left, right interface{}
	var err error

	if stage.leftOperator != nil {
		left, err = this.evaluateStage(stage.leftOperator, parameters)
		if err != nil {
			return nil, err
		}
	}

	if stage.isShortCircuitable() {
		switch stage.symbol {
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

	if right != shortCircuitHolder && stage.rightOperator != nil {
		right, err = this.evaluateStage(stage.rightOperator, parameters)
		if err != nil {
			return nil, err
		}
	}

	if this.ChecksTypes {
		if stage.typeCheck == nil {

			err = typeCheck(stage.leftTypeCheck, left, stage.symbol, stage.typeErrorFormat)
			if err != nil {
				return nil, err
			}

			err = typeCheck(stage.rightTypeCheck, right, stage.symbol, stage.typeErrorFormat)
			if err != nil {
				return nil, err
			}
		} else {
			// special case where the type check needs to know both sides to determine if the operator can handle it
			if !stage.typeCheck(left, right) {
				errorMsg := fmt.Sprintf(stage.typeErrorFormat, left, stage.symbol.String())
				return nil, errors.New(errorMsg)
			}
		}
	}

	return stage.operator(left, right, parameters)
}

func typeCheck(check stageTypeCheck, value interface{}, symbol OperatorSymbol, format string) error {

	if check == nil {
		return nil
	}

	if check(value) {
		return nil
	}

	errorMsg := fmt.Sprintf(format, value, symbol.String())
	return errors.New(errorMsg)
}
