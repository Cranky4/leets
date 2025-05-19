package main

import "testing"

func Test_shiftingLetters(t *testing.T) {
	type args struct {
		s      string
		shifts [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abc->ace",
			args: args{
				s: "abc",
				shifts: [][]int{
					{0, 1, 0},
					{1, 2, 1},
					{0, 2, 1},
				},
			},
			want: "ace",
		},
		{
			name: "double shift over the alphabet - forward",
			args: args{
				s: "z",
				shifts: [][]int{
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
					{0, 0, 1},
				},
			},
			want: "a",
		},
		{
			name: "double shift over the alphabet - backward",
			args: args{
				s: "a",
				shifts: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: "z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftingLetters(tt.args.s, tt.args.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
