package main

import (
	"fmt"
	"strconv"
)

func main() {
	// in := []string{"2", "1", "+", "3", "*"}
	in := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	fmt.Printf("%v\n", evalRPN(in))
}

func evalRPN(tokens []string) int {
	l := len(tokens)
	st := newStack(l)

	for _, s := range tokens {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			n2 := st.Pop()
			n1 := st.Pop()
			res := 0

			switch s {
			case "+":
				res = n1 + n2
			case "-":
				res = n1 - n2
			case "*":
				res = n1 * n2
			case "/":
				res = n1 / n2
			}

			st.Push(res)
		} else {
			i, _ := strconv.Atoi(s)
			st.Push(i)
		}
	}

	return st.Pop()
}

// basic stack
type stack struct {
	items []int
}

func newStack(n int) *stack {
	return &stack{
		items: make([]int, 0, n),
	}
}

func (s *stack) Push(t int) {
	s.items = append(s.items, t)
}

func (s *stack) Len() int {
	return len(s.items)
}

func (s *stack) Peek() int {
	return s.items[s.Len()-1]
}

func (s *stack) Pop() int {
	t := s.items[s.Len()-1]

	s.items = s.items[:s.Len()-1]

	return t
}
