package src

import "math"

func unboundedKnapsackDP(weg, val []int, cap int) int {
	n := len(weg)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, cap+1)
	}
	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			if weg[i-1] > c {
				dp[i][c] = dp[i-1][c]
			} else {
				dp[i][c] = int(math.Max(float64(dp[i-1][c]), float64(dp[i][c-weg[i-1]]+val[i-1])))
			}
		}
	}
	return dp[n][cap]
}

func unboundedKnapsackDPComp(weg, val []int, cap int) int {
	n := len(weg)
	dp := make([]int, cap+1)
	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			if weg[i-1] > c {
				dp[c] = dp[c]
			} else {
				dp[c] = int(math.Max(float64(dp[c]), float64(dp[c-weg[i-1]]+val[i-1])))
			}
		}
	}
	return dp[cap]
}
