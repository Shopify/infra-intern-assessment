package main

// Convert an index to a sub-grid number
func iToSg(i int) int {
	return (i/27)*3 + (i%9)/3
	// Note: /27*3 is not the same as /9 (because integer division).
}

// DFS to find the solution.
func solveSudokuHelper(i int, board [][]int, colUsed [9][9]bool, rowUsed [9][9]bool, sgUsed [9][9]bool) bool {
	// If the current index is off the end of the board... we're done!
	if i == 81 {
		return true
	}

	// If the current index is already set... skip to the next index.
	if board[i/9][i%9] != 0 {
		return solveSudokuHelper(i+1, board, colUsed, rowUsed, sgUsed)
	}

	// Otherwise, go through all feasible numbers.
	for target := range board {
		if colUsed[i%9][target] || rowUsed[i/9][target] || sgUsed[iToSg(i)][target] {
			continue
		}

		colUsed[i%9][target] = true
		rowUsed[i/9][target] = true
		sgUsed[iToSg(i)][target] = true

		if solveSudokuHelper(i+1, board, colUsed, rowUsed, sgUsed) {
			board[i/9][i%9] = target + 1
			return true
		}

		colUsed[i%9][target] = false
		rowUsed[i/9][target] = false
		sgUsed[iToSg(i)][target] = false
	}

	return false
}

func SolveSudoku(board [][]int) [][]int {
	colUsed := [9][9]bool{} // colUsed[i][j]==true implies j+1 has not yet been set in col i
	rowUsed := [9][9]bool{} // rowUsed[i][j]==true implies j+1 has not yet been set in row i
	sgUsed := [9][9]bool{}  // sgUsed[i][j]==true implies j+1 has not yet been set in sub-grid i

	for y := range board {
		for x := range board {
			target := board[y][x]
			if target == 0 {
				continue
			}

			rowUsed[y][target-1] = true
			colUsed[x][target-1] = true
			sgUsed[iToSg(9*y+x)][target-1] = true
		}
	}

	solution := make([][]int, len(board))
	for y := range solution {
		solution[y] = make([]int, len(board[y]))
		copy(solution[y], board[y])
	}

	solveSudokuHelper(0, solution, colUsed, rowUsed, sgUsed)
	return solution
}
