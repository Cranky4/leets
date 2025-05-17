package main

import "fmt"

func main() {
	in := []string{"a", "b", "c", "d"}
	groups := []int{1, 0, 1, 1}

	result := getLongestSubsequence(in, groups)

	fmt.Println(result)
}

func getLongestSubsequence(words []string, groups []int) []string {
	if len(words) == 0 {
		return nil
	}

	ln := len(words)
	dp := make([][]string, ln)
	dp[0] = append(dp[0], words[0])

	for i := 1; i < ln; i++ {
		if groups[i] != groups[i-1] {
			dp[i] = append(dp[i-1], words[i])
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[ln-1]
}
