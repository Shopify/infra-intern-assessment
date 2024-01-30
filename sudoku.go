package main

import "fmt"

const size = 9

// Function to check if a number can be placed in a particular cell
func isSafe(board [][]int, row, col, num int) bool {
	// Check if the number is not already present in the row and column
	for i := 0; i < size; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Check if the number is not already present in the 3x3 sub-grid
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

// Function to solve the Sudoku puzzle
func SolveSudoku(board [][]int) [][]int {
	var row, col int
	// Find an empty cell
	found := false
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	// If no empty cell is found, Sudoku is solved
	if !found {
		return board
	}

	// Try placing numbers 1 to 9 in the empty cell
	for num := 1; num <= size; num++ {
		if isSafe(board, row, col, num) {
			board[row][col] = num
			// Recursively solve Sudoku
			if SolveSudoku(board) != nil {
				return board
			}
			// Backtrack if placing the number leads to a solution
			board[row][col] = 0
		}
	}

	// If no number can be placed, trigger backtracking
	return nil
}

func main() {
	// Example input Sudoku puzzle
	board := [][]int{
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

	// Solve Sudoku
	solved := SolveSudoku(board)

	// Print solved Sudoku
	for i := 0; i < size; i++ {
		fmt.Println(solved[i])
	}
}
