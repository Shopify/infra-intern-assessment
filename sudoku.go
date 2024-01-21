package main

import "fmt"

// REQUIRES: board is a 2d array representing a sudoku with exactly 1 solution
//			 0 <= row, col <= 8
// MODIFIES: board
// EFFECTS: Performs a (recursive) backtracking algorithm to solve the sudoku
func SolveSudokuHelper(board [][]int, row int, col int) bool {
	// If sudoku is solved, return true
	// Otherwise
	// Try to place number 1 through 9 at current row and col
		// Check constraint: If placing this number does not violate any rules of sudoku
			// Recursion with the next unfilled cell on the board
			// If recursive function call returned true, then it means puzzle solved so return true
		// Otherwise, Backtrack: "unplace" the number at current cell
	// Reaching this point means that none of the numbers from 1 to 9 can be placed at current cell
	// i.e. the number in a previous cell needs to change, thus Backtrack -> return to the caller!
	return false
}

// Requires: The input grid will be a 9x9 two-dimensional array of integers.
//			 The input grid will have exactly one solution.
// Modifies: N/A
// Effects: Solves the sudoku, returns a 9 by 9 array of the solved sudoku
func SolveSudoku(sudoku [][]int) [][]int {
    for i := 0; i < len(sudoku); i++ {
        for j := 0; j < len(sudoku[i]); j++ {
            fmt.Printf("%d ", sudoku[i][j])
        }
        fmt.Println()
    }
	// calls helper function which does the backtracking
	// pass sudoku 2d array by reference
	SolveSudokuHelper(sudoku, 0, 0)
	return sudoku
}

func main() {
	input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	SolveSudoku(input)
}
