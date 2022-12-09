package main

import (
	"fmt"
	"math"
)

func main() {
	in := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 10,
			Right: &TreeNode{
				Val: 14,
				Left: &TreeNode{
					Val: 13,
				},
			},
		},
	}

	fmt.Printf("%v\n", maxAncestorDiff(in))
}

func maxAncestorDiff(root *TreeNode) int {
	return walk(root, math.MaxInt, 0)
}

func walk(n *TreeNode, min, max int) int {
	if n == nil {
		return max - min
	}

	if n.Val > max {
		max = n.Val
	}

	if n.Val < min {
		min = n.Val
	}

	return maximum(
		maximum(max-min, walk(n.Left, min, max)),
		maximum(max-min, walk(n.Right, min, max)),
	)
}

func maximum(i, j int) int {
	if i > j {
		return i
	}

	return j
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
