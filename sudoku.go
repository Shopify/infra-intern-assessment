package main

type SudokuBoard [][]int

// Checks whether a given value placement is valid (that is, the same number
// exists in neither the block, column, or row).
// row and int numbers start from 1 and in the top left corner, moving down and right.
func isValidCell(row int, col int, board *board) {
	// Check row and col
	for (i := 0; i < 9; i++) {
		if (*board[row][col] == board[row][i] || *board[row][col] == board[i][col]) {
			return false
		}
	}
	// Check block
	// Determine the block indices using some simple math
	block_row := row / 3
	block_col := col / 3
	for (i := 0; i < 3; i++) {
		for (j := 0; j < 3; j++) {
			if (*board[i + block_row * 3][j+ block_col * 3] == *board[row][col]) {
				return false
			}
		}
	}
	return true
}

// Use a simple backtracking algorithm. Assume there is a solution.
func SolveSudoku(board SudokuBoard) {
	// Maintain a stack of cells that represent our "path" to filling the board.

	// The latest cell in the stack is the cell we filled the most recently.

	for (i := 0; i < 9; i++) { // go row-by-row
		// Search for an empty cell
		for (j := 0; j < 9; j++) {
			if (board[i][j] == 0) { // empty cell
				valid := false
				for (val := 1; val < 9; val++) {
					board[i][j] = val
					valid := isValidCell(i, j, board)
					if (valid) { // valid value found
						break
					}
				}
				// If none of the values are valid, backtrack using our stack of cells
				while ()
			}
		}
	}
	return board
}