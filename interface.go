// Package goexpression provides ...
package goexpression

type goExpression interface {

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
	// DeleteExpr 删除包含key值的表达式
	DeleteExpr(key any)
	// DeleteAll 删除所有表达式
	DeleteAll()
	// String 返回对应key值的表达式，如果key为nil，返回AddSingleExpr添加的表达式
	String(key any) (string, error)
	// AllString 返回所有表达式
	AllString() map[any]string
}

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

// 内部测试使用
func newGoExpression() goExpression {
	return NewEvaluableExpression()
}

func NewGoEvaluator() GoEvaluator {
	return NewEvaluableExpression()
}
