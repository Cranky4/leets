package main

import (
	"log"
	"math"
)

// https://ru.wikipedia.org/wiki/%D0%9E%D0%B1%D1%80%D0%B0%D1%82%D0%BD%D0%B0%D1%8F_%D0%BF%D0%BE%D0%BB%D1%8C%D1%81%D0%BA%D0%B0%D1%8F_%D0%B7%D0%B0%D0%BF%D0%B8%D1%81%D1%8C

func rpn_calc(rpn string) int {
	stack := NewRPNStack(len(rpn))

	for _, v := range rpn {
		n1 := -1
		n2 := -1
		res := -1

		if v == '+' || v == '-' || v == '*' || v == '/' || v == '^' {
			for {
				n, ok := stack.Pop()

				if !ok {
					log.Fatalf("unexpected empty stack")
				}

				if n2 == -1 {
					n2 = n
				} else if n1 == -1 {
					n1 = n
				}

				if n1 != -1 && n2 != -1 {
					switch v {
					case '+':
						res = n1 + n2
					case '-':
						res = n1 - n2
					case '*':
						res = n1 * n2
					case '/':
						res = n1 / n2
					}

					stack.Add(res)

					break
				}
			}
		} else {
			stack.Add(int(v - '0'))
		}
	}

	if stack.Len() == 1 {
		r, _ := stack.Pop()

		return r
	}

	r := 0
	pow := 0

	for {
		v, ok := stack.Pop()

		if !ok {
			break
		}

		r += v * int(math.Pow10(pow))
		pow++
	}

	return r
}

// basic stack

func NewRPNStack(n int) rpn_stack {
	return rpn_stack{
		items: make([]int, 0, n),
	}
}

type rpn_stack struct {
	items []int
}

func (s *rpn_stack) Top() (int, bool) {
	if s.Len() == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

func (s *rpn_stack) Len() int {
	return len(s.items)
}

func (s *rpn_stack) Add(str int) {
	s.items = append(s.items, str)
}

func (s *rpn_stack) Pop() (int, bool) {
	l := len(s.items)

	if l == 0 {
		return 0, false
	}

	last := s.items[l-1]

	s.items = s.items[:l-1]

	return last, true
}
