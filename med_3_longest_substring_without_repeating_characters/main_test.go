package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abcabcbb",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "bbbbb",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "pwwkew",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "au",
			args: args{
				s: "au",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "a",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "azxcc",
			args: args{
				s: "azxcc",
			},
			want: 4,
		},
		{
			name: "aazxc",
			args: args{
				s: "aazxc",
			},
			want: 4,
		},
		{
			name: "azzxc",
			args: args{
				s: "azzxc",
			},
			want: 3,
		},
		{
			name: "dvdf",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
