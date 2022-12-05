package main

import "fmt"

func main() {
	in := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Printf("%v\n", levelOrder(in))
}

func levelOrder(root *TreeNode) [][]int {
	q := NewQueue(1)
	q.Add(&QueueItem{Node: root, Level: 0})
	res := make([][]int, 0, 1)

	if root == nil {
		return res
	}

	l := 0
	r := make([]int, 0, 2)

	for {
		if q.Len() == 0 {
			break
		}
		n := q.Get()

		if n.Level == l {
			r = append(r, n.Node.Val)
		} else {
			l = n.Level
			res = append(res, r)
			r = make([]int, 0, 2)
			r = append(r, n.Node.Val)
		}

		if n.Node.Left != nil {
			q.Add(&QueueItem{Node: n.Node.Left, Level: n.Level + 1})
		}

		if n.Node.Right != nil {
			q.Add(&QueueItem{Node: n.Node.Right, Level: n.Level + 1})
		}
	}

	if len(r) > 0 {
		res = append(res, r)
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queue struct {
	items []*QueueItem
}

func NewQueue(n int) queue {
	return queue{
		items: make([]*QueueItem, 0, n),
	}
}

func (q *queue) Add(i *QueueItem) {
	q.items = append(q.items, i)
}

func (q *queue) Len() int {
	return len(q.items)
}

func (q *queue) Get() *QueueItem {
	res := q.items[0]
	q.items = q.items[1:]
	return res
}

type QueueItem struct {
	Node  *TreeNode
	Level int
}
