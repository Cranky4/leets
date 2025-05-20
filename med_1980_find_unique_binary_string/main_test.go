package main

import "testing"

func Test_findDifferentBinaryString(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "\"01\",\"10\"",
			args: args{
				nums: []string{"01", "10"},
			},
			want: "00",
		},
		{
			name: "\"00\",\"01\"",
			args: args{
				nums: []string{"00", "01"},
			},
			want: "10",
		},
		{
			name: "\"111\",\"011\",\"001\"",
			args: args{
				nums: []string{"111", "011", "001"},
			},
			want: "000",
		},
		{
			name: "\"111\",\"011\",\"001\",\"000\", \"100\"",
			args: args{
				nums: []string{"111", "011", "001", "000", "100"},
			},
			want: "010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifferentBinaryString(tt.args.nums); got != tt.want {
				t.Errorf("findDifferentBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
