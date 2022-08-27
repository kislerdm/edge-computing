package logic

import "testing"

func TestColorName(t *testing.T) {
	type args struct {
		r, g, b uint8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "exact name",
			args: args{
				r: 124,
				g: 185,
				b: 232,
			},
			want: "Aero",
		},
		{
			name: "black",
			args: args{
				r: 0,
				g: 0,
				b: 0,
			},
			want: "Black",
		},
		{
			name: "white",
			args: args{
				r: 255,
				g: 255,
				b: 255,
			},
			want: "White",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Name(tt.args.r, tt.args.g, tt.args.b); got != tt.want {
					t.Errorf("Name() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
