package main

import "fmt"

func main() {
	in := &Node{
		Val: 1, Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}

	fmt.Printf("%v\n", preorder(in))
}

func preorder(root *Node) []int {
	stack := NewStack(10)
	stack.Add(root)
	res := make([]int, 0, 10)

	if root == nil {
		return []int{}
	}

	for stack.Len() > 0 {
		item := stack.Get()
		res = append(res, item.Val)

		for i := len(item.Children) - 1; i >= 0; i-- {
			stack.Add(item.Children[i])
		}
	}

	return res
}

type stack struct {
	items []*Node
}

func NewStack(n int) stack {
	return stack{
		items: make([]*Node, 0, n),
	}
}

func (s *stack) Add(i *Node) {
	s.items = append(s.items, i)
}

func (s *stack) Get() *Node {
	res := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return res
}

func (s *stack) Len() int {
	return len(s.items)
}

type Node struct {
	Val      int
	Children []*Node
}
