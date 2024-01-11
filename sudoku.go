package main

// Grid dimensions (9x9)
const G = 9

// SolveSudoku recieves a 9x9 sudoku board and returns a solved version of the board.
func SolveSudoku(board [][]int) [][]int {

	// Find an empty space on `board`
	for row := 0; row < G; row++ {
		for col := 0; col < G; col++ {

			// Select number from 1-9 and validate its location in 3x3 sub-grid
			if board[row][col] == 0 {
				ValidateSubGrid(board, row, col, board[row][col])
			}

			// Some back-tracking stuff probably

			// Guard clause for if the board cannot be solbved
		}
	}

	return board
}

// ValidateSubGrid (helper function) takes a 9x9 sudoku board, a row, a column, and a number
// and returns true if the number can be placed in the given 3x3 sub-grid.
func ValidateSubGrid(board [][]int, row int, col int, num int) bool {
	// Check row and col and verify num placed is valid

	return true
}
