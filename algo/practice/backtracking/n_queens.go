package backtracking

func backtrack(row, n int, state *[][]string, res *[][][]string, cols, diags1, diags2 *[]bool) {
	if row == n {
		newState := make([][]string, len(*state))
		for i := range newState {
			newState[i] = make([]string, len((*state)[0]))
			copy(newState[i], (*state)[i])
		}
		*res = append(*res, newState)
		return
	}
	for col := 0; col < len(*cols); col++ {
		diag1 := row - col + n - 1
		diag2 := row + col
		if !(*cols)[col] && !(*diags1)[diag1] && !(*diags2)[diag2] {
			(*state)[row][col] = "Q"
			(*cols)[col], (*diags1)[diag1], (*diags2)[diag2] = true, true, true
			backtrack(row+1, n, state, res, cols, diags1, diags2)
			(*cols)[col], (*diags1)[diag1], (*diags2)[diag2] = false, false, false
			(*state)[row][col] = "#"
		}
	}
}

func nQueens(n int) [][][]string {
	state := make([][]string, n)
	for i := 0; i < n; i++ {
		row := make([]string, n)
		for j := 0; j < n; j++ {
			row[j] = "#"
		}
		state[i] = row
	}
	res := make([][][]string, 0)
	cols := make([]bool, n)
	diags1 := make([]bool, 2*n-1)
	diags2 := make([]bool, 2*n-1)
	backtrack(0, n, &state, &res, &cols, &diags1, &diags2)
	return res
}
