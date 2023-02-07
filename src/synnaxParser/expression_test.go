package parserSecond

import (
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestEvaluableExpressionBasicLit_EvalString(t *testing.T) {
	type fields struct {
		stage       *evaluationNode
		ChecksTypes bool
		IsDebug     bool
	}
	type args struct {
		expression string
		parameters Parameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "nil",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "nil",
				parameters: MapParameters{},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "bool",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "bool",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "false",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
		{
			name: "integer",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "0",
				parameters: MapParameters{},
			},
			want:    int64(0),
			wantErr: false,
		},
		{
			name: "integer_max",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "9223372036854775807",
				parameters: MapParameters{},
			},
			want:    int64(9223372036854775807),
			wantErr: false,
		},
		{
			name: "integer_min",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "-9223372036854775808",
				parameters: MapParameters{},
			},
			want:    int64(-9223372036854775807),
			wantErr: false,
		},
		{
			name: "float",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "-0.0",
				parameters: MapParameters{},
			},
			want:    -0.0,
			wantErr: false,
		},
		{
			name: "float",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "-3.2678668756546343435563242",
				parameters: MapParameters{},
			},
			want:    -3.2678668756546343435563242,
			wantErr: false,
		},
		{
			name: "float_max",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "1.79769313486231570814527423731704356798070e+308",
				parameters: MapParameters{},
			},
			want:    math.MaxFloat64,
			wantErr: false,
		},
		{
			name: "float_min",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "4.9406564584124654417656879286822137236505980e-324",
				parameters: MapParameters{},
			},
			want:    math.SmallestNonzeroFloat64,
			wantErr: false,
		},
		{
			name: "string",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "\"3423asvas%6712你好\"",
				parameters: MapParameters{},
			},
			want:    "3423asvas%6712你好",
			wantErr: false,
		},
		{
			name: "string",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "\"3423asvas%	6712你好\"",
				parameters: MapParameters{},
			},
			want:    "3423asvas%	6712你好",
			wantErr: false,
		},
		{
			name: "string",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "\"3423QWERTYUIOPASDDFGHJKLZXCVBNM;[]\\\"/.,%6712你好\"",
				parameters: MapParameters{},
			},
			want:    "3423QWERTYUIOPASDDFGHJKLZXCVBNM;[]\"/.,%6712你好",
			wantErr: false,
		},
		{
			name: "string",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "\"3423QWERTYUIOPASDDFGHJKLZXCVBNM;[]'/.,%6712你好\"",
				parameters: MapParameters{},
			},
			want:    "3423QWERTYUIOPASDDFGHJKLZXCVBNM;[]'/.,%6712你好",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ee := &EvaluableExpression{
				stage:       tt.fields.stage,
				ChecksTypes: tt.fields.ChecksTypes,
				IsDebug:     true,
			}
			got, err := ee.EvalString(tt.args.expression, tt.args.parameters)
			t.Logf("%s; got:%+v; type:%+v", strings.Join(ee.recordStep, "; "), got, reflect.TypeOf(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalString() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestEvaluableExpressionOperandName_EvalString(t *testing.T) {
	type fields struct {
		stage       *evaluationNode
		ChecksTypes bool
		IsDebug     bool
	}
	type args struct {
		expression string
		parameters Parameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a",
				parameters: MapParameters{
					"a": 1,
				},
			},
			want:    int64(1),
			wantErr: false,
		},
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a",
				parameters: MapParameters{
					"a": true,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a",
				parameters: MapParameters{
					"a": struct {
						A int
					}{A: 234},
				},
			},
			want: struct {
				A int
			}{A: 234},
			wantErr: false,
		},
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a",
				parameters: MapParameters{
					"a": nil,
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a",
				parameters: MapParameters{
					"a": 1.0,
				},
			},
			want:    float64(1.0),
			wantErr: false,
		},
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a",
				parameters: MapParameters{
					"a": "aaaa",
				},
			},
			want:    "aaaa",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ee := &EvaluableExpression{
				stage:       tt.fields.stage,
				ChecksTypes: tt.fields.ChecksTypes,
				IsDebug:     true,
			}
			got, err := ee.EvalString(tt.args.expression, tt.args.parameters)
			t.Logf("%s; got:%+v; type:%+v", strings.Join(ee.recordStep, "; "), got, reflect.TypeOf(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalString() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestEvaluableExpressionOperand_EvalString(t *testing.T) {
	type fields struct {
		stage       *evaluationNode
		ChecksTypes bool
		IsDebug     bool
	}
	type args struct {
		expression string
		parameters Parameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a or b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "false or false",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ee := &EvaluableExpression{
				stage:       tt.fields.stage,
				ChecksTypes: tt.fields.ChecksTypes,
				IsDebug:     true,
			}
			got, err := ee.EvalString(tt.args.expression, tt.args.parameters)
			t.Logf("%s; got:%+v; type:%+v", strings.Join(ee.recordStep, "; "), got, reflect.TypeOf(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalString() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestEvaluableExpressionPrimaryExpr_EvalString(t *testing.T) {
	type fields struct {
		stage       *evaluationNode
		ChecksTypes bool
		IsDebug     bool
	}
	type args struct {
		expression string
		parameters Parameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a or b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "false or false",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ee := &EvaluableExpression{
				stage:       tt.fields.stage,
				ChecksTypes: tt.fields.ChecksTypes,
				IsDebug:     true,
			}
			got, err := ee.EvalString(tt.args.expression, tt.args.parameters)
			t.Logf("%s; got:%+v; type:%+v", strings.Join(ee.recordStep, "; "), got, reflect.TypeOf(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalString() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestEvaluableExpressionStmt_EvalString(t *testing.T) {
	type fields struct {
		stage       *evaluationNode
		ChecksTypes bool
		IsDebug     bool
	}
	type args struct {
		expression string
		parameters Parameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "+b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "+345",
				parameters: MapParameters{},
			},
			want:    int64(+345),
			wantErr: false,
		},
		{
			name: "-b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "-345",
				parameters: MapParameters{},
			},
			want:    int64(-345),
			wantErr: false,
		},
		{
			name: "!true",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "!true",
				parameters: MapParameters{},
			},
			want:    !true,
			wantErr: false,
		},
		{
			name: "!false",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "!false",
				parameters: MapParameters{},
			},
			want:    !false,
			wantErr: false,
		},
		{
			name: "not true",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "not true",
				parameters: MapParameters{},
			},
			want:    !true,
			wantErr: false,
		},
		{
			name: "not false",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "not false",
				parameters: MapParameters{},
			},
			want:    !false,
			wantErr: false,
		},
		{
			name: "^b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "^1",
				parameters: MapParameters{},
			},
			want:    int64(^1),
			wantErr: false,
		},
		{
			name: "a*b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123*345",
				parameters: MapParameters{},
			},
			want:    int64(123 * 345),
			wantErr: false,
		},
		{
			name: "a/b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123/345",
				parameters: MapParameters{},
			},
			want:    int64(123 / 345),
			wantErr: false,
		},
		{
			name: "a%b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "100%10",
				parameters: MapParameters{},
			},
			want:    int64(100 % 10),
			wantErr: false,
		},
		{
			name: "a%b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "101%10",
				parameters: MapParameters{},
			},
			want:    int64(101 % 10),
			wantErr: false,
		},
		{
			name: "a << b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "2 << 1",
				parameters: MapParameters{},
			},
			want:    int64(2 << 1),
			wantErr: false,
		},
		{
			name: "a >> b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "2 >> 1",
				parameters: MapParameters{},
			},
			want:    int64(2 >> 1),
			wantErr: false,
		},
		{
			name: "a & b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "100 & 50",
				parameters: MapParameters{},
			},
			want:    int64(100 & 50),
			wantErr: false,
		},
		{
			name: "a &^ b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "100 &^ 50",
				parameters: MapParameters{},
			},
			want:    int64(100 &^ 50),
			wantErr: false,
		},
		{
			name: "a+b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "0x123+345",
				parameters: MapParameters{},
			},
			want:    int64(0x123 + 345),
			wantErr: false,
		},
		{
			name: "a-b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123-345",
				parameters: MapParameters{},
			},
			want:    int64(123 - 345),
			wantErr: false,
		},
		{
			name: "a | b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "100 | 50",
				parameters: MapParameters{},
			},
			want:    int64(100 | 50),
			wantErr: false,
		},
		{
			name: "a ^ b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "100 ^ 50",
				parameters: MapParameters{},
			},
			want:    int64(100 ^ 50),
			wantErr: false,
		},
		{
			name: "a == b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123 == 345",
				parameters: MapParameters{},
			},
			want:    bool(123 == 345),
			wantErr: false,
		},
		{
			name: "a == b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 == 345",
				parameters: MapParameters{},
			},
			want:    bool(345 == 345),
			wantErr: false,
		},
		{
			name: "a == b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 == nil",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
		{
			name: "a == b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 != nil",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a == b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "\"ABC\" != \"AB\"",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a != b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123 != 345",
				parameters: MapParameters{},
			},
			want:    bool(123 != 345),
			wantErr: false,
		},
		{
			name: "a != b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 != 345",
				parameters: MapParameters{},
			},
			want:    bool(345 != 345),
			wantErr: false,
		},
		{
			name: "a > b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123 > 345",
				parameters: MapParameters{},
			},
			want:    bool(123 > 345),
			wantErr: false,
		},
		{
			name: "a > b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 > 345",
				parameters: MapParameters{},
			},
			want:    bool(345 > 345),
			wantErr: false,
		},
		{
			name: "a > b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 > 123",
				parameters: MapParameters{},
			},
			want:    bool(345 > 123),
			wantErr: false,
		},
		{
			name: "a >= b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123 >= 345",
				parameters: MapParameters{},
			},
			want:    bool(123 >= 345),
			wantErr: false,
		},
		{
			name: "a >= b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 >= 345",
				parameters: MapParameters{},
			},
			want:    bool(345 >= 345),
			wantErr: false,
		},
		{
			name: "a >= b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 >= 123",
				parameters: MapParameters{},
			},
			want:    bool(345 >= 123),
			wantErr: false,
		},
		{
			name: "a < b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123 < 345",
				parameters: MapParameters{},
			},
			want:    bool(123 < 345),
			wantErr: false,
		},
		{
			name: "a < b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 < 345",
				parameters: MapParameters{},
			},
			want:    bool(345 < 345),
			wantErr: false,
		},
		{
			name: "a < b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 < 123",
				parameters: MapParameters{},
			},
			want:    bool(345 < 123),
			wantErr: false,
		},
		{
			name: "a <= b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123 <= 345",
				parameters: MapParameters{},
			},
			want:    bool(123 <= 345),
			wantErr: false,
		},
		{
			name: "a <= b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 <= 345",
				parameters: MapParameters{},
			},
			want:    bool(345 <= 345),
			wantErr: false,
		},
		{
			name: "a <= b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "345 <= 123",
				parameters: MapParameters{},
			},
			want:    bool(345 <= 123),
			wantErr: false,
		},
		{
			name: "a && b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true && true",
				parameters: MapParameters{},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "a && b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true && false",
				parameters: MapParameters{},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "a && b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "false && false",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
		{
			name: "a && b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "1.0 && 0",
				parameters: MapParameters{},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "a || b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true || true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a || b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true || false",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a || b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "false || false",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
		{
			name: "a and b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true and true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a and b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true and false",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
		{
			name: "a and b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "false and false",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
		{
			name: "a or b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true or true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a or b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true or false",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a or b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "false or false",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
		{
			name: "a in b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123 in {123, \"dtr\",456}",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a in b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123 in {345, \"dtr\",456}",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
		{
			name: "a in b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "123 in {}",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
		{
			name: "a?b:c",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a1[0]?a1[1]:a1[2]",
				parameters: MapParameters{
					"a1": []string{
						"true",
						"b",
						"c",
						"d",
					},
				},
			},
			want:    "b",
			wantErr: false,
		},
		{
			name: "a?b:c",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true? 1.0 : 2",
				parameters: MapParameters{},
			},
			want:    1.0,
			wantErr: false,
		},
		{
			name: "a?b:c",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "1 ? 1.0 : 2",
				parameters: MapParameters{},
			},
			want:    1.0,
			wantErr: false,
		},
		{
			name: "a?b:c",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "0 ? 1.0 : 2",
				parameters: MapParameters{},
			},
			want:    int64(2),
			wantErr: false,
		},
		{
			name: "a?b:c",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a1[0]?a1[1]:a1[2]",
				parameters: MapParameters{
					"a1": []string{
						"",
						"b",
						"c",
						"d",
					},
				},
			},
			want:    "c",
			wantErr: false,
		},
		{
			name: "a?b:c",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true?a1[1]:a1[2]",
				parameters: MapParameters{
					"a1": []string{
						"",
						"b",
						"c",
						"d",
					},
				},
			},
			want:    "b",
			wantErr: false,
		},
		{
			name: "a1 * a2 - a2 + a2 * (a1 * a2 - a2 + a2)",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a1 * a2 - a2 + a2 * (a1 * a2 - a2 + a2)",
				parameters: MapParameters{
					"a1": 123,
					"a2": 345,
				},
			},
			want:    int64(123.0*345.0 - 345.0 + 345.0*(123.0*345.0-345.0+345.0)),
			wantErr: false,
		},
		{
			name: "a1.a2",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a1.A2",
				parameters: MapParameters{
					"a1": struct {
						A2 string
					}{
						A2: "345",
					},
				},
			},
			want:    "345",
			wantErr: false,
		},
		{
			name: "a1[1]",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a1[1]",
				parameters: MapParameters{
					"a1": []string{
						"345",
						"456",
					},
				},
			},
			want:    "456",
			wantErr: false,
		},
		{
			name: "a1[1]",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "a1[1:3]",
				parameters: MapParameters{
					"a1": []string{
						"a",
						"b",
						"c",
						"d",
					},
				},
			},
			want: []string{
				"b",
				"c",
			},
			wantErr: false,
		},
		{
			name: "{0,1,2}",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "{0,1,2}",
				parameters: MapParameters{
					"a1": []string{
						"",
						"b",
						"c",
						"d",
					},
				},
			},
			want:    []any{int64(0), int64(1), int64(2)},
			wantErr: false,
		},
		{
			name: "{0,1,2}[1]",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "{0,1,2}[1]",
				parameters: MapParameters{
					"a1": []string{
						"",
						"b",
						"c",
						"d",
					},
				},
			},
			want:    int64(1),
			wantErr: false,
		},
		{
			name: "{\"1\",\"2\",\"3\"}[1]",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "{\"1\",\"2\",\"3\"}[1]",
				parameters: MapParameters{
					"a1": 4,
				},
			},
			want:    "2",
			wantErr: false,
		},
		{
			name: "{\"1\",2,\"3\"}[1]",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "{\"1\",2,\"3\"}[1]",
				parameters: MapParameters{
					"a1": 4,
				},
			},
			want:    int64(2),
			wantErr: false,
		},
		{
			name: "{0,1,2}[1] > a1",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "{0,1,2}[1] > a1",
				parameters: MapParameters{
					"a1": 4,
				},
			},
			want:    false,
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ee := &EvaluableExpression{
				stage:       tt.fields.stage,
				ChecksTypes: tt.fields.ChecksTypes,
				IsDebug:     true,
			}
			got, err := ee.EvalString(tt.args.expression, tt.args.parameters)
			t.Logf("%s; got:%+v; type:%+v", strings.Join(ee.recordStep, "; "), got, reflect.TypeOf(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalString() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestEvaluableExpression_EvalString(t *testing.T) {
	type fields struct {
		stage       *evaluationNode
		ChecksTypes bool
		IsDebug     bool
	}
	type args struct {
		expression string
		parameters Parameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "true",
				parameters: MapParameters{},
			},
			want:    bool(true),
			wantErr: false,
		},
		{
			name: "a or b",
			fields: fields{
				stage:       nil,
				ChecksTypes: true,
			},
			args: args{
				expression: "false or false",
				parameters: MapParameters{},
			},
			want:    bool(false),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ee := &EvaluableExpression{
				stage:       tt.fields.stage,
				ChecksTypes: tt.fields.ChecksTypes,
				IsDebug:     true,
			}
			got, err := ee.EvalString(tt.args.expression, tt.args.parameters)
			t.Logf("%s; got:%+v; type:%+v", strings.Join(ee.recordStep, "; "), got, reflect.TypeOf(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalString() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}
