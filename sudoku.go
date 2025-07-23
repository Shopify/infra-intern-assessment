package main

func SolveSuduko(board [][]int) [][]int {
	//This will be similar to a DFS (in a way)
	//The plan can be as follows: assign a value to a cell which does not break any constraint.
	//Build off from there, if a constraint is broken, can revert back to the latest board state that does not violate any constraints -> then assign a different value

	solveEmptyCells(board)
	return board
}

func solveEmptyCells(board [][]int) bool {

	//find first occurance of empty cell
	row, col := getFirstEmptyCell(board)

	//if the board is completed (i.e, no more empty cells). immediately end recursion
	if row == -1 && col == -1 {
		return true
	}

	//fill the empty cell with every possible digit value
	for digit := 1; digit <= 9; digit++ {

		//If using that digit value doesn't violate any rules (constraints)
		if validateRules(row, col, digit, board) {

			//then assign the digit to the empty cell
			board[row][col] = digit

			//finally, take a look at the next empty cell using this new digit value
			//Call solveEmptyCells recursvively.
			if solveEmptyCells(board) {
				return true //if sucess, return to previous recursive level keeping the digit value.
			}

		}
	}

	//end of function. this senario is when no digit can satsify the rules.
	//resets cell back to 0, returns to previous recursive level
	board[row][col] = 0
	return false

}

func getFirstEmptyCell(board [][]int) (int, int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// If cell is empty, return the cell row and col coordinates
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	//return this code if the board is complete (i.e, no more empty cells)
	return -1, -1
}

func validateRules(row int, col int, digit int, board [][]int) bool {
	return !existsInCol(col, board, digit) && !existsInRow(row, board, digit) && !existsInSubGrid(row, col, board, digit)
}

//rule #1: check if digit exists in the respective 3x3 subgrid
func existsInSubGrid(row int, col int, board [][]int, digit int) bool {
	gridY := (row / 3) * 3
	gridX := (col / 3) * 3

	for y := gridY; y < gridY+3; y++ {
		for x := gridX; x < gridX+3; x++ {
			if board[y][x] == digit {
				return true
			}
		}
	}
	return false
}

//rule #2: check if digit exists in same column
func existsInCol(col int, board [][]int, digit int) bool {
	for row := 0; row < 9; row++ {
		if board[row][col] == digit {
			return true
		}
	}
	return false
}

//rule #3: check if digit exists in same row
func existsInRow(row int, board [][]int, digit int) bool {
	for col := 0; col < 9; col++ {
		if board[row][col] == digit {
			return true
		}
	}
	return false
}
