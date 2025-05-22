package main

import "testing"

func Test_repairCars(t *testing.T) {
	type args struct {
		ranks []int
		cars  int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "ranks = [4,2,3,1], cars = 10",
			args: args{
				ranks: []int{4, 2, 3, 1},
				cars:  10,
			},
			want: 16,
		},
		{
			name: "ranks = [5,1,8], cars = 6",
			args: args{
				ranks: []int{5, 1, 8},
				cars:  6,
			},
			want: 16,
		},
		{
			name: "ranks = [3,3,1,2,1,1,3,2,1], cars = 58",
			args: args{
				ranks: []int{3, 3, 1, 2, 1, 1, 3, 2, 1},
				cars:  58,
			},
			want: 75,
		},
		{
			name: "ranks = [1,3,1,2,1,1], cars = 21",
			args: args{
				ranks: []int{1, 3, 1, 2, 1, 1},
				cars:  21,
			},
			want: 18,
		},
		{
			name: "ranks = [31,31,5,19,19,10,31,18,19,3,16,20,4,16,2,25,10,16,23,18,21,23,28,6,7,29,11,11,19,20,24,19,26,12,29,29,1,14,17,26,24,7,11,28,22,14,31,12,3,19,16,26,11], cars = 736185",
			args: args{
				ranks: []int{31, 31, 5, 19, 19, 10, 31, 18, 19, 3, 16, 20, 4, 16, 2, 25, 10, 16, 23, 18, 21, 23, 28, 6, 7, 29, 11, 11, 19, 20, 24, 19, 26, 12, 29, 29, 1, 14, 17, 26, 24, 7, 11, 28, 22, 14, 31, 12, 3, 19, 16, 26, 11},
				cars:  736185,
			},
			want: 2358388332,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repairCars(tt.args.ranks, tt.args.cars); got != tt.want {
				t.Errorf("repairCars() = %v, want %v", got, tt.want)
			}
		})
	}
}
