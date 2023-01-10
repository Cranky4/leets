package main

import "fmt"

func main() {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Printf("%v\n", isSameTree(p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	q1 := NewQ(10)
	q2 := NewQ(10)

	q1.Push(p)
	q2.Push(q)

	for {
		n1 := q1.Pop()
		n2 := q2.Pop()

		if n1 == nil && n2 != nil ||
			n2 == nil && n1 != nil {
			return false
		}

		if n1 == nil && n2 == nil {
			break
		}

		if n1.Val != n2.Val ||
			(n1.Left == nil && n2.Left != nil) ||
			(n1.Left != nil && n2.Left == nil) ||
			(n1.Left != nil && n2.Left != nil && n1.Left.Val != n2.Left.Val) ||
			(n1.Right == nil && n2.Right != nil) ||
			(n1.Right != nil && n2.Right == nil) ||
			(n1.Right != nil && n2.Right != nil && n1.Right.Val != n2.Right.Val) {
			return false
		}

		if n1.Left != nil {
			q1.Push(n1.Left)
		}
		if n2.Right != nil {
			q1.Push(n1.Right)
		}

		if n2.Left != nil {
			q2.Push(n2.Left)
		}
		if n2.Right != nil {
			q2.Push(n2.Right)
		}
	}

	return true
}

type queue struct {
	items []*TreeNode
}

func NewQ(n int) *queue {
	return &queue{
		items: make([]*TreeNode, 0, n),
	}
}

func (q *queue) Push(n *TreeNode) {
	q.items = append(q.items, n)
}

func (q *queue) Len() int {
	return len(q.items)
}

func (q *queue) Pop() *TreeNode {
	if q.Len() == 0 {
		return nil
	}

	n := q.items[0]

	q.items = q.items[1:]

	return n
}
