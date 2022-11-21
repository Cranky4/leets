package main

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Breadth-first_search

func bfs(start [2]int, field [][]byte) int {
	res := -1

	rows := len(field)
	if rows == 0 {
		return res
	}

	cols := len(field[0])
	if cols == 0 {
		return res
	}

	visited := make(map[string]struct{}, rows*cols)
	visited[fmt.Sprintf("%d_%d", start[0], start[1])] = struct{}{}

	queue := walk(newQueue(rows*cols), [3]int{start[0], start[1], 0}, rows, cols)

	for {
		if queue.Len() == 0 {
			break
		}

		p, _ := queue.Get()

		key := fmt.Sprintf("%d_%d", p[0], p[1])
		_, ok := visited[key]
		if ok || isWall(field[p[0]][p[1]]) {
			continue
		}

		if isExit(rows, cols, p) {
			res = p[2]
			break
		}

		visited[key] = struct{}{}

		queue = walk(queue, p, rows, cols)
	}

	return res
}

func isExit(rowsCount, colsCount int, point [3]int) bool {
	return point[0] == rowsCount-1 || point[0] == 0 || point[1] == 0 || point[1] == colsCount-1
}

func isWall(p byte) bool {
	return p == 43
	// 46 .
	// 43 +
}

// basic queue
type queue struct {
	items [][3]int
}

func newQueue(n int) queue {
	return queue{
		items: make([][3]int, 0, n), // row, col, depth
	}
}

func (q *queue) Add(point [3]int) {
	q.items = append(q.items, point)
}

func (q *queue) Len() int {
	return len(q.items)
}

func (q *queue) Get() ([3]int, bool) {
	if len(q.items) == 0 {
		return [3]int{0, 0, 0}, false
	}

	first := q.items[0]

	q.items = q.items[1:]

	return first, true
}

func walk(queue queue, p [3]int, rows, cols int) queue {
	// up
	if p[0]-1 >= 0 {
		queue.Add([3]int{p[0] - 1, p[1], p[2] + 1})
	}

	// right
	if p[1]+1 < cols {
		queue.Add([3]int{p[0], p[1] + 1, p[2] + 1})
	}

	// down
	if p[0]+1 < rows {
		queue.Add([3]int{p[0] + 1, p[1], p[2] + 1})
	}

	// left
	if p[1]-1 >= 0 {
		queue.Add([3]int{p[0] - 1, p[1], p[2] + 1})
	}

	return queue
}
