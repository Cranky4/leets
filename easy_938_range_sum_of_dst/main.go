package main

import "fmt"

func main() {
	in := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	low := 7
	high := 15

	fmt.Printf("%v\n", rangeSumBST(in, low, high))
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	s := NewStack(10)
	s.Add(root)

	res := 0

	for {
		if s.Len() == 0 {
			break
		}

		n := s.Get()
		if n.Val >= low && n.Val <= high {
			res += n.Val
		}

		if n.Left != nil {
			s.Add(n.Left)
		}
		if n.Right != nil {
			s.Add(n.Right)
		}
	}

	return res
}

type stack struct {
	items []*TreeNode
}

func NewStack(n int) *stack {
	return &stack{
		items: make([]*TreeNode, 0, n),
	}
}

func (s *stack) Add(n *TreeNode) {
	s.items = append(s.items, n)
}

func (s *stack) Len() int {
	return len(s.items)
}

func (s *stack) Get() *TreeNode {
	if s.Len() == 0 {
		return nil
	}

	n := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
