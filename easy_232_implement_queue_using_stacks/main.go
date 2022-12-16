package main

import "fmt"

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Printf("peek %d\n", q.Peek())
	fmt.Printf("pop %d\n", q.Pop())
	fmt.Printf("empty %v\n", q.Empty())
	fmt.Printf("pop %d\n", q.Pop())
	fmt.Printf("empty %v\n", q.Empty())
}

type MyQueue struct {
	stack1 *stack
	stack2 *stack
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: NewStack(10),
		stack2: NewStack(10),
	}
}

func (this *MyQueue) Push(x int) {
	// put all from first stack to second
	for {
		if this.stack1.len() == 0 {
			this.stack1.push(x)
			break
		}

		i, _ := this.stack1.pop()
		this.stack2.push(i)
	}

	// put back
	for {
		if this.stack2.len() == 0 {
			break
		}

		i, _ := this.stack2.pop()
		this.stack1.push(i)
	}
}

func (this *MyQueue) Pop() int {
	i, _ := this.stack1.pop()
	return i
}

func (this *MyQueue) Peek() int {
	i, _ := this.stack1.peek()
	return i
}

func (this *MyQueue) Empty() bool {
	return this.stack1.len() == 0
}

// stack
type stack struct {
	items []int
}

func NewStack(n int) *stack {
	return &stack{
		items: make([]int, 0, n),
	}
}

func (s *stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *stack) peek() (int, bool) {
	if s.len() == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *stack) pop() (int, bool) {
	if s.len() == 0 {
		return 0, false
	}

	i := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return i, true
}

func (s *stack) len() int {
	return len(s.items)
}
