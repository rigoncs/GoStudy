package src

import "math"

func knapsackDFS(wgt, val []int, i, c int) int {
	if i == 0 || c == 0 {
		return 0
	}
	if wgt[i-1] > c {
		return knapsackDFS(wgt, val, i-1, c)
	}
	no := knapsackDFS(wgt, val, i-1, c)
	yes := knapsackDFS(wgt, val, i-1, c-wgt[i-1]) + val[i-1]
	return int(math.Max(float64(no), float64(yes)))
}

func knapsackDFSMem(wgt, val []int, mem [][]int, i, c int) int {
	if i == 0 || c == 0 {
		return 0
	}
	if mem[i][c] != -1 {
		return mem[i][c]
	}
	if wgt[i-1] > c {
		return knapsackDFSMem(wgt, val, mem, i-1, c)
	}
	no := knapsackDFSMem(wgt, val, mem, i-1, c)
	yes := knapsackDFSMem(wgt, val, mem, i-1, c-wgt[i-1]) + val[i-1]
	mem[i][c] = int(math.Max(float64(no), float64(yes)))
	return mem[i][c]
}

func knapsackDP(weg, val []int, cap int) int {
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
				dp[i][c] = int(math.Max(float64(dp[i-1][c]), float64(dp[i-1][c-weg[i-1]]+val[i-1])))
			}
		}
	}
	return dp[n][cap]
}

func knapsackDPComp(weg, val []int, cap int) int {
	n := len(weg)
	dp := make([]int, cap+1)
	for i := 1; i <= n; i++ {
		for c := cap; c >= 1; c-- {
			if weg[i-1] <= c {
				dp[c] = int(math.Max(float64(dp[c]), float64(dp[c-weg[i-1]]+val[i-1])))
			}
		}
	}
	return dp[cap]
}
