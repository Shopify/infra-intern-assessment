// Shopify Infrastructure Internship Assessment
// Makhdoomzada Ali Syed
// 2024-26-01
// Sudoku Solver

package main

// Main function to call to solve a given sudoku board
func SolveSudoku(board [][]int) [][]int {
	solve(board)
	return board
}

// Recursive solve function that uses backtracking to solve the board
func solve(board [][]int) bool {
	// Iterate over the board which is just a 2D array hence use row and col
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// If the cell is empty attempt to fill it with a digit
			if board[row][col] == 0 {
				for digit := 1; digit <= 9; digit++ {
					// If the digit is safe to place in a given cell place it
					// "safe" means the digit doesnt appear in the same row, col or 3x3 sub-grid
					if safe(board, row, col, digit) {
						board[row][col] = digit

						// Recursively call solve to solve the rest of the puzzle
						if solve(board) {
							// If the entire board is solved return true
							return true
						}

						// If the board is not solved we backtrack by resetting the digit
						// and trying the next digit
						board[row][col] = 0
					}
				}
				// If we have tried all digits and none are safe to place return false
				return false
			}
		}
	}
	// If we have iterated over the entire board and all cells are filled return true
	// this is the base case and indicates the puzzle is solved
	return true
}

// Helper function to check if a digit is safe to place in a given cell
// Conditions for a digit to be "safe":
// 1. Digit does not appear in the same row
// 2. Digit does not appear in the same col
// 3. Digit does not appear in the same 3x3 sub-grid
func safe(board [][]int, row int, col int, digit int) bool {

	// Doing the checks in order helps optimize as
	// if the digit is not safe to place in a row or col
	// we can return false early instead of checking both
	// conditions simultaneously in a return statement

	// Check rule 1 & 2 (saves time by checking both at once)
	if !safeRowCol(board, row, col, digit) {
		return false
	}

	// Check rule 3
	if !safeThreeByThree(board, row, col, digit) {
		return false
	}

	// If the digit is safe to place return true
	return true
}

// Helper function to check rule 1 & 2
func safeRowCol(board [][]int, row int, col int, digit int) bool {
	// Iterate over the row and col
	for i := 0; i < 9; i++ {
		// If the digit appears in the row or col return false
		if board[row][i] == digit || board[i][col] == digit {
			return false
		}
	}
	// If the digit does not appear in the row or col return true
	return true
}

// helper function to check rule 3
func safeThreeByThree(board [][]int, row int, col int, digit int) bool {
	// Calculate the top left corner of the 3x3 grid using the fact
	// that row%3 and col%3 calc the a cells current pos in 3x3 grid
	for i := row - row%3; i < row-row%3+3; i++ {
		// Iterating over the 3x3 grid using the top left corner as start
		for j := col - col%3; j < col-col%3+3; j++ {
			// If the digit appears in the 3x3 grid return false
			if board[i][j] == digit {
				return false
			}
		}
	}
	// If the digit does not appear in the 3x3 grid return true
	return true
}
