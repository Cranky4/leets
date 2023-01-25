package main

import (
	"fmt"
	"math"
)

func main() {
	edges := []int{2, 2, 3, -1}
	p1 := 0
	p2 := 1

	fmt.Printf("%v\n", closestMeetingNode(edges, p1, p2))
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	visited := make(map[string]struct{})
	q1 := newQ(len(edges))
	q1.Push(node1)

	q2 := newQ(len(edges))
	q2.Push(node2)

	res := [2]int{math.MaxInt, math.MaxInt}

	for q1.Len() > 0 || q2.Len() > 0 {
		step := [2]int{-1, -1}

		if q1.Len() > 0 {
			step[0] = q1.Pop()
		}

		if q2.Len() > 0 {
			step[1] = q2.Pop()
		}

		// check visited
		_, ok1 := visited[fmt.Sprintf("%d_", step[0])]
		if !ok1 {
			visited[fmt.Sprintf("%d_", step[0])] = struct{}{}

			// check meet
			_, ok := visited[fmt.Sprintf("_%d", step[0])]
			if ok {
				res[0] = step[0]
			} else if step[0] != -1 {
				// go next
				if edges[step[0]] != -1 {
					q1.Push(edges[step[0]])
				}
			}
		}

		_, ok2 := visited[fmt.Sprintf("_%d", step[1])]
		if !ok2 {
			visited[fmt.Sprintf("_%d", step[1])] = struct{}{}

			// check meet
			_, ok := visited[fmt.Sprintf("%d_", step[1])]
			if ok {
				res[1] = step[1]
			} else if step[1] != -1 {
				// go next
				if edges[step[1]] != -1 {
					q2.Push(edges[step[1]])
				}
			}
		}

		if res[0] != math.MaxInt || res[1] != math.MaxInt {
			break
		}
	}

	if res[0] < res[1] {
		return res[0]
	}

	if res[1] == math.MaxInt {
		return -1
	}

	return res[1]
}

type queue struct {
	items []int
}

func newQ(n int) *queue {
	return &queue{
		items: make([]int, 0, n),
	}
}

func (q *queue) Push(i int) {
	q.items = append(q.items, i)
}

func (q *queue) Len() int {
	return len(q.items)
}

func (q *queue) Pop() int {
	i := q.items[0]

	q.items = q.items[1:]

	return i
}
