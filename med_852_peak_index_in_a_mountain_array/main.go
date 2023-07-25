package main

import "fmt"

func main() {
	in := [][]int {
		{0,1,0},
		{0,2,1,0},
		{0,10,5,2},
		{55,59,63,99,97,94,84,81,79,66,40,38,33,23,22,21,17,9,7},
	}

	for _, i := range in {
		fmt.Println(peakIndexInMountainArray(i))
	}
}

func peakIndexInMountainArray(arr []int) int {
    l := 0
    r := len(arr) - 1

    var m int

    for l < r {
        m = (l + r) / 2
        if arr[m] < arr[m + 1] {
            l = m + 1
        } else {
            r = m
        }
    }

    return l
}