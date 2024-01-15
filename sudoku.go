package main

// Check whether the given item indexed by itemRow and itemCol
// is valid against the sudoku. It verifies the row, col and block
// where the item is located.
//
// Assumes sudoku is a 9x9 grid and itemRow and itemCol are valid indexes.
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

// Find the first empty cell in the sudoku
//
// Assumes sudoku is a 9x9 grid and itemRow and itemCol are valid indexes.
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

// Print the 2D array
func printSudoku(sudoku [][]int) {
	for _, row := range sudoku {
		for _, item := range row {
			print(item, " ")
		}
		println()
	}
}

func SolveSudoku(sudoku [][]int) [][]int {
	// Find the first empty cell
	x, y := findEmpty(sudoku)

	if x == -1 || y == -1 {
		// Print the complete sudoku
		printSudoku(sudoku)

		return sudoku
	}

	// Iterate through possible answers
	for i := 1; i < 10; i++ {
		sudoku[x][y] = i

		// Check if the in place value is valid
		if isItemValid(sudoku, x, y) {
			// Check if there is a solved solution
			if SolveSudoku(sudoku) != nil {
				return sudoku
			}
		}

		// Backtrack
		sudoku[x][y] = 0
	}

	// No answer is found
	return nil
}
