package main

func isItemValid(sudoku [][]int, itemRow int, itemCol int) bool {
	var row, col, block [10]bool

	// Locate the start of the block
	blockRowStart := itemRow / 3 * 3
	blockColStart := itemCol / 3 * 3

	for i := 0; i < 9; i++ {
		// Check if row has duplicated
		if sudoku[itemRow][i] != 0 && row[sudoku[itemRow][i]] {
			return false
		} else {
			row[sudoku[itemRow][i]] = true
		}

		// Check if column has duplicated
		if sudoku[i][itemCol] != 0 && col[sudoku[i][itemCol]] {
			return false
		} else {
			col[sudoku[i][itemCol]] = true
		}

		// Check if block has duplicated
		x := blockRowStart + (i / 3)
		y := blockColStart + (i % 3)

		if sudoku[x][y] != 0 && block[sudoku[x][y]] {
			return false
		} else {
			block[sudoku[x][y]] = true
		}
	}

	return true
}

func findEmpty(sudoku [][]int) (int, int) {
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku[i]); j++ {
			if sudoku[i][j] == 0 {
				return i, j
			}
		}
	}

	return -1, -1
}

func SolveSudoku(sudoku [][]int) [][]int {
	x, y := findEmpty(sudoku)

	if x == -1 || y == -1 {
		return sudoku
	}

	for i := 1; i < 10; i++ {
		sudoku[x][y] = i

		if isItemValid(sudoku, x, y) {
			if SolveSudoku(sudoku) != nil {
				return sudoku
			}
		}

		sudoku[x][y] = 0
	}

	return nil
}
