package main

import "fmt"

func main() {
	in := [][]int{
		{0, 0, 1}, {1, 0, 0}, {0, 0, 0},
	}

	fmt.Println(highestPeak(in))
}

/**
      0 0 1
      1 0 0
	  0 0 0

	  1 1 0
	  0 1 1
	  1 2 2
*/

func highestPeak(isWater [][]int) [][]int {
	if len(isWater) == 0 {
		return nil
	}

	totalRows := len(isWater)
	totalCols := len(isWater[0])

	// init
	result := make([][]int, totalRows)

	for i := 0; i < len(isWater); i++ {
		result[i] = make([]int, totalCols)
	}

	q := newQueue(10) //nwm

	for row, cols := range result {
		for col := range cols {
			if isWater[row][col] == 1 {
				result[row][col] = 0
				continue
			}

			pushPoint(q, row, col, totalRows, totalCols, 1, -1, -1)

			// read
			for {
				item, ok := q.get()
				if !ok {
					break
				}

				if isWater[item.r][item.c] == 1 {
					result[row][col] = item.d
					// found the nearest water so flush the queue
					q.flush()
					break
				}

				pushPoint(q, item.r, item.c, totalRows, totalCols, item.d+1, row, col)
			}
		}
	}

	return result
}

func pushPoint(q *queue, row, col, rows, cols, depth, prevRow, prevCol int) {
	for _, coords := range getPaths() {
		// skip out of range and visited
		if row+coords[0] < 0 || row+coords[0] >= rows || col+coords[1] < 0 || col+coords[1] >= cols || prevRow == row+coords[0] && prevCol == col+coords[1] {
			continue
		}

		q.push(&queueItem{r: row + coords[0], c: col + coords[1], d: depth})
	}
}

func getPaths() [][2]int {
	return [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
}

type queue struct {
	items []*queueItem
}

type queueItem struct {
	r, c, d int
}

func newQueue(l int) *queue {
	return &queue{
		items: make([]*queueItem, 0, l),
	}
}

func (q *queue) push(item *queueItem) {
	q.items = append(q.items, item)
}

func (q *queue) get() (*queueItem, bool) {
	if len(q.items) == 0 {
		return nil, false
	}

	item := q.items[0]

	q.items = q.items[1:]

	return item, true
}

func (q *queue) flush() {
	q.items = q.items[:0]
}
