package parserSecond

import "regexp"

// 是字符串
func isString(value any) bool {

	switch value.(type) {
	case string:
		return true
	}
	return false
}

// regex or string
func isRegexOrString(value any) bool {

	switch value.(type) {
	case string:
		return true
	case *regexp.Regexp:
		return true
	}
	return false
}

// 布尔
func isBool(value any) bool {
	switch value.(type) {
	case bool:
		return true
	}
	return false
}

func isNumber(value any) bool {
	if isFloat64(value) || isInt64(value) {
		return true
	}
	return false
}

// 字符
func isFloat64(value any) bool {
	switch value.(type) {
	case float32, float64:
		return true
	}
	return false
}

// 字符
func isInt64(value any) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

// 加法，支持字符串和数字
func additionTypeCheck(left any, right any) bool {

	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return true
	}
	if isString(left) && isString(right) {
		return true
	}
	return false
}

// 比较操作类型检测
func comparatorTypeCheck(left any, right any) bool {

	if (isFloat64(left) || isInt64(left)) && (isFloat64(right) || isInt64(right)) {
		return true
	}
	if isString(left) && isString(right) {
		return true
	}
	return false
}

// 数组
func isArray(value any) bool {
	switch value.(type) {
	case []any:
		return true
	}
	return false
}
