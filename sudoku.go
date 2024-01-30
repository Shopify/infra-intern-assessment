package main

func SolveSudoku(board [][]int) [][]int {
	// makes copy of the original board
	solved := make([][]int, len(board))
	for i := range board {
		solved[i] = make([]int, len(board[i]))
		copy(solved[i], board[i])
	}

	// start solving from the top-left corner
	if solveSudokuPos(solved, 0, 0) {
		return solved
	}
	return [][]int{}
}

// this function recursively solves the Sudoku puzzle
func solveSudokuPos(board [][]int, row, col int) bool {
	// find the next empty position on the board
	solved, row, col := getNextPos(board, row, col)
	if solved {
		return true // if no empty position found, end recursion -> puzzle is solved
	}

	// tries nums 1 to 9 in the current position
	for i := 1; i <= 9; i++ {
		// check if current num is valid
		if isValidValue(board, i, row, col) {
			putValue(board, i, row, col)                // add the number to the board
			solved := solveSudokuPos(board, row, col+1) // recursively solve for next position
			if solved {
				return true // if Sudoku is solved from this position,  end recursion -> puzzle is solved
			}
			removeValue(board, row, col) // backtrack if solution not found
		}
	}
	return false // return false if no valid number can be placed
}

// this function finds the next empty position on the board
func getNextPos(board [][]int, row, col int) (bool, int, int) {
	// search rows and columns
	for row < 9 {
		for col < 9 {
			if board[row][col] == 0 {
				return false, row, col // return the empty position
			}
			col++
		}
		row++
		col = 0
	}
	return true, row, col // return true if board is full
}

// this function removes a value from the specified position
func removeValue(board [][]int, row, col int) {
	board[row][col] = 0
}

// this function places a value at the specified position
func putValue(board [][]int, val, row, col int) {
	board[row][col] = val
}

// this function checks if a value can be placed at the specified position
func isValidValue(board [][]int, val, row, col int) bool {
	// checks for value in row
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}

	// checks for value in column
	for i := 0; i < 9; i++ {
		if board[i][col] == val {
			return false
		}
	}

	// checks for value in box
	rowStart, rowEnd := (row/3)*3, (row/3+1)*3
	colStart, colEnd := (col/3)*3, (col/3+1)*3
	for i := rowStart; i < rowEnd; i++ {
		for j := colStart; j < colEnd; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}

	return true // value placed if no conflicts
}
