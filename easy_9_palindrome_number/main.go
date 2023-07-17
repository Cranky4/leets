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

    n := x
    rev := 0

    for n > 0 {
        rev = rev * 10 + (n % 10)
        n /= 10
    }

    return x == rev
}