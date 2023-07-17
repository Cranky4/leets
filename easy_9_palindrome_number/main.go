package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x < 10 {
        return true
    }

    arr := make([]int, 0)

    for x > 9 {
        arr = append(arr, x % 10)
        x /= 10
    }

    arr = append(arr, x)

   
   i :=0 
   j := len(arr)-1

   for i < j {
       if arr[i] != arr[j] {
           return false
       }
       i++
       j--
   }

    return true
}