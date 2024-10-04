package dynamicprogramming

func editDistanceDFS(s, t string, i, j int) int {
	if i == 0 && j == 0 {
		return 0
	}
	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}
	if s[i-1] == t[j-1] {
		return editDistanceDFS(s, t, i-1, j-1)
	}
	insert := editDistanceDFS(s, t, i, j-1)
	delete := editDistanceDFS(s, t, i-1, j)
	replace := editDistanceDFS(s, t, i-1, j-1)
	return MinInt(insert, MinInt(delete, replace)) + 1
}

func editDistanceDFSMem(s, t string, mem [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return 0
	}
	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}
	if mem[i][j] != -1 {
		return mem[i][j]
	}
	if s[i-1] == t[j-1] {
		return editDistanceDFS(s, t, i-1, j-1)
	}
	insert := editDistanceDFS(s, t, i, j-1)
	delete := editDistanceDFS(s, t, i-1, j)
	replace := editDistanceDFS(s, t, i-1, j-1)
	mem[i][j] = MinInt(insert, MinInt(delete, replace)) + 1
	return mem[i][j]
}

func editDistanceDP(s, t string) int {
	n := len(s)
	m := len(t)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= m; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = MinInt(dp[i-1][j], MinInt(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}
	return dp[n][m]
}

func editDistanceDPComp(s, t string) int {
	n := len(s)
	m := len(t)
	dp := make([]int, m+1)
	for j := 1; j <= m; j++ {
		dp[j] = j
	}
	for i := 1; i <= n; i++ {
		leftUp := dp[0]
		dp[0] = i
		for j := 1; j <= m; j++ {
			temp := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = leftUp
			} else {
				dp[j] = MinInt(MinInt(dp[j-1], dp[j]), leftUp) + 1
			}
			leftUp = temp
		}
	}
	return dp[m]
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
