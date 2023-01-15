package parserSecond

import (
	"encoding/json"
	"testing"
)

func TestVisitorParserString(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "add",
			args: args{
				expression: "abc + aa",
			},
			want:    "123",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VisitorParserString(tt.args.expression)
			t.Errorf("%+v", *got)
			gotStr, err := json.Marshal(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("VisitorParserString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(gotStr) != tt.want {
				t.Errorf("VisitorParserString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
