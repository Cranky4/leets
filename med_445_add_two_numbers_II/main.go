package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	l3 := addTwoNumbers(l1, l2)

	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    s1 := NewStack(1)
    node := l1

    for node != nil {
        s1.Add(node)
        node = node.Next
    }

    s2 := NewStack(1)
    node = l2

    for node != nil {
        s2.Add(node)
        node = node.Next
    }

    //
    s3 := NewStack(3)
    caret := 0

    for s1.Len() > 0 || s2.Len() > 0 {
        n1 := 0
        if s1.Len() > 0 {
            n1 = s1.Pop().Val
        }

        n2 := 0
        if s2.Len() > 0 {
            n2 = s2.Pop().Val
        }

        sum := n1 + n2

        if caret == 1 {
            sum++
            caret = 0
        }


        if sum > 9 {
            sum = sum % 10
            caret = 1
        }

        s3.Add(&ListNode{Val: sum})
    }

    if caret > 0 {
        s3.Add(&ListNode{Val: 1})
    }

    cur := s3.Pop()
    first := cur
    for s3.Len() > 0 {
        next := s3.Pop()
        cur.Next = next
        cur = next
    }

    return first
}

type Stack struct {
    items []*ListNode
}

func NewStack(l int) *Stack {
    return &Stack{
        items: make([]*ListNode, 0, l),
    }
}

func (s *Stack) Add(i *ListNode) {
    s.items = append(s.items, i)
}

func (s *Stack) Len() int {
    return len(s.items)
}

func (s *Stack) Pop() *ListNode {
    if s.Len() < 1 {
        return nil
    }

    i := s.items[s.Len() - 1]

    s.items = s.items[0:s.Len() - 1]

    return i
}