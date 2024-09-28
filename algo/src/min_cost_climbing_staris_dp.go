package src

func minCostClimbingStairsDP(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 2 {
		return n
	}
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	dp := make([]int, n+1)
	dp[1] = cost[1]
	dp[2] = cost[2]
	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return dp[n]
}

func minCostClimbingStairsDPComp(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 2 {
		return n
	}
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	a, b := cost[1], cost[2]
	for i := 3; i <= n; i++ {
		tmp := b
		b = min(a, tmp) + cost[i]
		a = tmp
	}
	return b
}

func climbingStairsConstraintDP(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([][3]int, n+1)
	dp[1][1] = 1
	dp[1][2] = 0
	dp[2][1] = 0
	dp[2][2] = 1
	for i := 3; i <= n; i++ {
		dp[i][1] = dp[i-1][2]
		dp[i][2] = dp[i-2][1] + dp[i-2][2]
	}
	return dp[n][1] + dp[n][2]
}
