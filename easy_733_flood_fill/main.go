package main

import "fmt"

func main() {
	in := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr := 1
	sc := 1
	color := 2

	fmt.Printf("%v\n", floodFill(in, sr, sc, color))
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows := len(image)
	if rows == 0 {
		return image
	}
	cols := len(image[0])

	q := newQueue(cols * rows)
	q.Add([2]int{sr, sc})

	orig := image[sr][sc]

	for {
		if q.Len() == 0 {
			break
		}

		n := q.Get()

		if image[n[0]][n[1]] == color || image[n[0]][n[1]] != orig {
			continue
		}

		// color
		image[n[0]][n[1]] = color

		// top
		if n[0]-1 >= 0 {
			q.Add([2]int{n[0] - 1, n[1]})
		}

		// right
		if n[1]+1 < cols {
			q.Add([2]int{n[0], n[1] + 1})
		}

		// bot
		if n[0]+1 < rows {
			q.Add([2]int{n[0] + 1, n[1]})
		}

		// left
		if n[1]-1 >= 0 {
			q.Add([2]int{n[0], n[1] - 1})
		}
	}

	return image
}

type queue struct {
	items [][2]int
}

func newQueue(n int) *queue {
	return &queue{
		items: make([][2]int, 0, n),
	}
}

func (q *queue) Add(n [2]int) {
	q.items = append(q.items, n)
}

func (q *queue) Len() int {
	return len(q.items)
}

func (q *queue) Get() [2]int {
	n := q.items[0]
	q.items = q.items[1:]

	return n
}
