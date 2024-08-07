package main

const N = 9 // Define the size of the Sudoku grid

// isSafe checks if it's valid to put a num in the given row and column.
func isSafe(board [][]int, row, col, num int) bool {
	// Check the row
	for x := 0; x < N; x++ {
		if board[row][x] == num {
			return false
		}
	}

	// Check the column
	for x := 0; x < N; x++ {
		if board[x][col] == num {
			return false
		}
	}

	// Check the 3x3 box containing the cell
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

// SolveSudoku solves the Sudoku puzzle using backtracking.
func SolveSudoku(board [][]int) [][]int {
	solve(board)
	return board
}

// solve is the recursive function to solve the puzzle using backtracking.
func solve(board [][]int) bool {
	row, col := -1, -1
	isEmpty := true

	// Find an empty cell
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				isEmpty = false
				break
			}
		}
		if !isEmpty {
			break
		}
	}

	// If there are no empty cells, we're done
	if isEmpty {
		return true
	}

	// Try all numbers from 1 to 9 in the empty cell
	for num := 1; num <= N; num++ {
		if isSafe(board, row, col, num) {
			board[row][col] = num
			if solve(board) {
				return true
			}
			board[row][col] = 0 // Backtrack
		}
	}

	return false
}
