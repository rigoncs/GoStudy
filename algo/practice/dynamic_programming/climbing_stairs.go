package dynamicprogramming

func backtrack(choices []int, state, n int, res []int) {
	if state == n {
		res[0]++
	}
	for _, choice := range choices {
		if state+choice > n {
			continue
		}
		backtrack(choices, state+choice, n, res)
	}
}

func climbingStairsBacktrack(n int) int {
	state := 0
	res := make([]int, 1)
	choices := []int{1, 2}
	res[0] = 0
	backtrack(choices, state, n, res)
	return res[0]
}

func dfs(i int) int {
	if i == 1 || i == 2 {
		return i
	}
	count := dfs(i-1) + dfs(i-2)
	return count
}

func climbingStairsDFS(n int) int {
	return dfs(n)
}

func dfsMem(i int, mem []int) int {
	if i == 1 || i == 2 {
		return i
	}
	if mem[i] != -1 {
		return mem[i]
	}
	count := dfsMem(i-1, mem) + dfsMem(i-2, mem)
	mem[i] = count
	return count
}

func climbingStairsDFSMem(n int) int {
	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = -1
	}
	return dfsMem(n, mem)
}

func climbingStairsDP(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbingStairsDPComp(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
