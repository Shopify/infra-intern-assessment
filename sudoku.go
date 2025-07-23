package main

import (
	"fmt"
)

const N = 9

// SolveSudoku solves the Sudoku puzzle
func SolveSudoku(grid [][]int) [][]int {
	if solveSudokuHelper(grid) {
		return grid
	}
	return nil //if the puzzle is unsolvable
}

// solveSudokuHelper is a helper function uses backtracking
func solveSudokuHelper(grid [][]int) bool {
	row, col := findEmptyCell(grid)
	if row == -1 && col == -1 {
		return true // Puzzle solved
	}

	for num := 1; num <= N; num++ {
		if isValid(grid, row, col, num) {
			grid[row][col] = num
			if solveSudokuHelper(grid) {
				return true
			}
			grid[row][col] = 0 // Backtrack
		}
	}
	return false
}

// findEmptyCell returns the row and column of an empty cell
func findEmptyCell(grid [][]int) (int, int) {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if grid[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

// isValid checks if a number can be placed in the given cell
func isValid(grid [][]int, row, col, num int) bool {
	// Check row and column
	for i := 0; i < N; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	// Check 3x3 sub-grid
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

// printGrid prints the Sudoku grid
func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
