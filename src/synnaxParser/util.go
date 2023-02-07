package parserSecond

import (
	"reflect"
)

func castToFloat64(value any) float64 {
	switch value.(type) {
	case uint:
		return float64(value.(uint))
	case uint8:
		return float64(value.(uint8))
	case uint16:
		return float64(value.(uint16))
	case uint32:
		return float64(value.(uint32))
	case uint64:
		return float64(value.(uint64))
	case int:
		return float64(value.(int))
	case int8:
		return float64(value.(int8))
	case int16:
		return float64(value.(int16))
	case int32:
		return float64(value.(int32))
	case int64:
		return float64(value.(int64))
	case float32:
		return float64(value.(float32))
	case float64:
		return value.(float64)
	}
	return 0.0
}

func castToInt64(value any) int64 {
	switch value.(type) {
	case uint:
		return int64(value.(uint))
	case uint8:
		return int64(value.(uint8))
	case uint16:
		return int64(value.(uint16))
	case uint32:
		return int64(value.(uint32))
	case uint64:
		return int64(value.(uint64))
	case int:
		return int64(value.(int))
	case int8:
		return int64(value.(int8))
	case int16:
		return int64(value.(int16))
	case int32:
		return int64(value.(int32))
	case int64:
		return value.(int64)
	}
	return 0
}

//func CheckTypeByReflectNil(arg interface{}) bool {
//	if reflect.ValueOf(arg).IsNil() { //利用反射直接判空，指针用isNil
//		// 函数解释：isNil() bool	判断值是否为 nil
//		// 如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的v== nil操作
//		return reflect.ValueOf(arg).IsValid()
//	}
//}
//
//func CheckTypeByReflectZero(arg interface{}) {
//	if reflect.ValueOf(arg).IsZero() { //利用反射直接判空，基础数据类型用isZero
//		fmt.Printf("反射判断：数据类型为%s,数据值为：%v,nil：%v \n",
//			reflect.TypeOf(arg).Kind(), reflect.ValueOf(arg), reflect.ValueOf(arg).IsValid())
//	}
//}

func boolJudge(value any) bool {
	return !reflect.ValueOf(value).IsZero()
}
