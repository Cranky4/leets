package main

import "fmt"

func main() {
	in1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	in2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	fmt.Printf("%v\n", leafSimilar(in1, in2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	nodes := make([][]int, 2)
	s := NewStack(10)

	for i, root := range []*TreeNode{root1, root2} {
		s.Add(root)
		for {
			if s.Len() == 0 {
				break
			}
			n := s.Get()

			if n.Left == nil && n.Right == nil {
				nodes[i] = append(nodes[i], n.Val)
				continue
			}

			if n.Left != nil {
				s.Add(n.Left)
			}

			if n.Right != nil {
				s.Add(n.Right)
			}
		}
	}

	if len(nodes[0]) != len(nodes[1]) {
		return false
	}

	for i := 0; i < len(nodes[0]); i++ {
		if nodes[0][i] != nodes[1][i] {
			return false
		}
	}

	return true
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

	n := s.items[s.Len()-1]

	s.items = s.items[:s.Len()-1]

	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
