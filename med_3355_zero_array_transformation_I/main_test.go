package main

import "testing"

func Test_isZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1,0,2 -> 0,2 -> true",
			args: args{
				nums:    []int{1, 0, 2},
				queries: [][]int{{0, 2}, {1, 2}},
			},
			want: true,
		},
		{
			name: "4,3,2,1 -> [1,3] [0,2] -> false",
			args: args{
				nums:    []int{4, 3, 2, 1},
				queries: [][]int{{1, 3}, {0, 2}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("isZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
