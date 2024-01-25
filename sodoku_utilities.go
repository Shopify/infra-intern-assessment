package main

func backtrackSolution(sodoku *Sodoku, rowPosition, colPosition uint8) bool {
	// reached the last element
	if rowPosition == sodoku.NumberOfRows-1 && colPosition >= sodoku.NumberOfCols {
		return true
	}

	if colPosition >= sodoku.NumberOfCols {
		rowPosition++
		colPosition = 0
	}

	if sodoku.grid[rowPosition][colPosition] > EMPTY_SPACE {
		return backtrackSolution(sodoku, rowPosition, colPosition+1)
	} else { // spot is zero
		for candidate := 1; candidate <= MAX_POSSIBLE_NUM; candidate++ {
			if !sodoku.ConflictCheck(candidate, rowPosition, colPosition) {
				sodoku.solvedGrid[rowPosition][colPosition] = candidate

				if backtrackSolution(sodoku, rowPosition, colPosition+1) {
					return true
				}
			}

			sodoku.solvedGrid[rowPosition][colPosition] = EMPTY_SPACE
		}
	}

	return false
}

func (sodoku *Sodoku) Solve() {
	if backtrackSolution(sodoku, 0, 0) {
		sodoku.solved = true
	}
}

func (sodoku *Sodoku) rowConflict(candidate int, numRow int) bool {
	for row := 0; row < GRID_ROW_SIZE; row++ {
		if sodoku.solvedGrid[numRow][row] == candidate {
			return true
		}
	}

	return false
}

func (sodoku *Sodoku) colConflict(candidate int, numCol int) bool {
	for col := 0; col < GRID_COL_SIZE; col++ {
		if sodoku.solvedGrid[col][numCol] == candidate {
			return true
		}
	}

	return false
}

func (sodoku *Sodoku) squareConflict(candidate int, numRow, numCol int) bool {
	squareRow := numRow - numRow%SQUARE_SIZE
	squareCol := numCol - numCol%SQUARE_SIZE

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sodoku.solvedGrid[i+squareRow][j+squareCol] == candidate {
				return true
			}
		}
	}

	return false
}

func (sodoku *Sodoku) ConflictCheck(candidate int, numRow, numCol uint8) bool {
	return sodoku.colConflict(candidate, int(numCol)) ||
		sodoku.rowConflict(candidate, int(numRow)) ||
		sodoku.squareConflict(candidate, int(numRow), int(numCol))
}
