package logic_test

import (
	"reflect"
	"testing"

	"edgecomputing/logic"
)

func TestStart(t *testing.T) {
	type args struct {
		r, g, b uint8
	}
	tests := []struct {
		name    string
		args    args
		want    logic.Output
		wantErr bool
	}{
		{
			name: "happy path: Black",
			args: args{0, 0, 0},
			want: logic.Output{
				Name:   "Black",
				IsWarm: false,
			},
			wantErr: false,
		},
		{
			name: "happy path: Red",
			args: args{255, 0, 0},
			want: logic.Output{
				Name:   "Red",
				IsWarm: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := logic.Start(tt.args.r, tt.args.g, tt.args.b)
				if (err != nil) != tt.wantErr {
					t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Start() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
