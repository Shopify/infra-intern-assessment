package main

// Convert an index to a subgrid number
func iToSg(i int) int {
	if i < 27 {
		return (i % 9) / 3
	} else if i < 54 {
		return 3 + (i%9)/3
	} else {
		return 6 + (i%9)/3
	}
}

func solveSudokuHelper(i int, board [][]int, colUsed [9][9]bool, rowUsed [9][9]bool, sgUsed [9][9]bool) bool {
	// If the last looked at index was the last index, we are done
	if i == 81 {
		return true
	}

	// Find the next free index to look at
	for board[i/9][i%9] != 0 {
		i++

		if i == 81 {
			return true
		}
	}

	// Go through all possible numbers not yet set
	for target, _ := range board {
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
	rowUsed := [9][9]bool{} // rowUsed[i][j]==true implies j+1 has not yet been set in row j
	sgUsed := [9][9]bool{}  // sgUsed[i][j]==true implies j+1 has not yet been set in sub grid j

	for i, _ := range board {
		for j, _ := range board {
			target := board[i][j]
			if target != 0 {
				rowUsed[i][target-1] = true
				colUsed[j][target-1] = true
				sgUsed[iToSg(9*i+j)][target-1] = true
			}
		}
	}

	solveSudokuHelper(0, board, colUsed, rowUsed, sgUsed)
	return board
}
