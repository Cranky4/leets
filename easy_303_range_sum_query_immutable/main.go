package main

import "fmt"

/*
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
*/
func main() {
	/*
	  -2  0  3 -5  2 -1
	   0 -2 -2  1 -4  2  1
	*/
	n := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(n.SumRange(0, 2))
	fmt.Println(n.SumRange(2, 5))
	fmt.Println(n.SumRange(0, 5))
}

type NumArray struct {
	prefs []int
}

func Constructor(nums []int) NumArray {
	prefs := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefs[i+1] = prefs[i] + nums[i]
	}

	return NumArray{
		prefs: prefs,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefs[right+1] - this.prefs[left]
}
