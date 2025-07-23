package main

func isValid(grid [][]int, row int, col int, num int) bool {
	// Checks if the suggested number is in the same row
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	//Checks if the suggested number is in the same column
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}

	// To get the box the row and col belongs to
	boxRow := row - (row % 3)
	boxCol := col - (col % 3)

	// Checks is the suggested number is in the same box
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}

	//The number is valid
	return true
}

func solver(board [][]int, row int, col int) bool {
	//Checks if are at the last col
	//If we are it will check if we are at last row too
	//If we are at last row and column, we will return as it end of board
	//Otherwise we will increment row and reset col
	if col == 9 {
		if row == 8 {
			return true
		}
		row++
		col = 0
	}

	//on work on cells which are empty or "0"
	if board[row][col] != 0 {
		return solver(board, row, col+1)
	}

	//Go over every possible number until a possible solution for a cell is found. We used backtracking
	for i := 1; i < 10; i++ {
		//checks if suggested number is valid, if it is we assign
		if isValid(board, row, col, i) {
			board[row][col] = i

			if solver(board, row, col+1) {
				return true
			}
			//if it is not valid, we revert back to zero
			board[row][col] = 0
		}
	}

	return false

}

func SolveSudoku(board [][]int) [][]int {
	solver(board, 0, 0)
	//fmt.Println(board)
	return board
}
