# goexpression
> 一个`Golang`语言支持动态运行规则表达式的库。

本项目基于 [govaluate](https://github.com/Knetic/govaluate) 和 [antlr4](https://github.com/antlr/antlr4) 
## 特性
- 以`golang`原始语法为主，学习成本低；
- 支持常量预编译计算，提高运行速度；
- 支持三元运算，`in`运算，方便使用；
- 支持多表达式同时运行，贴合业务实际使用；

## 快速开始
```golang
package main

import "fmt"
import "github.com/stoneqi/goexpression"

func main() {
	goExpr := goexpression.NewGoExpression()
	_ = goExpr.AddSingleExpr("1+3*4-num1*num2")
	ret, _ := goExpr.EvalSingleString(map[string]any{
		"num1": 10,
		"num2": 4,
	})
	fmt.Printf("%v\n", ret)
}

```

## 接口定义
```go
type GoExpression interface {
	// AddExpr 添加表达式，key为表达式唯一标识，expr为表达式
	AddExpr(key any, expr string) error
	// AddExprWithParameters 包含预参数，如果表达式中使用的变量在parameters中存在，则会预选读取编译
	AddExprWithParameters(key any, expr string, parameters Parameters) error
	// AddExprWithMap 同AddExprWithParameters，包装map到Parameters
	AddExprWithMap(key any, expr string, parameters map[string]any) error
	// EvalAllExprWithParameters 执行所有表达式，单独返回每个表达式的结果否或错误
	EvalAllExprWithParameters(parameters Parameters) (map[any]any, map[any]error)
	// EvalAllExpr 同EvalAllExprWithParameters，包装map到Parameters
	EvalAllExpr(parameters map[string]any) (map[any]any, map[any]error)
	// EvalExprByKeyWithParameters 执行标识为key的表达式，返回该表达式的执行结果或错误
	EvalExprByKeyWithParameters(key any, parameters Parameters) (any, error)
	// EvalExprByKey 同EvalExprByKeyWithParameters，包装map到Parameters
	EvalExprByKey(key any, parameters map[string]any) (any, error)
	// AddSingleExpr 添加单个表达式
	AddSingleExpr(expr string) error
	// EvalSingleStringWithParameters 执行添加的单个表达式，返回该表达式的执行结果或错误
	EvalSingleStringWithParameters(parameters Parameters) (any, error)
	// EvalSingleString EvalSingleStringWithParameters，包装map到Parameters
	EvalSingleString(parameters map[string]any) (any, error)
	// DeleteExpr 删除包含key值的表达式
	DeleteExpr(key any)
	// DeleteAll 删除所有表达式
	DeleteAll()
	// String 返回对应key值的表达式，如果key为nil，返回AddSingleExpr添加的表达式
	String(key any) (string, error)
	// AllString 返回所有表达式
	AllString() map[any]string
}
```

## 字面量数据类型
- `nil` 空 
- `bool` 布尔，`TRUE`，`True`，`true`，`false`，`False`，`FALSE`
- `int` 整数，内部会把所有整数转为go的int64，注意精度丢失和范围
- `float` 浮点数，内部会把所有浮点数转为go的float64，注意精度丢失和范围
- `string` 字符串，`"string"`双引号内部包含的字符串
- `{}` 数组，`{1,2,a,"string"}`

## 变量
以英文字母开头，类型可以为golang中的任何数据类型，操作时，`int`会统一转为`int64`，`float`会统一转为`float64`
`a+b`，变量通过MapParameter传入
## 操作
### 一元操作
- `+`：正值，值必须为`int`或`float`，`+345`，`+345.0`
- `-`：负值，值必须为`int`或`float`，`-345`，`-345.0`
- `!` 和 `not`：取反，返回`bool`类型，值可为任意值 `!345`, `!true`
- `^`：按位取反，值必须为`int`或`float`，`^123`

### 二元操作
- `*`，`/`，`+`，`-`：值必须为`int`或`float`，两个操作值为`int`则按`int`计算,否则统一转为`float`计算， `+` 支持`string`计算
- `%`
- `<<`
- `>>`
- `&`
- `&^`
- `|`
- `^`
- `==`
- `!=`
- `<`
- `<=`
- `>`
- `>=`
- `and` 和 `&&`
- `or` 和 `||`
- `in` 例如 `1 in {1,2,3}`, 右侧值需要为数组

### 三元操作
- `A?B:C`: 支持三元判断操作，ABC都可为任意表达式， A支持任意类型，为`nil`认为`false`，否则为`true`。例如 `true?1:2`

### 取值操作
- `index`：索引，`{1,3,"AA"}[2]`
- `slice`: 切片， `{1,3,"AA"}[1,2]`
- `.`: 访问， 访问结构体或MAP结构。 `A.B.C`

### 函数调用
- `funName(argum)`。支持函数调用，函数可以在表达式执行时通过接口的参数传入。对函数参数不做限制，要求必须有返回值，返回值为2个是，第二个需要是`error`类型， 执行函数时，返回error不为空，则中断表达式执行，返回错误。
```go
func example(a int, b int ) (int, err) {
	return a + b, nil
}
```