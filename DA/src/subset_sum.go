package src

import "sort"

func backTrackSubsetSumI(start, target int, state, choice *[]int, res *[][]int) {
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
		backTrackSubsetSumI(i, target-(*choice)[i], state, choice, res)
		*state = (*state)[:len(*state)-1]
	}
}

func backTrackSubsetSumII(start, target int, state, choice *[]int, res *[][]int) {
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
		backTrackSubsetSumI(i+1, target-(*choice)[i], state, choice, res)
		*state = (*state)[:len(*state)-1]
	}
}

func subsetSum(nums []int, target int) [][]int {
	state := make([]int, 0)
	sort.Ints(nums)
	start := 0
	res := make([][]int, 0)
	backTrackSubsetSumI(start, target, &state, &nums, &res)
	return res
}
