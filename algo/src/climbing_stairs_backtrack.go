package src

func backtrack(choices []int, state, n int, res []int) {
	if state == n {
		res[0] = res[0] + 1
	}
	for _, choice := range choices {
		if state+choice > n {
			continue
		}
		backtrack(choices, state+choice, n, res)
	}
}

func climbingStairsBacktrack(n int) int {
	choices := []int{1, 2}
	state := 0
	res := make([]int, 1)
	res[0] = 0
	backtrack(choices, state, n, res)
	return res[0]
}

func dfsBruteForce(i int) int {
	if i == 1 || i == 2 {
		return i
	}
	count := dfsBruteForce(i-1) + dfsBruteForce(i-2)
	return count
}

func climbingStairsDFS(n int) int {
	return dfsBruteForce(n)
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
