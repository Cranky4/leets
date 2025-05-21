package main

import "testing"

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name  string
		args  args
		wants [][]int
	}{
		{
			name: "[1 1 1] [1 0 1] [1 1 1]",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			wants: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "[0 1 2 0] [3 4 5 2] [1 3 1 5]",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			wants: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)

			for row, cols := range tt.args.matrix {
				for col, val := range cols {
					if val != tt.wants[row][col] {
						t.Errorf("setZeroes() = %v, want %v", tt.args.matrix, tt.wants)
					}
				}
			}
		})
	}
}
