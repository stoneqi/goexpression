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
					symbol: 0,
					leftOperator: &evaluationNode{
						symbol:          LITERAL,
						leftOperator:    nil,
						rightOperator:   nil,
						operator:        makeAccessorOperator([]string{"a"}),
						leftTypeCheck:   nil,
						rightTypeCheck:  nil,
						typeCheck:       nil,
						typeErrorFormat: "",
					},
					rightOperator: &evaluationNode{
						symbol:          LITERAL,
						leftOperator:    nil,
						rightOperator:   nil,
						operator:        makeLiteralOperator(345.0),
						leftTypeCheck:   nil,
						rightTypeCheck:  nil,
						typeCheck:       nil,
						typeErrorFormat: "",
					},
					operator:        addOperator,
					leftTypeCheck:   isFloat64,
					rightTypeCheck:  isFloat64,
					typeCheck:       nil,
					typeErrorFormat: "",
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
