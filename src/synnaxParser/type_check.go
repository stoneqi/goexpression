package parserSecond

import "regexp"

// 是字符串
func isString(value interface{}) bool {

	switch value.(type) {
	case string:
		return true
	}
	return false
}

// regex or string
func isRegexOrString(value interface{}) bool {

	switch value.(type) {
	case string:
		return true
	case *regexp.Regexp:
		return true
	}
	return false
}

// 布尔
func isBool(value interface{}) bool {
	switch value.(type) {
	case bool:
		return true
	}
	return false
}

// 字符
func isFloat64(value interface{}) bool {
	switch value.(type) {
	case float64:
		return true
	}
	return false
}

// 加法，支持字符串和数字
func additionTypeCheck(left interface{}, right interface{}) bool {

	if isFloat64(left) && isFloat64(right) {
		return true
	}
	if !isString(left) && !isString(right) {
		return false
	}
	return true
}

// 比较操作类型检测
func comparatorTypeCheck(left interface{}, right interface{}) bool {

	if isFloat64(left) && isFloat64(right) {
		return true
	}
	if isString(left) && isString(right) {
		return true
	}
	return false
}

// 数组
func isArray(value interface{}) bool {
	switch value.(type) {
	case []interface{}:
		return true
	}
	return false
}
