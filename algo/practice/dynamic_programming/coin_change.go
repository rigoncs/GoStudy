package dynamicprogramming

import "math"

func coinChangeDP(coins []int, amt int) int {
	n := len(coins)
	max := amt + 1
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amt+1)
	}
	for j := 1; j <= amt; j++ {
		dp[0][j] = max
	}
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i] > a {
				dp[i][a] = dp[i-1][a]
			} else {
				dp[i][a] = int(math.Min(float64(dp[i-1][a]), float64(dp[i][a-coins[i-1]]+1)))
			}
		}
	}
	if dp[n][amt] != max {
		return dp[n][amt]
	}
	return -1
}

func coinChangeDPComp(coins []int, amt int) int {
	n := len(coins)
	max := amt + 1
	dp := make([]int, amt+1)
	for i := 1; i <= amt; i++ {
		dp[i] = max
	}
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				dp[a] = dp[a]
			} else {
				dp[a] = int(math.Min(float64(dp[a]), float64(dp[a-coins[i-1]]+1)))
			}
		}
	}
	if dp[amt] != max {
		return dp[amt]
	}
	return -1
}

func coinChangeIIDP(coins []int, amt int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amt+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				dp[i][a] = dp[i-1][a]
			} else {
				dp[i][a] = dp[i-1][a] + dp[i][a-coins[i]]
			}
		}
	}
	return dp[n][amt]
}

func coinChangeIIDPComp(coins []int, amt int) int {
	n := len(coins)
	dp := make([]int, amt+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				dp[a] = dp[a]
			} else {
				dp[a] = dp[a] + dp[a-coins[i]]
			}
		}
	}
	return dp[amt]
}
