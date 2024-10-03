package backtracking

func backtrackI(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		if !(*selected)[i] {
			(*selected)[i] = true
			*state = append(*state, choice)
			backtrackI(state, choices, selected, res)
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

func permutationsI(nums []int) [][]int {
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	res := make([][]int, 0)
	backtrackI(&state, &nums, &selected, &res)
	return res
}

func backtrackII(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}
	duplicated := make(map[int]struct{}, 0)
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		if _, ok := duplicated[choice]; !ok && !(*selected)[i] {
			(*selected)[i] = true
			duplicated[choice] = struct{}{}
			*state = append(*state, choice)
			backtrackI(state, choices, selected, res)
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}
