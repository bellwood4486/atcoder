package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{6, []int{0, 153, 10, 10, 23}}, 53},
		{"", args{2, []int{3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N, tt.args.B); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
