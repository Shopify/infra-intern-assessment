package main

import "fmt"

// SolveSudoku solves the Sudoku puzzle using backtracking and recursion.
// It takes a SudokuGrid as input and returns the solved grid
func SolveSudoku(grid [][]int) [][]int {
	// Find and store the coordinates of empty cells in a stack
	emptyCells := findEmptyCells(grid)

	// Attempt to solve the Sudoku puzzle using backtracking and recursion
	solvedGrid, success := solveSudokuHelper(grid, emptyCells)

	// If successful, return the solved grid; otherwise, print a message
	if success {
		return solvedGrid
	}
	fmt.Println("\nNo solution exists.")
	return [][]int{}
}

// solveSudokuHelper is a recursive helper function for solving the Sudoku puzzle.
// It uses backtracking to explore possible solutions.
func solveSudokuHelper(grid [][]int, emptyCells Stack) ([][]int, bool) {
	// If there are no more empty cells, the puzzle is solved
	if len(emptyCells) == 0 {
		return grid, true
	}

	// Pop an empty cell from the stack
	cell, _ := emptyCells.Pop()
	row, col := cell.Row, cell.Col

	// Try placing numbers from 1 to gridSize in the current cell
	for num := 1; num <= gridSize; num++ {
		// Check if placing 'num' in the current cell is valid
		if IsValidPlacement(grid, row, col, num) {
			// Place 'num' in the current cell
			grid[row][col] = num

			// Recursively attempt to solve the puzzle
			solvedGrid, success := solveSudokuHelper(grid, emptyCells)
			if success {
				return solvedGrid, true
			}

			// If the current placement doesn't lead to a solution, backtrack
			grid[row][col] = 0
		}
	}

	// If no valid placement for the current cell, backtrack
	emptyCells.Push(cell)

	return grid, false
}
