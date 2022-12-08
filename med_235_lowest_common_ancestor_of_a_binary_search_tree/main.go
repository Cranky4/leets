package main

import "fmt"

func main() {
	q := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	p := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 0,
		},
		Right: q,
	}

	in := &TreeNode{
		Val:  6,
		Left: p,
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Printf("%v\n", lowestCommonAncestor(in, p, q))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	s := newStack(100)
	s.Add(root)

	res := root

	for {
		if s.Len() == 0 {
			break
		}

		n := s.Get()

		if p.Val <= n.Val && q.Val >= n.Val {
			res = n
			break
		}

		if p.Val < n.Val && q.Val < n.Val && n.Left != nil {
			res = n.Left
			s.Add(n.Left)

			continue
		}

		if p.Val > n.Val && q.Val > n.Val && n.Right != nil {
			res = n.Right
			s.Add(n.Right)

			continue
		}
	}

	return res
}

type stack struct {
	items []*TreeNode
}

func newStack(n int) *stack {
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
	n := s.items[0]

	s.items = s.items[1:]

	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
