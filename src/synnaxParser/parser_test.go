// Package main provides ...
package parserSecond

import (
	"testing"
)

func TestParserString(t *testing.T) {
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
			name: "1323",
			args: args{
				expression: "a + b",
			},
			want:    "a+b",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParserString(tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParserString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParserString() = %v, want %v", got, tt.want)
			}
		})
	}
}
