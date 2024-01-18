// Name: sudoku.go
// Author: Jay Turnsek
// Date Modified: Jan 18 2024

package main

// findEmpty gets the next empty cell in the Sudoku board. 
// takes board and loc as arguments, but returns a bool stating if a 
// empty cell was found; loc is modified in the case that there IS an empty cell.
func findEmpty(board [][]int, loc []int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (board[i][j] == 0) {

				// empty cell found, update loc and return true
				loc[0] = i
				loc[1] = j
				return true
			}
		}
	}

	// no empty cells
	return false
}

// validRow checks if a given value can be placed in the row; simply by
// checking if the row (subarray) contains the value 'val'.
func validRow(board [][]int, row int, val int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}
	return true
}

// validCol checks if a given value can be placed in the column; simply by
// checking if each row's i'th value is equal to val.
func validCol(board [][]int, col int, val int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == val {
			return false
		}
	}
	return true
}

// validBox checks if a given value can be placed in the current box.
// the adjusted values force the indexes into the top left corner of the box
// ie (4, 7) => (3, 6), as (4, 7) is within the middle right box, but (3, 6)
// is the top left of that box. 
// If the value 'val' is found within this box, the function returns false; 
// returns true otherwise.
func validBox(board [][]int, row int, col int, val int) bool {
	
	// this just puts the indexes to the top left of the box.
	row_adjusted := row - (row % 3)
	col_adjusted := col - (col % 3)

	// check all values in box for val
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i + row_adjusted][j + col_adjusted] == val {

				// not valid :(
				return false
			}
		}
	}

	// it's valid!
	return true
}

// locationValid is just a grouping function that ensures all Sudoku constraints are
// satisfied for a value entry into a specific row and column. All 3 must be true for a 
// move to be valid.
func locationValid(board [][]int, row int, col int, val int) bool {
	return validRow(board, row, val) && validCol(board, col, val) && validBox(board, row, col, val)
}

// solve executes the actual backtracking algorithm. the basic steps are as follows:
// 1. find the next empty cell (0)
// 2. Try to fit values from 1-9 into that cell
// 3. If there are no valid values to put into the cell, we backtrack to the previous call
// 4. If there is a valid value, recursively call solve to the next available cell.
// Once there are no empty cells, we are done and the Sudoku is solved.
func solve(board [][]int) bool {
	loc := []int{0, 0}

	// No empty cells, we're done and its solved
	if !findEmpty(board, loc) {
		return true
	}

	// unpack row and col values
	row, col := loc[0], loc[1]

	// try all values in cell 0 - 9
	for n := 1; n <= 9; n++ {
		
		// if the move is valid, update the cell.
		if locationValid(board, row, col, n) {
			board[row][col] = n
			
			// recursively call solving function, until all cells are filled
			if solve(board) {
				return true
			}
			
			// if we reached here, the move led to no valid moves; reset to 0
			// and continue to next value from 0 - 9
			board[row][col] = 0
		}
	}

	// this triggers the backtracking
	return false
}

// SolveSudoku just wraps the solve function, as by it's recursive/backtracking nature
// its easier to have a helper function that copies the board to apply the function's changes;
// the solved board is returned as the same type as the board passed into it for the purposes of testing.
func SolveSudoku(board [][]int) [][]int {

	// copy board, make changes to solve, return solved board.
	solvedBoard := board
	solve(solvedBoard)
	return solvedBoard
}