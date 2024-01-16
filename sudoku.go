package main

// canPlaceNumber checks to see if we can place a number in a given cell of the Sudoku board
// by ensuring that the same number is not already placed in the same row, column, or the 3*3
// subgrid the cell lies in.

// Parameters:
// board: 2d array that represents the Sudoku board.
// row: the row index of the cell.
// col: The column index of the cell.
// num: The number we want to place on the cell.

// Return:
// bool: true is the number can be placed and false otherwise
func canPlaceNumber(board [][]int, row, col, num int) bool {
	// Check if the given number is unique in its row
	for i := 0; i <= 8; i++ {
		if board[row][i] == num {
			return false
		}
	}
	// Check if the given number is unique in its column
	for j := 0; j <= 8; j++ {
		if board[j][col] == num {
			return false
		}
	}
	// Check if the given number is unique in its 3*3 subgrid
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

// solveSudokuHelper is a recursive function that tries solving the Sudoku puzzle by implementing
// backtracking.  It tries to place a number from 1 to size in each cell and recursively checks if it leads
// to the solution

// Parameter:
// board: 2d array that represents the Sudoku board.

// Returns:
// bool: Returns true if the Sudoku is solvable and false otherwise
func solveSudokuHelper(board [][]int) bool {
	size := 9 //The variable represents the size of the Sudoku board which for a standard board is 9
	row, col := -1, -1
	isItEmpty := true

	// Finding an empty cell.
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				isItEmpty = false
				break
			}
		}
		if !isItEmpty {
			break
		}
	}
	// When no empty cells are found, the Sudoku will have been solved.
	if isItEmpty {
		return true
	}
	// Try one of the numbers from 1 to size in the found empty cell and
	// recurse.
	for k := 1; k <= size; k++ {
		if canPlaceNumber(board, row, col, k) {
			board[row][col] = k
			if solveSudokuHelper(board) {
				return true
			}
			board[row][col] = 0 //Backtracking happens here when k does not lead to a solution
		}
	}
	return false
}

// SolveSudoku is a public function that attempts to solve a given Sudoku puzzle.
// If there is no solution, it returns a board filled with zeroes.

// Parameter:
// board: 2d array that represents the Sudoku board.

// Returns:
// [][]int: The final solved Sudoku board or a board filled with zeroes if there is no solution
func SolveSudoku(board [][]int) [][]int {
	if solveSudokuHelper(board) {
		return board
	} else {
		noSolution := make([][]int, 9)
		for i := range noSolution {
			noSolution[i] = make([]int, 9)
		}
		return noSolution
	}
}
