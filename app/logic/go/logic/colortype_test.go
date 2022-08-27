package logic

import "testing"

func TestType(t *testing.T) {
	type args struct {
		r, g, b uint8
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "happy path: warm",
			args:    args{254, 37, 0},
			want:    true,
			wantErr: false,
		},
		{
			name:    "happy path: cold",
			args:    args{182, 221, 199},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := Type(tt.args.r, tt.args.g, tt.args.b)
				if (err != nil) != tt.wantErr {
					t.Errorf("Type() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Type() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
