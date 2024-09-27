package src

func backTrackPermutationI(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		if !(*selected)[i] {
			(*selected)[i] = true
			*state = append(*state, choice)
			backTrackPermutationI(state, choices, selected, res)
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

func backTrackPermutationII(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}
	duplicated := make(map[int]struct{}, 0)
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		if _, ok := duplicated[choice]; !ok && !(*selected)[i] {
			duplicated[choice] = struct{}{}
			(*selected)[i] = true
			*state = append(*state, choice)
			backTrackPermutationI(state, choices, selected, res)
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

func permutations(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	backTrackPermutationI(&state, &nums, &selected, &res)
	return res
}
