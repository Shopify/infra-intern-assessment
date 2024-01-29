package main

// Constants representing values in the Sudoku board
const emptyCell int = 0
const boardSize int = 9

//
func isNumberInRow(board [][]int, row, col, value int) bool {
	for columnPosition := 0; columnPosition < boardSize; columnPosition++ {
		if board[row][columnPosition] == value {
			return false
		}
	}
	return true
}

// isNumberInRow checks if the given value is present in the specified row of the Sudoku board
func isNumberInColumn(board [][]int, row, col, value int) bool {
	for rowPosition := 0; rowPosition < boardSize; rowPosition++ {
		if board[rowPosition][col] == value {
			return false
		}
	}
	return true
}

// isNumberInColumn checks if the given value is present in the specified column of the Sudoku board
func findEmptyCell(board [][]int) (int, int) {
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if board[row][col] == emptyCell {
				return row, col
			}
		}
	}
	return -1, -1 // Return -1, -1 if no empty cell is found
}

// isNumberInBox checks if the given value is present in the 3x3 box containing the specified cell
func isNumberInBox(board [][]int, row, col, value int) bool {
	startingBoxRow, startingBoxColumn := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startingBoxRow][j+startingBoxColumn] == value {
				return false
			}
		}
	}
	return true
}

// isNumberValid checks if placing the given value in the specified cell is a valid move according to Sudoku rules
func isNumberValid(board [][]int, row, col, value int) bool {
	return (isNumberInRow(board, row, col, value) &&
		isNumberInColumn(board, row, col, value) &&
		isNumberInBox(board, row, col, value))
}

// SolveSudoku attempts to solve the Sudoku puzzle using backtracking algorithm
func SolveSudoku(board [][]int) [][]int {
	emptyRow, emptyColumn := findEmptyCell(board)

	if emptyRow == -1 && emptyColumn == -1 {
		return board // If no empty cell is found, the board is solved
	}

	for number := 1; number <= boardSize; number++ {
		if isNumberValid(board, emptyRow, emptyColumn, number) {
			board[emptyRow][emptyColumn] = number

			// Recursively attempt to solve the Sudoku with the updated board
			if SolveSudoku(board) != nil {
				return board
			}

			board[emptyRow][emptyColumn] = emptyCell // Backtrack if the current placement leads to no solution
		}
	}
	return nil // Return nil if no solution is found for the current board configuration
}
