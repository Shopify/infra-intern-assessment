package main

const BOARD_SIZE = 9

const EMPTY_CELL = 0

func SolveSudoku(board [][]int) [][]int {
	solveWrapper(board)
	return board
}

// solveWrapper is a recursive helper function for solving the Sudoku puzzle. It returns true if it can be solved and false otherwise.
func solveWrapper(board [][]int) bool {

	// loop over every position of the board, ...
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			if board[row][col] == EMPTY_CELL {
				// ... trying to place a valid number from 1 to 9 on every empty cell
				for num := 1; num <= 9; num++ {
					if !hasDuplicateNumbersIn3x3Grid(board, row, col, num) && !hasDuplicateNumbersInLine(board, row, col, num) {
						board[row][col] = num // set the current number when it does not violate sudoku rules...
						if solveWrapper(board) { // ... and determine if it can be part of the solution for the sudoku
							return true
						} else {
							board[row][col] = EMPTY_CELL // if it cannot be part of the solution, reinitialize to empty cell
						}
					}
				}

				// when no number can be placed, backtrack by returning false
				return false
			}
		}
	}

	// when every cell has been filled, return true
	return true
}

// hasDuplicateNumbersInLine returns true if a duplicate exists horizontally/vertically at board[row][col] and false otherwise
func hasDuplicateNumbersInLine(board [][]int, row, col, num int) bool {
	for i := 0; i < BOARD_SIZE; i++ {
		if board[row][i] == num || board[i][col] == num {
			return true
		}
	}
	return false
}

// hasDuplicateNumbersIn3x3Grid returns true if a duplicate exists in a 3x3 subgrid and false otherwise
// It gets the middle of the 3x3 grid and iterates over all 9 possible locations through coordinates
func hasDuplicateNumbersIn3x3Grid(board [][]int, row, col, num int) bool {
	middleCellRowValue, middleCellColValue := row - row%3 + 1, col - col%3 + 1
	offsets := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, offset := range offsets {
		row, col = middleCellRowValue + offset[0], middleCellColValue + offset[1]
		if board[row][col] == num {
			return true
		}
	}
	return false
}

