package main

import (
	"fmt"
)

const N = 9 // Constant for the size of the Sudoku board

// CanPlaceNumber checks if a number can be placed in a given row, column, and 3x3 square
// without violating Sudoku rules.
func CanPlaceNumber(num int, row, col int, rows, cols, squares []map[int]bool) bool {
	// Check if the number is not already in the given row, column, and square
	return !rows[row][num] && !cols[col][num] && !squares[3*(row/3)+col/3][num]
}

// InitializeBoard sets up the board by filling maps with existing numbers on the board.
// This helps in checking if a number can be placed in a certain position.
func InitializeBoard(board [][]int) ([]map[int]bool, []map[int]bool, []map[int]bool) {
	// Create slices of maps for rows, columns, and squares
	rows := make([]map[int]bool, N)
	cols := make([]map[int]bool, N)
	squares := make([]map[int]bool, N)

	// Initialize maps
	for i := 0; i < N; i++ {
		rows[i] = make(map[int]bool)
		cols[i] = make(map[int]bool)
		squares[i] = make(map[int]bool)
	}

	// Fill the maps with existing numbers on the board
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			num := board[i][j]
			if num != 0 {
				rows[i][num] = true
				cols[j][num] = true
				squares[(i/3)*3+j/3][num] = true
			}
		}
	}

	return rows, cols, squares
}

// SolveSudoku solves the Sudoku puzzle using a backtracking approach.
func SolveSudoku(board [][]int) [][]int {
	// Create a copy of the board to work on
	boardCopy := make([][]int, len(board))
	for i := range board {
		boardCopy[i] = make([]int, len(board[i]))
		copy(boardCopy[i], board[i])
	}

	// Initialize maps to keep track of existing numbers
	rows, cols, squares := InitializeBoard(boardCopy)

	// Recursive function to solve the puzzle
	var solve func(row, col int) bool
	solve = func(row, col int) bool {

		// If we have reached the end of the board, the puzzle is solved
		if row == N {
			return true
		}

		// Determine the next cell to fill
		nextRow, nextCol := row, col+1
		if col == N-1 {
			nextRow, nextCol = row+1, 0
		}

		// Skip filled cells
		if boardCopy[row][col] != 0 {
			return solve(nextRow, nextCol)
		}

		// Try placing numbers 1-9 in the current cell
		for num := 1; num <= N; num++ {
			if CanPlaceNumber(num, row, col, rows, cols, squares) {
				// Place the number and update the state
				boardCopy[row][col] = num
				rows[row][num] = true
				cols[col][num] = true
				squares[3*(row/3)+col/3][num] = true

				// Recursively solve the rest of the board
				if solve(nextRow, nextCol) {
					return true
				}

				// Backtrack if the number placed doesn't lead to a solution
				boardCopy[row][col] = 0
				delete(rows[row], num)
				delete(cols[col], num)
				delete(squares[3*(row/3)+col/3], num)
			}
		}

		// Return false if no number can be placed in the current cell
		return false
	}

	// Start the solving process from the top-left cell
	if solve(0, 0) {
		return boardCopy // Return the solved board
	} else {
		return nil // Return nil if the puzzle cannot be solved
	}
}

// printBoard displays the Sudoku board in a readable format.
func printBoard(board [][]int) {
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
