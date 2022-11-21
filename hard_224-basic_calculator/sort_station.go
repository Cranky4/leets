package main

import (
	"log"
)

// https://ru.wikipedia.org/wiki/%D0%90%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC_%D1%81%D0%BE%D1%80%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%BE%D1%87%D0%BD%D0%BE%D0%B9_%D1%81%D1%82%D0%B0%D0%BD%D1%86%D0%B8%D0%B8

func sort_station(str string) string {
	out_q := make([]rune, 0, len(str))
	stack := NewSortStationStack(len(str))

	for _, v := range str {
		// Если токен — число, то добавить его в очередь вывода.
		if v >= '0' && v <= '9' {
			out_q = append(out_q, v)
		} else
		// Если символ является открывающей скобкой, помещаем его в стек.
		if v == '(' {
			stack.Add(v)
		} else
		// Если символ является закрывающей скобкой
		if v == ')' {
			// До тех пор, пока верхним элементом стека не станет открывающая скобка, выталкиваем элементы из стека в выходную строку
			for {
				s, ok := stack.Pop()

				if !ok {
					log.Fatal("unexpected empty stack")
				}

				if s == '(' {
					break
				}

				out_q = append(out_q, s)
			}
		} else
		// Если символ является бинарной операцией о1
		if v == '+' || v == '-' || v == '*' || v == '/' || v == '^' {
			top, ok := stack.Top()

			// операция на вершине стека приоритетнее или такого же уровня приоритета как o1
			// выталкиваем ее из стека в выходную строку
			if ok && ((v == '+' || v == '-') && (top == '+' || top == '-') || (top == '*' || top == '/' || top == '^')) {
				s, _ := stack.Pop()

				out_q = append(out_q, s)
			}

			stack.Add(v)
		}
	}

	for stack.Len() > 0 {
		v, _ := stack.Pop()
		out_q = append(out_q, v)
	}

	return string(out_q)
}

func NewSortStationStack(n int) sort_station_stack {
	return sort_station_stack{
		items: make([]rune, 0, n),
	}
}

// basic stack
type sort_station_stack struct {
	items []rune
}

func (s *sort_station_stack) Top() (rune, bool) {
	if s.Len() == 0 {
		return '\n', false
	}

	return s.items[len(s.items)-1], true
}

func (s *sort_station_stack) Len() int {
	return len(s.items)
}

func (s *sort_station_stack) Add(str rune) {
	s.items = append(s.items, str)
}

func (s *sort_station_stack) Pop() (rune, bool) {
	l := len(s.items)

	if l == 0 {
		return '\n', false
	}

	last := s.items[l-1]

	s.items = s.items[:l-1]

	return last, true
}
