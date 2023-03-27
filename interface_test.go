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
			got := NewGoExpression()
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

func ExampleEvalSingleString() {
	goExpr := NewGoExpression()
	_ = goExpr.AddSingleExpr("1+3*4-num1*num2")
	ret, _ := goExpr.EvalSingleString(map[string]any{
		"num1": 10,
		"num2": 4,
	})
	fmt.Printf("%v\n", ret)
}
