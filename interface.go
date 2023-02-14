// Package goexpression provides ...
package goexpression

type GoExpression interface {
	AddExpr(key any, expr string) error
	EvalAllExpr(parameters Parameters) (map[any]any, map[any]error)
	EvalExprByKey(key any, parameters Parameters) (any, error)
	EvaluateAllExpr(parameters map[string]any) (map[any]any, map[any]error)
	EvaluateExprByKey(key any, parameters map[string]any) (any, error)
	AddSingleExpr(expr string) error
	EvalSingleString(parameters Parameters) (any, error)
	EvaluateSingleString(parameters map[string]any) (any, error)
	String(key any) (string, error)
}

type Parameters interface {
	Get(name string) (any, error)
}
