package main


// SolveSudoku SolveSudoku Problem
func SolveSudoku(board [][]int) [][]int {
	solve(board)
	return board
}


// Solve uses a recursive approach to attempt solving the Sudoku puzzle.
// The parameter 'board' represents the Sudoku board, which is a 9x9 2D array.
// The function tries to fill each empty cell to satisfy Sudoku rules.
// It returns true if a solution is found, and false otherwise.
func solve(board [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {// If the current cell is empty
			if board[i][j] == 0 {// Try filling numbers 1 to 9
				for num := 1; num <= 9; num++ {// Check if the current number is valid at the current position
					if isValid(board, i, j, num) {// If valid, place the number in the current position
						board[i][j] = num// Recursively call solve to continue filling the next empty cell
						if solve(board) {// If a solution is found, return true
							return true
						}// If recursion fails to find a solution, reset the current position to 0 and backtrack
						board[i][j] = 0
					}
				}// If all numbers have been tried and none is valid, return false indicating no solution
				return false
			}
		}
	}// All cells are filled, indicating a solution is found, return true
	return true
}

// isValid checks if the given position is valid, if the current number can be placed without violating Sudoku rules.
// The parameters 'board' represents the Sudoku board, 'row' and 'col' indicate the position to check, and 'num' is the number to be placed.
// It returns true if the current position allows placing the number, and false if there is a violation.
func isValid(board [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {// Check if the current number is repeated in the current row or column
		if board[row][i] == num || board[i][col] == num {
			return false
		}
		// check neighborhood (3*3)
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

