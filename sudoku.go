package main

import "fmt"

// REQUIRES: existing board is valid
// 			 0 <= row, col <= 8; board[row][col] is unfilled
//			 1 <= num <= 9
// MODIFIES: N/A
// EFFECTS: Determines if placing num at board[row][col] is valid
func Promising(board [][]int, row int, col int, num int) bool {
	// check all cells on the same row and same column as the cell at [row][col]
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}
	// check all cells in the same 3 by 3 grid as the cell at [row][col]
	startRowIdx := (row / 3) * 3
	startColIdx := (col / 3) * 3
	// fmt.Printf("Start row idx: %d, Start col idx: %d\n", startRowIdx, startColIdx)
	for i := startRowIdx; i < startRowIdx + 3; i++ {
		for j := startColIdx; j < startColIdx + 3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
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
//			 0 <= row, col <= 8
// MODIFIES: board
// EFFECTS: Performs a (recursive) backtracking algorithm to solve the sudoku
func SolveSudokuHelper(board [][]int, row int, col int) bool {
	// If sudoku is solved, return true
	// All cells filled -> Solved!
	if row < 0 {
		return true
	}
	// Otherwise
	// Try to place number 1 through 9 at current row and col
	for num := 1; num <= 9; num++ {
		// Check constraint: If placing this number does not violate any rules of sudoku
		if Promising(board, row, col, num) {
			// Place number
			board[row][col] = num
			// Recursion with the next unfilled cell on the board
			nextRow, nextCol := FindNextUnfilled(board, row, col)
			// If recursive function call returned true
			if SolveSudokuHelper(board, nextRow, nextCol) {
				// then it means puzzle solved so return true
				return true
			}
			// Otherwise, Backtrack: "unplace" the number at current cell
			board[row][col] = 0
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
	// for storing the index of the first unfilled cell
	row := -1
	col := -1
    for i := 0; i < len(sudoku); i++ {
        for j := 0; j < len(sudoku[i]); j++ {
			// Preprocess the board by filling all unfilled values with 0
			if sudoku[i][j] < 1 || sudoku[i][j] > 9 {
				sudoku[i][j] = 0
				// Find the index of the first unfilled cell
				if row == -1 {
					row = i
					col = j
				}
			}
        }
    }
	// calls helper function which does the backtracking
	// pass sudoku 2d array by reference
	SolveSudokuHelper(sudoku, row, col)
	return sudoku
}

func testPromising() {
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
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 2, 0, 5))
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 2, 0, 7))
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 2, 0, 4))
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 2, 0, 9))
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 2, 0, 3))
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 2, 0, 8))
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 2, 0, 6))
	fmt.Printf("Expected: %t; Got: %t\n", true, Promising(input, 2, 0, 1))
	fmt.Printf("Expected: %t; Got: %t\n", true, Promising(input, 2, 0, 2))

	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 7, 6, 7))
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 7, 6, 9))
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 7, 6, 8))

	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 8, 5, 4))
	fmt.Printf("Expected: %t; Got: %t\n", false, Promising(input, 8, 5, 1))
	fmt.Printf("Expected: %t; Got: %t\n", true, Promising(input, 8, 5, 6))
	fmt.Printf("Expected: %t; Got: %t\n", true, Promising(input, 8, 5, 2))
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

    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[i]); j++ {
            fmt.Printf("%d ", input[i][j])
        }
        fmt.Println()
    }

	// testPromising()
}
