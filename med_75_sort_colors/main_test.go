package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name   string
		args   []int
		wanted []int
	}{
		{
			name:   "[2,0,2,1,1,0]",
			args:   []int{2, 0, 2, 1, 1, 0},
			wanted: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:   "[2,0,1]",
			args:   []int{2, 0, 1},
			wanted: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args)

			assert.Equal(t, tt.wanted, tt.args)
		})
	}
}
