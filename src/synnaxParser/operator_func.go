package parserSecond

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
)

var (
	_true  = interface{}(true)
	_false = interface{}(false)
)

// left
func leftOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return left, nil
}

// right
func rightOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return right, nil
}

// 加法： +2=2
func unAryAddOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {

	return right.(float64), nil
}

// 加法： "a"+"b" = "ab"; 1+2=3
func addOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {

	// string concat if either are strings
	if isString(left) || isString(right) {
		return fmt.Sprintf("%v%v", left, right), nil
	}

	return left.(float64) + right.(float64), nil
}

// 减法： 1 - 3 = -2
func subtractOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return left.(float64) - right.(float64), nil
}

// 乘法：
func multiplyOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return left.(float64) * right.(float64), nil
}

// 除法
func divideOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return left.(float64) / right.(float64), nil
}

// 平方
func exponentOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return math.Pow(left.(float64), right.(float64)), nil
}

// 取模
func modulusOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return math.Mod(left.(float64), right.(float64)), nil
}

// 大于等于
func gteOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	if isString(left) && isString(right) {
		return boolIface(left.(string) >= right.(string)), nil
	}
	return boolIface(left.(float64) >= right.(float64)), nil
}

// 大于
func gtOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	if isString(left) && isString(right) {
		return boolIface(left.(string) > right.(string)), nil
	}
	return boolIface(left.(float64) > right.(float64)), nil
}

// 小于等于
func lteOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	if isString(left) && isString(right) {
		return boolIface(left.(string) <= right.(string)), nil
	}
	return boolIface(left.(float64) <= right.(float64)), nil
}

// 小于
func ltOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	if isString(left) && isString(right) {
		return boolIface(left.(string) < right.(string)), nil
	}
	return boolIface(left.(float64) < right.(float64)), nil
}

// 等于
func equalOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return boolIface(reflect.DeepEqual(left, right)), nil
}

// 不等于
func notEqualOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return boolIface(!reflect.DeepEqual(left, right)), nil
}

// and
func andOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return boolIface(left.(bool) && right.(bool)), nil
}

// or
func orOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return boolIface(left.(bool) || right.(bool)), nil
}

// 取负数
func negateOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return -right.(float64), nil
}

// 取反
func invertOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return boolIface(!right.(bool)), nil
}

// 按位取反
func bitwiseNotOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return float64(^int64(right.(float64))), nil
}

//func ternaryIfOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
//	if left.(bool) {
//		return right, nil
//	}
//	return nil, nil
//}
//func ternaryElseOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
//	if left != nil {
//		return left, nil
//	}
//	return right, nil
//}

// 布尔
func regexOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {

	var pattern *regexp.Regexp
	var err error

	switch right.(type) {
	case string:
		pattern, err = regexp.Compile(right.(string))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Unable to compile regexp pattern '%v': %v", right, err))
		}
	case *regexp.Regexp:
		pattern = right.(*regexp.Regexp)
	}

	return pattern.Match([]byte(left.(string))), nil
}

func notRegexOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {

	ret, err := regexOperator(left, right, parameters)
	if err != nil {
		return nil, err
	}

	return !(ret.(bool)), nil
}

// 二元按位或
func bitwiseOrOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return float64(int64(left.(float64)) | int64(right.(float64))), nil
}

// 二元按位且
func bitwiseAndOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return float64(int64(left.(float64)) & int64(right.(float64))), nil
}

// 二元按位异或
func bitwiseXOROperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return float64(int64(left.(float64)) ^ int64(right.(float64))), nil
}

// 位清零
func bitwiseAndNotOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return float64(int64(left.(float64)) &^ int64(right.(float64))), nil
}

// 左移
func leftShiftOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return float64(uint64(left.(float64)) << uint64(right.(float64))), nil
}

// 左移
func rightShiftOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
	return float64(uint64(left.(float64)) >> uint64(right.(float64))), nil
}

// 获取参数值
func makeParameterOperator(parameterName string) EvaluationOperator {

	return func(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
		value, err := parameters.Get(parameterName)
		if err != nil {
			return nil, err
		}

		return value, nil
	}
}

// liter值
func makeLiteralOperator(literal interface{}) EvaluationOperator {
	return func(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
		return literal, nil
	}
}

// 函数执行
func makeFunctionOperator(function ExpressionFunction) EvaluationOperator {

	return func(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {

		if right == nil {
			return function()
		}

		switch right.(type) {
		case []interface{}:
			return function(right.([]interface{})...)
		default:
			return function(right)
		}
	}
}

// 函数执行
func makeExpressionListOperator(expression []*evaluationNode) EvaluationOperator {
	return func(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
		res := make([]interface{}, 0, len(expression))
		for _, node := range expression {
			stage, err := evaluateStage(node, parameters)
			if err != nil {
				return nil, err
			}
			res = append(res, stage)
		}
		return res, nil
	}
}

func evaluateStage(stage *evaluationNode, parameters Parameters) (interface{}, error) {

	var left, right interface{}
	var err error

	if stage.LeftOperator != nil {
		left, err = evaluateStage(stage.LeftOperator, parameters)
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
		}
	}

	if right != shortCircuitHolder && stage.RightOperator != nil {
		right, err = evaluateStage(stage.RightOperator, parameters)
		if err != nil {
			return nil, err
		}
	}

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
			errorMsg := fmt.Sprintf(stage.TypeErrorFormat, left, stage.Symbol.String())
			return nil, errors.New(errorMsg)
		}
	}

	return stage.Operator(left, right, parameters)
}

// 类型转换
func typeConvertParam(p reflect.Value, t reflect.Type) (ret reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			errorMsg := fmt.Sprintf("Argument type conversion failed: failed to convert '%s' to '%s'", p.Kind().String(), t.Kind().String())
			err = errors.New(errorMsg)
			ret = p
		}
	}()

	return p.Convert(t), nil
}

// 类型转换参数
func typeConvertParams(method reflect.Value, params []reflect.Value) ([]reflect.Value, error) {

	methodType := method.Type()
	numIn := methodType.NumIn()
	numParams := len(params)

	if numIn != numParams {
		if numIn > numParams {
			return nil, fmt.Errorf("Too few arguments to parameter call: got %d arguments, expected %d", len(params), numIn)
		}
		return nil, fmt.Errorf("Too many arguments to parameter call: got %d arguments, expected %d", len(params), numIn)
	}

	for i := 0; i < numIn; i++ {
		t := methodType.In(i)
		p := params[i]
		pt := p.Type()

		if t.Kind() != pt.Kind() {
			np, err := typeConvertParam(p, t)
			if err != nil {
				return nil, err
			}
			params[i] = np
		}
	}

	return params, nil
}

// 访问参数 a.b.c
func makeAccessorOperator(pair []string) EvaluationOperator {

	reconstructed := strings.Join(pair, ".")

	return func(left interface{}, right interface{}, parameters Parameters) (ret interface{}, err error) {

		var params []reflect.Value

		value, err := parameters.Get(pair[0])
		if err != nil {
			return nil, err
		}

		// while this library generally tries to handle panic-inducing cases on its own,
		// accessors are a sticky case which have a lot of possible ways to fail.
		// therefore every call to an accessor sets up a defer that tries to recover from panics, converting them to errors.
		defer func() {
			if r := recover(); r != nil {
				errorMsg := fmt.Sprintf("Failed to access '%s': %v", reconstructed, r.(string))
				err = errors.New(errorMsg)
				ret = nil
			}
		}()

		for i := 1; i < len(pair); i++ {

			coreValue := reflect.ValueOf(value)

			var corePtrVal reflect.Value

			// if this is a pointer, resolve it.
			if coreValue.Kind() == reflect.Ptr {
				corePtrVal = coreValue
				coreValue = coreValue.Elem()
			}

			if coreValue.Kind() != reflect.Struct {
				return nil, errors.New("Unable to access '" + pair[i] + "', '" + pair[i-1] + "' is not a struct")
			}

			field := coreValue.FieldByName(pair[i])
			if field != (reflect.Value{}) {
				value = field.Interface()
				continue
			}

			method := coreValue.MethodByName(pair[i])
			if method == (reflect.Value{}) {
				if corePtrVal.IsValid() {
					method = corePtrVal.MethodByName(pair[i])
				}
				if method == (reflect.Value{}) {
					return nil, errors.New("No method or field '" + pair[i] + "' present on parameter '" + pair[i-1] + "'")
				}
			}

			switch right.(type) {
			case []interface{}:

				givenParams := right.([]interface{})
				params = make([]reflect.Value, len(givenParams))
				for idx, _ := range givenParams {
					params[idx] = reflect.ValueOf(givenParams[idx])
				}

			default:

				if right == nil {
					params = []reflect.Value{}
					break
				}

				params = []reflect.Value{reflect.ValueOf(right.(interface{}))}
			}

			params, err = typeConvertParams(method, params)

			if err != nil {
				return nil, errors.New("Method call failed - '" + pair[0] + "." + pair[1] + "': " + err.Error())
			}

			returned := method.Call(params)
			retLength := len(returned)

			if retLength == 0 {
				return nil, errors.New("Method call '" + pair[i-1] + "." + pair[i] + "' did not return any values.")
			}

			if retLength == 1 {

				value = returned[0].Interface()
				continue
			}

			if retLength == 2 {

				errIface := returned[1].Interface()
				err, validType := errIface.(error)

				if validType && errIface != nil {
					return returned[0].Interface(), err
				}

				value = returned[0].Interface()
				continue
			}

			return nil, errors.New("Method call '" + pair[0] + "." + pair[1] + "' did not return either one value, or a value and an error. Cannot interpret meaning.")
		}

		value = castToFloat64(value)
		return value, nil
	}
}

//func separatorOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
//
//	var ret []interface{}
//
//	switch left.(type) {
//	case []interface{}:
//		ret = append(left.([]interface{}), right)
//	default:
//		ret = []interface{}{left, right}
//	}
//
//	return ret, nil
//}

// in 函数
func inOperator(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {

	for _, value := range right.([]interface{}) {
		if left == value {
			return true, nil
		}
	}
	return false, nil
}

/*
Converting a boolean to an interface{} requires an allocation.
We can use interned bools to avoid this cost.
*/
func boolIface(b bool) interface{} {
	if b {
		return _true
	}
	return _false
}
