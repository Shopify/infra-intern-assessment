package main

import "fmt"

// REQUIRES: memoRow, memoCol, memoGrid accurately reflect the current board
// 			 0 <= row, col <= 8
//			 memo corresponding to [row][col] is currently not updated (false)
//			 1 <= num <= 9
// MODIFIES: N/A
// EFFECTS: Determines if placing num at board[row][col] is valid
func Promising(memoRow [9][10]bool, memoCol [9][10]bool,
			   memoGrid [3][3][10]bool,row int, col int, num int) bool {
	// Return true iff num does not exist in the row, column, or grid
	// which cell [row][col] corresponds to
	return !memoRow[row][num] && !memoCol[col][num] && !memoGrid[row/3][col/3][num]
}

// REQUIRES: Existing board is valid
// 			 0 <= curRow, curCol <= 8
//			 curRow, curCol corresponds to the cell that was just filled
// MODIFIES: N/A
// EFFECTS: Returns the row and column index of the next unfilled cell
//			Returns -1, -1 if all cells are filled
func FindNextUnfilled(board [][]int, curRow int, curCol int) (int, int) {
	// Look for unfilled cells in current row first
	for j := curCol + 1; j < 9; j++ {
		if board[curRow][j] == 0 {
			return curRow, j
		}
	}
	// Then move on to the next row(s)
	for i := curRow + 1; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	// No unfilled cell found
	return -1, -1
}

// REQUIRES: board is a 2d array representing a sudoku with exactly 1 solution
//			 memoRow, memoCol, memoGrid accurately reflect the current board
//			 0 <= row, col <= 8
// MODIFIES: board
// EFFECTS: Performs a (recursive) backtracking algorithm to solve the sudoku
func SolveSudokuHelper(board [][]int, memoRow [9][10]bool, memoCol [9][10]bool,
					   memoGrid [3][3][10]bool, row int, col int) bool {
	// If sudoku is solved, i.e. all cells have been filled
	if row < 0 {
		return true
	}
	// Otherwise
	// Try to place number 1 through 9 at current [rol][col]
	for num := 1; num <= 9; num++ {
		// Check constraint: If placing this number does not violate any rules of sudoku
		if Promising(memoRow, memoCol, memoGrid, row, col, num) {
			// Place number & update memos
			board[row][col] = num
			memoRow[row][num] = true
			memoCol[col][num] = true
			memoGrid[row/3][col/3][num] = true
			// Recursion with the next unfilled cell on the board
			nextRow, nextCol := FindNextUnfilled(board, row, col)
			// If recursive function call returned true
			if SolveSudokuHelper(board, memoRow, memoCol, memoGrid, nextRow, nextCol) {
				// Then it means puzzle solved
				return true
			}
			// Otherwise, Backtrack: "unplace" the number at current cell
			board[row][col] = 0
			memoRow[row][num] = false
			memoCol[col][num] = false
			memoGrid[row/3][col/3][num] = false
		}
	}
	// Reaching this point means that none of the numbers from 1 to 9 can be placed at current cell
	// i.e. the number in a previous cell needs to change, thus Backtrack -> return to the caller!
	return false
}

// REQUIRES: The input grid will be a 9x9 two-dimensional array of integers.
//			 The input grid will have exactly one solution.
// MODIFIES: N/A
// EFFECTS: Solves the sudoku, returns a 9 by 9 array of the solved sudoku
func SolveSudoku(sudoku [][]int) [][]int {
	// For storing the index of the first unfilled cell
	row := -1
	col := -1
	// For storing whether a number in a row, col, grid is already placed
	// All elements in array automatically initialized to false
	var rowsMemo, colsMemo [9][10]bool
	var gridsMemo [3][3][10]bool

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
			// Get number at current cell
			num := sudoku[i][j]
			// Preprocess the board by filling all unfilled values with 0
			if num < 1 || num > 9 {
				sudoku[i][j] = 0
				// Find the index of the first unfilled cell
				if row == -1 {
					row = i
					col = j
				}
			// Preprocess which nums already exist in current row/col/grid
			} else {
				// Update memo to true to indicate that num exists in this row/col/grid
				rowsMemo[i][num] = true
				colsMemo[j][num] = true
				gridsMemo[i/3][j/3][num] = true
			}
        }
    }
	// Calls helper function which does the backtracking
	// Pass sudoku 2d array and 3 memos by reference
	SolveSudokuHelper(sudoku, rowsMemo, colsMemo, gridsMemo, row, col)
	return sudoku
}

// REQUIRES: The input grid will be a 9x9 two-dimensional array of integers.
// MODIFIES: N/A
// EFFECTS: Prints the sudoku board
func PrintSudoku(sudoku [][]int) {
	for i := 0; i < len(sudoku); i++ {
        for j := 0; j < len(sudoku[i]); j++ {
            fmt.Printf("%d ", sudoku[i][j])
        }
        fmt.Println()
    }
}
