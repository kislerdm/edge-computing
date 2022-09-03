package main

import (
	"testing"
	"unsafe"
)

func Test_t(ts *testing.T) {
	type args struct {
		r, g, b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path: warm",
			args: args{254, 37, 0},
			want: true,
		},
		{
			name: "happy path: cold",
			args: args{182, 221, 199},
			want: false,
		},
	}
	for _, tt := range tests {
		ts.Run(
			tt.name, func(ts *testing.T) {
				if got := t(tt.args.r, tt.args.g, tt.args.b); got != tt.want {
					ts.Errorf("t() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_n(t *testing.T) {
	type args struct {
		r, g, b int
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
				n(tt.args.r, tt.args.g, tt.args.b)
				if &name[0] != getNAddress() {
					t.Fatal("getNAddress() failed")
				}
				if len(name) != getNLen() {
					t.Fatal("getNLen() failed")
				}

				got := *(*string)(unsafe.Pointer(&name))
				if got != tt.want {
					t.Errorf("n() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
