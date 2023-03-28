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
	goExpr := goexpression.NewGoEvaluator()
	// 直接编译表达式并执行
	result, err := goExpr.Eval("1+1", nil)
	if err != nil {
		panic("出错：" + err.Error())
	}
	fmt.Println(result) // 输出：2
}

```

```golang
package main

import "fmt"
import "github.com/stoneqi/goexpression"

func main() {
	goExpr := goexpression.NewGoEvaluator()
	// 预编译表达式，后续可重复执行
	err := goExpr.Parse("1+1", nil)
	if err != nil {
		return
	}
	if err != nil {
		panic("表达式有误：" + err.Error())
	}

	// 计算表达式结果
	result, err := goExpr.Evaluate(nil)

	if err != nil {
		panic("计算结果出错：" + err.Error())
	}

	fmt.Println(result) // 输出：2
}

```

```golang
package main

import "fmt"
import "github.com/stoneqi/goexpression"

func main() {
	var MyFunction = func(a string) (interface{}, error) {
		return a, nil
	}
	goExpr := goexpression.NewGoEvaluator()
	// 编译表达式
	err := goExpr.Parse("MyFunction(hello)", nil)
	if err != nil {
		panic("表达式有误：" + err.Error())
	}
	parameters := make(map[string]any, 0)
	parameters["MyFunction"] = MyFunction
	parameters["hello"] = "Hello World"
	// 计算表达式结果
	result, err := goExpr.Evaluate(goexpression.MapParameters(parameters))

	if err != nil {
		panic("计算结果出错：" + err.Error())
	}

	fmt.Println(result) // 输出：Hello World
}

```

## 接口定义
```go
type GoEvaluator interface {
    // Parse 编译传入的表达式
    Parse(expression string, inputs Parameters) error
    // Evaluate 执行编译好的表达式
    Evaluate(inputs Parameters) (any, error)
    // Eval 添加并执行表达式
    Eval(expression string, data map[string]any) (any, error)
    // EvalWithParameters 添加并执行表达式
    EvalWithParameters(expression string, data Parameters) (any, error)
}

// SyntaxChecker 语法检查器接口
type SyntaxChecker interface {
    SyntaxCheck(expression string) error
}

type Parameters interface {
    Get(name string) (any, error)
}
```
## 接口介绍

### `Parse(expression string, inputs Parameters) error`
解析一个表达式，准备进行执行。

参数：
- expression：待解析的表达式。
- inputs：解析编译时提供的可选参数，可用于提供表达式中包含的参数提前编译。如果参数无效，解析器可能会返回错误。

返回值：
- error：如果解析期间发生错误，则返回错误信息。

### `Evaluate(inputs Parameters) (any, error)`
执行表达式并返回结果。

参数：
- inputs：可选参数，用于在执行表达式时传递使用的变量。如果表达式需要的参数无效，则方法可能会返回错误。

返回值：
- any：计算表达式的结果。
- error：如果表达式的计算过程中发生错误，则返回错误。

### `Eval(expression string, data map[string]any) (any, error)`
解析并执行一个新的表达式。

参数：

- expression：待解析和执行的表达式字符串。
- data：可选参数，用于在执行表达式时传递所需的变量。

返回值：

- any：计算表达式的结果。
- error：如果表达式执行过程中发生错误，则返回错误。

### `EvalWithParameters(expression string, data Parameters) (any, error)`
解析并执行一个新的表达式。

参数：

- expression：待解析和执行的表达式字符串。
- data：可选参数，用于在执行表达式时传递所需的变量。

返回值：

- any：计算表达式的结果。
- error：如果表达式执行过程中发生错误，则返回错误。

### `SyntaxCheck(expression string) error`
检查表达式的语法。如果表达式的语法正确，则方法将返回nil。

参数：

- expression：待检查的表达式字符串。

返回值：

- error：如果表达式语法不正确，则返回错误信息。

### `type Parameters interface`
仅包含一个`Get`方法。实现该接口的对象可以在执行表达式时读取其中的参数。内部实现了一个 MapParameters ，针对map数据结构实现了该接口

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