package main

import "log"

// SolveSudoku function returns boolean determining whether the Sudoku puzzle is solved or not.
func solve(grid [][]int) bool {
	var row int
	var col int
	emptyCell := true

	// Find an empty cell
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				row, col = i, j
				emptyCell = false
				break
			}
		}
		if !emptyCell {
			break
		}
	}

	// No empty cell = solved puzzle
	if emptyCell {
		return true
	}

	// Try digits 1 to 9 in the empty cell
	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num

			if solve(grid) {
				return true
			}

			// Backtrack
			grid[row][col] = 0
		}
	}
	return false
}

// function to check if number is valid to place in that specific cell.
func isSafe(grid [][]int, row, col, num int) bool {
	// Checks the row and column
	for x := 0; x < 9; x++ {
		if grid[row][x] == num || grid[x][col] == num {
			return false
		}
	}

	// Check the 3x3 square
	startRow := row - (row % 3)
	startCol := col - (col % 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

// SolveSudoku solves the Sudoku puzzle using a backtracking algorithm.
func SolveSudoku(board [][]int) [][]int {
	if !solve(board) {
		log.Fatal("Failed to solve the sudoku puzzle. Printing the grid in it's current state ")
	}

	log.Println("Successfully solved the sudoku puzzle.")
	return board
}
