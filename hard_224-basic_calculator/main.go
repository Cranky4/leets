package main

import "fmt"

func main() {
	in := " 30"
	rpn := sort_station(in)
	fmt.Printf("ОПН (RPN) - %v -> %v\n", in, rpn)

	fmt.Printf("%v", rpn_calc(rpn))
}
