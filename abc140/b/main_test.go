package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N int
		A []int
		B []int
		C []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{3, []int{3, 1, 2}, []int{2, 5, 4}, []int{3, 6}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N, tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
