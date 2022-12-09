package main

import "fmt"

func main() {
	fmt.Printf("%v\n", fib(10))
}

func fib(n int) int {
	prev := 0
	prevprev := 0
	res := 0

	for i := 1; i <= n; i++ {
		if prev == 0 {
			res = i
		} else {
			res = prev + prevprev
			prevprev = prev
		}

		prev = res
	}

	return res
}
