package main

import (
	"fmt"
)

// Define the size of the Sudoku grid
const gridSize = 9

// SolveSudoku solves the Sudoku puzzle using backtracking
func SolveSudoku(grid [][]int) bool {
	// Find an empty cell
	row, col := findEmptyCell(grid)

	// If there are no empty cells, the puzzle is solved
	if row == -1 && col == -1 {
		return true
	}

	// Try filling the empty cell with numbers from 1 to 9
	for num := 1; num <= gridSize; num++ {
		if isSafe(grid, row, col, num) {
			// Place the number in the cell
			grid[row][col] = num

			// Recursively try to solve the remaining puzzle
			if SolveSudoku(grid) {
				return true
			}

			// If placing the number doesn't lead to a solution, backtrack
			grid[row][col] = 0
		}
	}

	// If no number can be placed, backtrack
	return false
}

// findEmptyCell finds the first empty cell in the Sudoku grid
func findEmptyCell(grid [][]int) (int, int) {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if grid[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

// isSafe checks if it's safe to place a number in a specific cell
func isSafe(grid [][]int, row, col, num int) bool {
	// Check if the number is not present in the current row and column
	for i := 0; i < gridSize; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	// Check if the number is not present in the current 3x3 subgrid
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

// PrintGrid prints the Sudoku grid
func PrintGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func main() {
	// Example Sudoku grid
	inputGrid := [][]int{
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

	fmt.Println("Input Sudoku:")
	PrintGrid(inputGrid)

	// Solve the Sudoku puzzle
	if SolveSudoku(inputGrid) {
		fmt.Println("\nSolved Sudoku:")
		PrintGrid(inputGrid)
	} else {
		fmt.Println("\nNo solution found.")
	}
}
