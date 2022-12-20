package main

import "fmt"

func main() {
	in := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}

	fmt.Printf("%v\n", canVisitAllRooms(in))
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make(map[int]struct{}, n)

	q := newQ(n)
	visited[0] = struct{}{}
	for _, key := range rooms[0] {
		q.Push(key)
	}

	for {
		if q.Len() == 0 {
			break
		}

		n := q.Pop()

		_, ok := visited[n]
		if ok {
			continue
		}

		visited[n] = struct{}{}

		for _, key := range rooms[n] {
			q.Push(key)
		}
	}

	return len(visited) == n
}

type queue struct {
	items []int
}

func newQ(n int) *queue {
	return &queue{
		items: make([]int, 0, n),
	}
}

func (q *queue) Push(n int) {
	q.items = append(q.items, n)
}

func (q *queue) Len() int {
	return len(q.items)
}

func (q *queue) Pop() int {
	n := q.items[0]

	q.items = q.items[1:]

	return n
}
