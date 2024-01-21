package main

// SolveSudoku is the main function that wraps core logic
func SolveSudoku(board [][]int) [][]int {
	rowsFilled := make([][]bool, 9)
	colsFilled := make([][]bool, 9)
	boxesFilled := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowsFilled[i] = make([]bool, 10) // Using indices 1-9 for convenience
		colsFilled[i] = make([]bool, 10)
		boxesFilled[i] = make([]bool, 10)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num := board[row][col]
			if num != 0 {
				markFilled(row, col, num, true, rowsFilled, colsFilled, boxesFilled)
			}
		}
	}

	backtrack(0, board, rowsFilled, colsFilled, boxesFilled)
	return board
}

// boxNum calculates the box number based on row and column for a 3x3 Sudoku grid.
func boxNum(row, col int) int {
	return (row/3)*3 + (col / 3)
}

// backtrack attempts to solve the puzzle using backtracking, returning true if solved.
func backtrack(cellNum int, board [][]int, rowsFilled, colsFilled, boxesFilled [][]bool) bool {
	if cellNum == 81 { // Base case: all cells are filled.
		return true
	}

	row, col := cellNum/9, cellNum%9
	if board[row][col] == 0 { // If current cell is empty.
		for num := 1; num <= 9; num++ {
			if canPlace(row, col, num, rowsFilled, colsFilled, boxesFilled) {
				board[row][col] = num
				markFilled(row, col, num, true, rowsFilled, colsFilled, boxesFilled)
				if backtrack(cellNum+1, board, rowsFilled, colsFilled, boxesFilled) {
					return true // Continue with the next cell.
				}
				markFilled(row, col, num, false, rowsFilled, colsFilled, boxesFilled)
				board[row][col] = 0
			}
		}
		return false // No valid number found for this cell. Returns up call stack to try filling previous cells with different numbers.
	} else { // If current cell is already filled by default
		return backtrack(cellNum+1, board, rowsFilled, colsFilled, boxesFilled)
	}
}

// markFilled updates the filled arrays to reflect the status of a number in a row, column, and box.
func markFilled(row, col, num int, filledState bool, rowsFilled, colsFilled, boxesFilled [][]bool) {
	boxIndex := boxNum(row, col)
	rowsFilled[row][num] = filledState
	colsFilled[col][num] = filledState
	boxesFilled[boxIndex][num] = filledState
}

// canPlace checks if a number can be placed in the specified row and column without violating Sudoku rules.
func canPlace(row, col, num int, rowsFilled, colsFilled, boxesFilled [][]bool) bool {
	boxIndex := boxNum(row, col)
	return !rowsFilled[row][num] && !colsFilled[col][num] && !boxesFilled[boxIndex][num]
}
