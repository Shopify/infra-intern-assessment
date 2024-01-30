package main

func isMappingValid(g grid, c cell, newValue int) bool {
	// Check row and column
	for i := 0; i < 9; i++ {
		if g[c.row()][i] == newValue || g[i][c.column()] == newValue {
			return false
		}
	}

	// Check box
	boxRow := c.row() / 3 * 3
    boxColumn := c.column() / 3 * 3
    for i := boxRow; i < boxRow + 3; i++ {
        for j := boxColumn; j < boxColumn + 3; j++ {
			if i == c.row() || j == c.column() { // row and column already checked, so skip
				continue
			}

            if g[i][j] == newValue {
                return false
            }
        }
    }

	return true
}

func SolveSudokuRecursive(inputGrid grid, c cell) (solvedGrid [][]int) {
	if c == END_OF_GRID {
		return Copy(inputGrid)
	}

	if inputGrid.At(c) != 0 {
		return SolveSudokuRecursive(inputGrid, c.next())
	}

	for i := 1; i <= 9; i++ {
		if !isMappingValid(inputGrid, c, i) {
			continue
		}

		inputGrid.Set(c, i)
		solveChild := SolveSudokuRecursive(inputGrid, c.next())
		inputGrid.Reset(c)

		if solveChild != nil {
			return solveChild
		}
	}

	return nil
}

func SolveSudoku(inputGrid grid) (solvedGrid [][]int) {
	return SolveSudokuRecursive(inputGrid, 0)
}

func main() {
}