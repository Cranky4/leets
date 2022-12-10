package main

import "fmt"

func main() {
	in := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Printf("%v\n", maxProduct(in))
}

func maxProduct(root *TreeNode) int {
	m := make(map[*TreeNode]int)
	total := walk(root, m)

	return walkForMax(root, m, total, 0) % 1000000007
}

func walk(n *TreeNode, m map[*TreeNode]int) int {
	if n == nil {
		return 0
	}

	sum := n.Val
	sum += walk(n.Left, m)
	sum += walk(n.Right, m)

	m[n] = sum

	return sum
}

func walkForMax(n *TreeNode, m map[*TreeNode]int, total, max int) int {
	if n == nil {
		return max
	}

	max = maximum((total-m[n])*m[n], max)

	max = walkForMax(n.Left, m, total, max)
	max = walkForMax(n.Right, m, total, max)

	return max
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
