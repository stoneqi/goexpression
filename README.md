# goexpression

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

## 字面量数据类型
- `nil` 空 
- `bool` 布尔，`TRUE`，`True`，`true`，`false`，`False`，`FALSE`
- `int` 整数，内部会把所有整数转为go的int64，注意精度丢失和范围
- `float` 浮点数，内部会把所有浮点数转为go的float64，注意精度丢失和范围
- `string` 字符串，`"string"`双引号内部包含的字符串
- `{}` 数组，`{1,2,a,"string"}`

## 变量
以英文字母开头，类型可以为golang中的任何数据类型，操作时，int*会统一转为int64，float*会统一转魏float64
`a+b`，变量通过MapParameter传入
## 操作
### 一元操作
- `+`：正值，值必须为`int`或`float`，`+345`，`+345.0`
- `-`：负值，值必须为`int`或`float`，`-345`，`-345.0`
- `!` 和 `not`：取反，返回`bool`类型，值可为任意值 `!345`, `!true`
- `^`：按位取反，值必须为`int`或`float`，`^123`

### 二元操作
- `*`，`/`，`+`，`-`：值必须为`int`或`float`，两个操作值为`int`则按int计算,否则统一转为float计算+` 值可以为`string`
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
- `in'

### 三元操作
- `A?B:C`: `true?1:2`

### 取值操作
- `index`：索引，`{1,3,"AA"}[2]`
- `slice`: 切片， `{1,3,"AA"}[1,2]`

### 函数调用
- `funName(argum)`