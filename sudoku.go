package main

import "fmt"

const N = 9

// SolveSudoku solves the Sudoku puzzle. It returns the solved grid.
func SolveSudoku(grid [][]int) [][]int {
	if solveSudokuHelper(grid) {
		return grid // Return the solved grid
	}
	return grid // Return the original grid if unsolvable
}

// solveSudokuHelper is a helper function that applies the backtracking algorithm.
func solveSudokuHelper(grid [][]int) bool {
	var row, col int
	if !findEmptyLocation(grid, &row, &col) {
		return true // Puzzle solved
	}

	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num
			if solveSudokuHelper(grid) {
				return true
			}
			grid[row][col] = 0 // Failure, unmake & try again
		}
	}
	return false // Triggers backtracking
}

// findEmptyLocation searches for a location that is still unassigned
func findEmptyLocation(grid [][]int, row, col *int) bool {
	for *row = 0; *row < N; *row++ {
		for *col = 0; *col < N; *col++ {
			if grid[*row][*col] == 0 {
				return true
			}
		}
	}
	return false
}

// isSafe checks if it will be legal to assign num to the given row, col
func isSafe(grid [][]int, row, col, num int) bool {
	return !usedInRow(grid, row, num) && !usedInCol(grid, col, num) && !usedInBox(grid, row-row%3, col-col%3, num)
}

// usedInRow checks if a number is in the given row
func usedInRow(grid [][]int, row, num int) bool {
	for col := 0; col < N; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// usedInCol checks if a number is in the given column
func usedInCol(grid [][]int, col, num int) bool {
	for row := 0; row < N; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// usedInBox checks if a number is in the given 3x3 box
func usedInBox(grid [][]int, boxStartRow, boxStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[row+boxStartRow][col+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}

// printGrid prints the grid
func printGrid(grid [][]int) {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			fmt.Printf("%2d", grid[row][col])
		}
		fmt.Println()
	}
}

func main() {
	// Example puzzle
	grid := [][]int{
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

	grid = SolveSudoku(grid)
	printGrid(grid)
}
