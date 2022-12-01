package main

import "testing"

func Test_clmbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "test #1",
			args: args{n: 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clmbStairs(tt.args.n); got != tt.want {
				t.Errorf("clmbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
