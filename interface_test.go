package goexpression

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoExpression(t *testing.T) {
	tests := []struct {
		name    string
		expr    map[string]string
		param   map[string]any
		want    map[string]any
		wantErr map[string]bool
	}{
		{
			name: "multi_expr",
			expr: map[string]string{
				"test1": "1+2",
			},
			param: nil,
			want: map[string]any{
				"test1": int64(1 + 2),
			},
			wantErr: map[string]bool{
				"test1": false,
			},
		},
		{
			name: "multi_expr",
			expr: map[string]string{
				"test1": "-1+2",
			},
			param: nil,
			want: map[string]any{
				"test1": int64(-1 + 2),
			},
			wantErr: map[string]bool{
				"test1": false,
			},
		},
		{
			name: "multi_expr",
			expr: map[string]string{
				"test1": "-0+2",
			},
			param: nil,
			want: map[string]any{
				"test1": int64(-0 + 2),
			},
			wantErr: map[string]bool{
				"test1": false,
			},
		},
		{
			name: "multi_expr",
			expr: map[string]string{
				"test1": "1+2",
				"test2": "3+4",
			},
			param: nil,
			want: map[string]any{
				"test1": int64(1 + 2),
				"test2": int64(3 + 4),
			},
			wantErr: map[string]bool{
				"test1": false,
				"test2": false,
			},
		},
		{
			name: "multi_expr",
			expr: map[string]string{
				"test1": "a+b",
				"test2": "b+c",
			},
			param: map[string]any{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			want: map[string]any{
				"test1": int64(1 + 2),
				"test2": int64(2 + 3),
			},
			wantErr: map[string]bool{
				"test1": false,
				"test2": false,
			},
		},
		{
			name: "multi_expr",
			expr: map[string]string{
				"test1": "a+b",
				"test2": "b > c",
			},
			param: map[string]any{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			want: map[string]any{
				"test1": int64(1 + 2),
				"test2": bool(2 > 3),
			},
			wantErr: map[string]bool{
				"test1": false,
				"test2": false,
			},
		},
		{
			name: "multi_expr",
			expr: map[string]string{
				"test1": "a+b",
				"test2": "b > c",
				"test3": "a > c",
			},
			param: map[string]any{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			want: map[string]any{
				"test1": int64(1 + 2),
				"test2": bool(2 > 3),
				"test3": bool(1 > 3),
			},
			wantErr: map[string]bool{
				"test1": false,
				"test2": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newGoExpression()
			for key, value := range tt.expr {
				err := got.AddExpr(key, value)
				if err != nil {
					t.Errorf("addExpr error:%s", err)
				}
			}
			rets, errs := got.EvalAllExpr(tt.param)
			for key, value := range tt.wantErr {
				if value && errs[key] != nil {
					continue
				}
				if rets[key] != nil && reflect.DeepEqual(rets[key], tt.want[key]) {
					continue
				}
				t.Errorf("errs= %v, rets= %v", errs[key], rets[key])
			}
		})
	}
}

func TestExampleEvalSingleString(t *testing.T) {
	goExpr := NewGoEvaluator()
	err := goExpr.Parse("1+1", nil)
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

func TestExampleEvalSingleString2(t *testing.T) {

	var MyFunction = func() (interface{}, error) {
		return "Hello World", nil
	}
	goExpr := NewGoEvaluator()
	// 编译表达式
	err := goExpr.Parse("MyFunction()", nil)
	if err != nil {
		panic("表达式有误：" + err.Error())
	}
	parameters := make(map[string]any, 0)
	parameters["MyFunction"] = MyFunction
	// 计算表达式结果
	result, err := goExpr.Evaluate(MapParameters(parameters))

	if err != nil {
		panic("计算结果出错：" + err.Error())
	}

	fmt.Println(result) // 输出：Hello World
}

func TestExampleEvalSingleString3(t *testing.T) {
	goExpr := NewGoEvaluator()
	// 编译表达式并执行
	result, err := goExpr.Eval("1+1", nil)
	if err != nil {
		panic("出错：" + err.Error())
	}
	fmt.Println(result) // 输出：2
}
