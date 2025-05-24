package main

import (
	"reflect"
	"testing"
)

func Test_findWordsContaining(t *testing.T) {
	type args struct {
		words []string
		x     byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "words = [\"leet\",\"code\"], x = \"e\"",
			args: args{
				words: []string{"leet", "code"},
				x:     'e',
			},
			want: []int{0, 1},
		},
		{
			name: "words = [\"abc\",\"bcd\",\"aaaa\",\"cbc\"], x = \"a\"",
			args: args{
				words: []string{"abc", "bcd", "aaaa", "cbc"},
				x:     'a',
			},
			want: []int{0, 2},
		},
		{
			name: "words = [\"abc\",\"bcd\",\"aaaa\",\"cbc\"], x = \"z\"",
			args: args{
				words: []string{"abc", "bcd", "aaaa", "cbc"},
				x:     'z',
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWordsContaining(tt.args.words, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
