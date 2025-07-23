package main

// Albert Nguyen-Tran
// January 27, 2024
// SolveSudoku takes a partially filled Sudoku board represented as a 9x9 grid of integers
// and returns the solved Sudoku board.
func SolveSudoku(board [][]int) [][]int {
	// Start the backtracking solution from the top-left cell (0, 0).
	backtrack(board, 0, 0)

	// Return the solved Sudoku board after the backtracking algorithm has completed.
	return board
}

// backtrack is a recursive backtracking function that attempts to fill the Sudoku board.
// It explores different possibilities by trying numbers from 1 to 9 in each empty cell.
// Therefore in the worst case, the TC = O(9^(m*n))
// And the SC = O(m*n) because the actions are done in place
// If a valid solution is found, it returns true; otherwise, it backtracks.
func backtrack(board [][]int, i, j int) bool {
	// If the current row index (i) is 9, the entire board is filled, and a solution is found.
	if i == 9 {
		return true
	}
	// If the current column index (j) is 9, move to the next row.
	if j == 9 {
		return backtrack(board, i+1, 0)
	}
	// If the current cell is not empty (contains a non-zero value), move to the next cell.
	if board[i][j] != 0 {
		return backtrack(board, i, j+1)
	}
	// Try filling the current cell with numbers from 1 to 9.
	for temp := 1; temp <= 9; temp++ {
		// Check if the number temp is valid in the current cell.
		if !isValid(board, i, j, temp) {
			continue
		}
		// Set the current cell to temp.
		board[i][j] = temp
		// Recursively move to the next cell and continue the exploration.
		if backtrack(board, i, j+1) {
			return true
		}
		// If the recursive call does not lead to a solution, backtrack by resetting the current cell to 0.
		board[i][j] = 0
	}
	// No valid number was found for the current cell, backtrack further.
	return false
}

// isValid checks whether placing the number temp in the cell at position (i, j) is valid.
// It checks the row, column, and the 3x3 sub-box for any conflicts.
func isValid(board [][]int, i, j, temp int) bool {
	// Check the column for conflicts with the same number temp.
	for x := 0; x < 9; x++ {
		if board[x][j] == temp {
			return false
		}
	}
	// Check the row for conflicts with the same number temp.
	for y := 0; y < 9; y++ {
		if board[i][y] == temp {
			return false
		}
	}
	// Check the 3x3 sub-box for conflicts with the same number temp.
	row, col := i-i%3, j-j%3
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if board[row+x][col+y] == temp {
				return false
			}
		}
	}
	// No conflicts found, the number temp can be placed in the current cell.
	return true
}
