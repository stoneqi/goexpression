package src

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// left
func leftOperator(left any, right any, parameters Parameters) (any, error) {
	return left, nil
}

// right
func rightOperator(left any, right any, parameters Parameters) (any, error) {
	return right, nil
}

// 加法： +2=2
func unAryAddOperator(left any, right any, parameters Parameters) (any, error) {
	return right, nil
}

// 加法： "a"+"b" = "ab"; 1+2=3
func addOperator(left any, right any, parameters Parameters) (any, error) {

	if isString(left) || isString(right) {
		return fmt.Sprintf("%v%v", left, right), nil
	}
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) + castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return castToFloat64(left) + castToFloat64(right), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 减法： 1 - 3 = -2
func subtractOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) - castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return castToFloat64(left) - castToFloat64(right), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 乘法：
func multiplyOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) * castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return castToFloat64(left) * castToFloat64(right), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 除法
func divideOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) / castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return castToFloat64(left) / castToFloat64(right), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 平方
//func exponentOperator(left any, right any, parameters Parameters) (any, error) {
//	return math.Pow(left.(float64), right.(float64)), nil
//}

// 取模
func modulusOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) % castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return math.Mod(castToFloat64(left), castToFloat64(right)), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 大于等于
func gteOperator(left any, right any, parameters Parameters) (any, error) {
	if isString(left) && isString(right) {
		return left.(string) >= right.(string), nil
	}
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) >= castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return castToFloat64(left) >= castToFloat64(right), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 大于
func gtOperator(left any, right any, parameters Parameters) (any, error) {
	if isString(left) && isString(right) {
		return left.(string) > right.(string), nil
	}
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) > castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return castToFloat64(left) > castToFloat64(right), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 小于等于
func lteOperator(left any, right any, parameters Parameters) (any, error) {
	if isString(left) && isString(right) {
		return left.(string) <= right.(string), nil
	}
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) <= castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return castToFloat64(left) <= castToFloat64(right), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 小于
func ltOperator(left any, right any, parameters Parameters) (any, error) {
	if isString(left) && isString(right) {
		return left.(string) < right.(string), nil
	}
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) < castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return castToFloat64(left) < castToFloat64(right), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 等于
func equalOperator(left any, right any, parameters Parameters) (any, error) {
	return reflect.DeepEqual(left, right), nil
}

// 不等于
func notEqualOperator(left any, right any, parameters Parameters) (any, error) {
	return !reflect.DeepEqual(left, right), nil
}

// and
func andOperator(left any, right any, parameters Parameters) (any, error) {
	return boolJudge(left) && boolJudge(right), nil
}

// or
func orOperator(left any, right any, parameters Parameters) (any, error) {
	return boolJudge(left) || boolJudge(right), nil
}

// 取负数
func negateOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(right) {
		return -castToInt64(right), nil
	}
	if isFloat64(right) {
		return -castToFloat64(right), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 取反
func invertOperator(left any, right any, parameters Parameters) (any, error) {
	return !boolJudge(right), nil
}

// 按位取反
func bitwiseNotOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(right) {
		return ^castToInt64(right), nil
	}
	if isFloat64(right) {
		return ^int64(castToFloat64(right)), nil
	}
	return 0, errors.New("parameters type is nil")
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
func regexOperator(left any, right any, parameters Parameters) (any, error) {

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

func notRegexOperator(left any, right any, parameters Parameters) (any, error) {

	ret, err := regexOperator(left, right, parameters)
	if err != nil {
		return nil, err
	}

	return !(ret.(bool)), nil
}

// 二元按位或
func bitwiseOrOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) | castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return float64(int64(castToFloat64(left)) | int64(castToFloat64(right))), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 二元按位且
func bitwiseAndOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) & castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return float64(int64(castToFloat64(left)) & int64(castToFloat64(right))), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 二元按位异或
func bitwiseXOROperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) ^ castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return float64(int64(castToFloat64(left)) ^ int64(castToFloat64(right))), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 位清零
func bitwiseAndNotOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) &^ castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return float64(int64(castToFloat64(left)) &^ int64(castToFloat64(right))), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 左移
func leftShiftOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) << castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return float64(int64(castToFloat64(left)) << int64(castToFloat64(right))), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 左移
func rightShiftOperator(left any, right any, parameters Parameters) (any, error) {
	if isInt64(left) && isInt64(right) {
		return castToInt64(left) >> castToInt64(right), nil
	}
	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return float64(int64(castToFloat64(left)) >> int64(castToFloat64(right))), nil
	}
	return 0, errors.New("parameters type is nil")
}

// 获取参数值
func makeParameterOperator(parameterName string) EvaluationOperator {

	return func(left any, right any, parameters Parameters) (any, error) {
		value, err := parameters.Get(parameterName)
		if err != nil {
			return nil, err
		}
		if isInt64(value) {
			return castToInt64(value), nil
		}
		if isFloat64(value) {
			return castToFloat64(value), nil
		}
		return value, nil
	}
}

// index
func indexOperator(left any, right any, parameters Parameters) (any, error) {

	leftValue := reflect.ValueOf(left)

	if leftValue.Kind() == reflect.Map || leftValue.Kind() == reflect.Array || leftValue.Kind() == reflect.Slice || leftValue.Kind() == reflect.String {
		var indexFloat int64 = -1
		if isInt64(right) {
			indexFloat = castToInt64(right)
		}
		if isFloat64(right) {
			indexFloat = int64(castToFloat64(right))
		}
		if indexFloat < 0 || indexFloat >= int64(leftValue.Len()) {
			return nil, errors.New("map index out of range")
		}
		var value reflect.Value
		if leftValue.Kind() == reflect.Map {
			value = leftValue.MapIndex(reflect.ValueOf(right))
		} else {
			value = leftValue.Index(int(indexFloat))
		}
		if value.IsValid() {
			valueRet := value.Interface()
			if isInt64(valueRet) {
				return castToInt64(valueRet), nil
			}
			if isFloat64(valueRet) {
				return castToFloat64(valueRet), nil
			}
			return valueRet, nil
		} else {
			return nil, errors.New(reflect.ValueOf(right).String() + "index out of range")
		}
	}
	return nil, errors.New("lefRight is no array or slice or map or string")

}

// liter值
func makeLiteralOperator(literal any) EvaluationOperator {
	return func(left any, right any, parameters Parameters) (any, error) {
		if isInt64(literal) {
			return castToInt64(literal), nil
		}
		if isFloat64(literal) {
			return castToFloat64(literal), nil
		}
		return literal, nil
	}
}

// 三元运算符
func conditionalOperator(left any, right any, parameters Parameters) (any, error) {
	boolValue := boolJudge(left)
	conditionalValue, ok := right.([]any)
	if !ok || len(conditionalValue) != 2 {
		return nil, errors.New("conditional size is not 2")
	}
	if boolValue {
		return conditionalValue[0], nil
	} else {
		return conditionalValue[1], nil
	}
}

// 函数执行
func makeFunctionOperator(left any, right any, parameters Parameters) (any, error) {
	if left == nil {
		return nil, errors.New("no Function")
	}
	if right == nil {
		return nil, errors.New("no Parameters")
	}

	fun, ok := left.(func(arguments ...any) (any, error))
	if !ok {
		return nil, errors.New("left is not func type")
	}
	switch right.(type) {
	case []any:
		return fun(right.([]any)...)
	default:
		return fun(right)
	}
}

// 函数执行
func makeFunction2Operator(left any, right any, parameters Parameters) (any, error) {
	if left == nil {
		return nil, errors.New("no Function")
	}
	if right == nil {
		return nil, errors.New("no Parameters")
	}

	funValue := reflect.ValueOf(left)
	funType := reflect.TypeOf(left)
	if funType.Kind() != reflect.Func {
		return nil, errors.New("no func type")
	}
	if funType.NumOut() == 0 {
		return nil, errors.New("return param gt 2")
	}
	if funType.NumOut() > 2 {
		return nil, errors.New("return param gt 2")
	}
	parametersValue, ok := right.([]any)
	if !ok {
		parametersValue = append(parametersValue, right)
	}
	if funType.NumIn() != len(parametersValue) {
		return nil, errors.New("parameters len no eq")
	}

	var params []reflect.Value
	for i := 0; i < funType.NumIn(); i++ {
		paramsType := reflect.TypeOf(parametersValue[i])
		if !paramsType.AssignableTo(funType.In(i)) {
			return nil, errors.New("parameters type no eq")
		} else {
			params = append(params, reflect.ValueOf(parametersValue[i]))
		}
	}

	returned := funValue.Call(params)
	retLength := len(returned)

	if retLength == 0 {
		return nil, errors.New("Method call '" + funValue.String() + "' did not return any values.")
	}

	valueRet := returned[0].Interface()
	if isInt64(valueRet) {
		valueRet = castToInt64(valueRet)
	}
	if isFloat64(valueRet) {
		valueRet = castToFloat64(valueRet)
	}

	if retLength == 1 {
		return valueRet, nil
	}

	if retLength == 2 {
		errIface := returned[1].Interface()
		err, validType := errIface.(error)
		if validType && errIface != nil {
			return valueRet, err
		}
		return valueRet, nil
	}

	return nil, errors.New("method call fail")

	//if !ok {
	//	return nil, errors.New("left is not func type")
	//}
	//switch right.(type) {
	//case []any:
	//	return fun(right.([]any)...)
	//default:
	//	return fun(right)
	//}
}

//// 函数执行
//func makeFuncOperator(function ExpressionFunction) EvaluationOperator {
//	return func(left interface{}, right interface{}, parameters Parameters) (interface{}, error) {
//		res := make([]interface{}, 0, len(expression))
//		for _, node := range expression {
//			singleExpr, err := evaluateStage(node, parameters)
//			if err != nil {
//				return nil, err
//			}
//			res = append(res, singleExpr)
//		}
//		return res, nil
//	}
//}

// 函数参数
func makeSliceOperator(left any, right any, parameters Parameters) (any, error) {
	rightValue, ok := right.([]any)
	if !ok {
		return nil, errors.New("slice index is not array")
	}
	var leftIndex int64 = -1
	if isInt64(rightValue[0]) {
		leftIndex = castToInt64(rightValue[0])
	}
	if isFloat64(rightValue[0]) {
		leftIndex = int64(castToFloat64(rightValue[0]))
	}
	var rightIndex int64 = 0
	if isInt64(rightValue[1]) {
		rightIndex = castToInt64(rightValue[1])
	}
	if isFloat64(rightValue[1]) {
		rightIndex = int64(castToFloat64(rightValue[1]))
	}

	if len(rightValue) == 2 {
		leftValue := reflect.ValueOf(left)
		if leftValue.Kind() == reflect.Array || leftValue.Kind() == reflect.Slice {
			if int(leftIndex) < 0 || rightIndex == 0 || int(rightIndex) > leftValue.Len() || leftIndex >= rightIndex {
				return nil, errors.New("slice bounds out of range")
			}
			value := leftValue.Slice(int(leftIndex), int(rightIndex))
			if value.IsValid() {
				return value.Interface(), nil
			} else {
				return nil, errors.New("no found in array" + reflect.ValueOf(right).String())
			}
		}
	}

	return nil, errors.New("slice params length must 2, actual:" + strconv.Itoa(len(rightValue)))
}

// 函数参数
//func makeMultiExpressionOperator(expression []*evaluationNode) EvaluationOperator {
//	return func(left any, right any, parameters Parameters) (any, error) {
//		res := make([]any, 0, len(expression))
//		for _, node := range expression {
//			singleExpr, err := evaluateStage(node, parameters)
//			if err != nil {
//				return nil, err
//			}
//			res = append(res, singleExpr)
//		}
//		return res, nil
//	}
//}

//func evaluateStage(singleExpr *evaluationNode, parameters Parameters) (any, error) {
//
//	var left, right any
//	var err error
//
//	if singleExpr.LeftOperator != nil {
//		left, err = evaluateStage(singleExpr.LeftOperator, parameters)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	if singleExpr.isShortCircuitable() {
//		switch singleExpr.Symbol {
//		case LOGICAL_AND:
//			if left == false {
//				return false, nil
//			}
//		case LOGICAL_OR:
//			if left == true {
//				return true, nil
//			}
//		}
//	}
//
//	if right != shortCircuitHolder && singleExpr.RightOperator != nil {
//		right, err = evaluateStage(singleExpr.RightOperator, parameters)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	if singleExpr.TypeCheck == nil {
//
//		err = typeCheck(singleExpr.LeftTypeCheck, left, singleExpr.Symbol, singleExpr.TypeErrorFormat)
//		if err != nil {
//			return nil, err
//		}
//
//		err = typeCheck(singleExpr.RightTypeCheck, right, singleExpr.Symbol, singleExpr.TypeErrorFormat)
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		// special case where the type check needs to know both sides to determine if the Operator can handle it
//		if !singleExpr.TypeCheck(left, right) {
//			errorMsg := fmt.Sprintf(singleExpr.TypeErrorFormat, left, singleExpr.Symbol.String())
//			return nil, errors.New(errorMsg)
//		}
//	}
//
//	return singleExpr.Operator(left, right, parameters)
//}

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

func makeAccessOperator(pair string) EvaluationOperator {

	return func(left any, right any, parameters Parameters) (ret any, err error) {

		value := left
		// while this library generally tries to handle panic-inducing cases on its own,
		// accessors are a sticky case which have a lot of possible ways to fail.
		// therefore every call to an accessor sets up a defer that tries to recover from panics, converting them to errors.
		defer func() {
			if r := recover(); r != nil {
				errorMsg := fmt.Sprintf("Failed to access '%s': %v", pair, r.(string))
				err = errors.New(errorMsg)
				ret = nil
			}
		}()

		//for i := 1; i < len(pair); i++ {

		coreValue := reflect.ValueOf(value)

		var corePtrVal reflect.Value

		// if this is a pointer, resolve it.
		if coreValue.Kind() == reflect.Ptr {
			coreValue = coreValue.Elem()
		}

		if coreValue.Kind() != reflect.Struct {
			return nil, errors.New("Unable to access '" + pair + "', '" + coreValue.String() + "' is not a struct or map")
		}

		field := coreValue.FieldByName(pair)
		if field.IsValid() {
			return field.Interface(), nil
		}

		method := coreValue.MethodByName(pair)
		if method == (reflect.Value{}) {
			if corePtrVal.IsValid() {
				method = corePtrVal.MethodByName(pair)
			}
			if method.IsValid() {
				return method.Interface(), nil
			}
		}
		return nil, errors.New("")
	}
}

// 访问参数 a.b.c
func makeAccessorOperator(pair []string) EvaluationOperator {

	reconstructed := strings.Join(pair, ".")

	return func(left any, right any, parameters Parameters) (ret any, err error) {

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
			case []any:

				givenParams := right.([]any)
				params = make([]reflect.Value, len(givenParams))
				for idx, _ := range givenParams {
					params[idx] = reflect.ValueOf(givenParams[idx])
				}

			default:

				if right == nil {
					params = []reflect.Value{}
					break
				}

				params = []reflect.Value{reflect.ValueOf(right.(any))}
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

		//value = castToFloat64(value)
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
func inOperator(left any, right any, parameters Parameters) (any, error) {
	//leftValue := reflect.ValueOf(left)
	rightValue := reflect.ValueOf(right)
	if rightValue.Kind() == reflect.Array || rightValue.Kind() == reflect.Slice {
		for i := 0; i < rightValue.Len(); i++ {
			itemIndex := rightValue.Index(i).Interface()
			if reflect.DeepEqual(left, itemIndex) {
				return true, nil
			}
		}
		return false, nil
	}
	return nil, errors.New("right value type is err: " + rightValue.Kind().String())

}
