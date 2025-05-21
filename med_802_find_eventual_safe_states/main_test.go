package main

import (
	"reflect"
	"testing"
)

func Test_eventualSafeNodes(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[[],[0,2,3,4],[3],[4],[]]",
			args: args{
				graph: [][]int{
					{},
					{0, 2, 3, 4},
					{3},
					{4},
					{},
				},
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eventualSafeNodes(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventualSafeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
