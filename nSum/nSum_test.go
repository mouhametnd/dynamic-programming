package main
import "testing"

func Test_nSumRecursive(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nSumRecursive(tt.args.n); got != tt.want {
				t.Errorf("nSumRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "edge case #1"
			args:{
				n:0
			},
			want: 0
		},
		{
			name: "edge case #2"
			args:{
				n:1
			},
			want: 1
		},
		{
			name: "simple test #1"
			args:{
				n:5
			},
			want: 15
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nSum(tt.args.n); got != tt.want {
				t.Errorf("nSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
