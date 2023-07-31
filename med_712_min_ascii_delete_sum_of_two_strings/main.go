package main

import "fmt"

func main() {
	fmt.Println(minimumDeleteSum("delete", "leet"))
	fmt.Println(minimumDeleteSum("sea", "eat"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i, r := range s1 {
		dp[i+1][0] = int(r) + dp[i][0]
	}

	for i, r := range s2 {
		dp[0][i+1] = int(r) + dp[0][i]
	}

	for i, r1 := range s1 {
		for j, r2 := range s2 {
			if r1 == r2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1]+int(r1), dp[i+1][j]+int(r2))
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
