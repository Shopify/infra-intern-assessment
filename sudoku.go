package main

// isValid checks whether it will be legal to assign num to the given row, col
func isValid(board [][]int, row, col, num int) bool {
	// Check row and column
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Check box
	startRow := row - row%3
	startCol := col - col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}

// SolveSudoku is the function that attempts to solve a given Sudoku puzzle
// using a backtracking algorithm. It will return the solved puzzle if it's solvable.
func SolveSudoku(board [][]int) [][]int {
	if solve(board, 0, 0) {
		return board
	}
	// Return nil if no solution exists
	return nil
}

// solve is the recursive function that applies the backtracking algorithm
// to solve the Sudoku puzzle.
func solve(board [][]int, row, col int) bool {
	// If we have reached the end of the board, return true
	if row == 9 {
		return true
	}

	// Move to the next row if we have reached the end of a column
	if col == 9 {
		return solve(board, row+1, 0)
	}

	// If the current cell is not empty, continue to the next cell
	if board[row][col] != 0 {
		return solve(board, row, col+1)
	}

	// Try placing numbers 1-9 in the current empty cell
	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solve(board, row, col+1) {
				return true
			}
			// Reset the cell if placing num doesn't lead to a solution
			board[row][col] = 0
		}
	}

	// If no valid number can be placed in this cell, backtrack
	return false
}