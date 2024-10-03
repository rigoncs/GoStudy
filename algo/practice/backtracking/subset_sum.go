package backtracking

import "sort"

func backtrackSubsetSumINaive(total, target int, state, choices *[]int, res *[][]int) {
	if total == target {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
		return
	}
	for i := 0; i < len(*choices); i++ {
		if total+(*choices)[i] > target {
			continue
		}
		*state = append(*state, (*choices)[i])
		backtrackSubsetSumINaive(total, target, state, choices, res)
		*state = (*state)[:len(*state)-1]
	}
}

func subsetSumINaive(nums []int, target int) [][]int {
	state := make([]int, 0)
	res := make([][]int, 0)
	total := 0
	backtrackSubsetSumINaive(total, target, &state, &nums, &res)
	return res
}

func backtrackSubsetSumI(start, target int, state, choice *[]int, res *[][]int) {
	if target == 0 {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
		return
	}
	for i := start; i < len(*choice); i++ {
		if target-(*choice)[i] < 0 {
			break
		}
		*state = append(*state, (*choice)[i])
		backtrackSubsetSumI(i, target-(*choice)[i], state, choice, res)
		*state = (*state)[:len(*state)-1]
	}
}

func subsetSumI(nums []int, target int) [][]int {
	state := make([]int, 0)
	res := make([][]int, 0)
	sort.Ints(nums)
	start := 0
	backtrackSubsetSumI(start, target, &state, &nums, &res)
	return res
}

func backtrackSubsetSumII(start, target int, state, choice *[]int, res *[][]int) {
	if target == 0 {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
		return
	}
	for i := start; i < len(*choice); i++ {
		if target-(*choice)[i] < 0 {
			break
		}
		if i > start && (*choice)[i] == (*choice)[i-1] {
			continue
		}
		*state = append(*state, (*choice)[i])
		backtrackSubsetSumI(i, target-(*choice)[i], state, choice, res)
		*state = (*state)[:len(*state)-1]
	}
}
