package main

import "fmt"

type queue struct {
	items []item
}

type item struct {
	n int
	r rune
}

func NewQueue(l int) queue {
	return queue{
		items: make([]item, 0, l),
	}
}

func (q *queue) Add(i item) {
	q.items = append(q.items, i)

	q.shiftUp(len(q.items) - 1)
}

func (q *queue) Get() item {
	if len(q.items) == 0 {
		return item{n: -1}
	}

	res := q.items[0]

	q.Swap(0, len(q.items)-1)
	q.items = q.items[:len(q.items)-1]
	q.shiftDown(0)

	return res
}

func (q *queue) shiftDown(idx int) {
	left := q.getLeft(idx)
	right := q.getRight(idx)

	if left == -1 && right == -1 {
		return
	}

	// shift down to left
	if (left != -1 && right != -1 && q.items[left].n >= q.items[right].n) || (right == -1 && left != -1) {
		if q.items[idx].n < q.items[left].n {
			q.Swap(idx, left)
			q.shiftDown(left)
		}

		return
	}

	// shift down to right
	if (left != -1 && right != -1 && q.items[left].n < q.items[right].n) || (right != -1 && left == -1) {
		if q.items[idx].n < q.items[right].n {
			q.Swap(idx, right)
			q.shiftDown(right)
		}

		return
	}

	fmt.Printf("f %v\n", idx)
}

func (q *queue) shiftUp(idx int) {
	parent := q.getParent(idx)
	if parent == -1 {
		return
	}

	if q.items[idx].n <= q.items[parent].n {
		return
	}

	q.Swap(idx, parent)
	q.shiftUp(parent)
}

func (q *queue) Swap(idx1, idx2 int) {
	q.items[idx1], q.items[idx2] = q.items[idx2], q.items[idx1]
}

func (q *queue) getParent(idx int) int {
	parent := int((idx - 1) / 2)

	if parent >= 0 {
		return parent
	}

	return -1
}

func (q *queue) getRight(idx int) int {
	right := idx*2 + 2

	if len(q.items)-1 >= right {
		return right
	}
	return -1
}

func (q *queue) getLeft(idx int) int {
	left := idx*2 + 1

	if len(q.items)-1 >= left {
		return left
	}
	return -1
}
