package parserSecond

import (
	"reflect"
	"testing"
)

func TestEvaluableExpression_evaluateStage(t *testing.T) {
	type fields struct {
		ChecksTypes bool
	}
	type args struct {
		stage      *evaluationNode
		parameters Parameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				ChecksTypes: true,
			},
			args: args{
				stage: &evaluationNode{
					Symbol: 0,
					LeftOperator: &evaluationNode{
						Symbol:          LITERAL,
						LeftOperator:    nil,
						RightOperator:   nil,
						Operator:        makeAccessorOperator([]string{"a"}),
						LeftTypeCheck:   nil,
						RightTypeCheck:  nil,
						TypeCheck:       nil,
						TypeErrorFormat: "",
					},
					RightOperator: &evaluationNode{
						Symbol:          LITERAL,
						LeftOperator:    nil,
						RightOperator:   nil,
						Operator:        makeLiteralOperator(345.0),
						LeftTypeCheck:   nil,
						RightTypeCheck:  nil,
						TypeCheck:       nil,
						TypeErrorFormat: "",
					},
					Operator:        addOperator,
					LeftTypeCheck:   isFloat64,
					RightTypeCheck:  isFloat64,
					TypeCheck:       nil,
					TypeErrorFormat: "",
				},
				parameters: MapParameters{
					"a": 123,
				},
			},
			want:    468.0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := EvaluableExpression{
				ChecksTypes: tt.fields.ChecksTypes,
			}
			got, err := this.evaluateStage(tt.args.stage, tt.args.parameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("evaluateStage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evaluateStage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvaluableExpression_EvalString(t *testing.T) {
	type fields struct {
		stage       *evaluationNode
		ChecksTypes bool
	}
	type args struct {
		expression string
		parameters Parameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
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
			want:    float64(0x123 + 345),
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ee := &EvaluableExpression{
				stage:       tt.fields.stage,
				ChecksTypes: tt.fields.ChecksTypes,
			}
			got, err := ee.EvalString(tt.args.expression, tt.args.parameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
