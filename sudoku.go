package main

import "math"

// Size of grid
const N = 9

// Size of sub-grid
var rootN = int(math.Sqrt(N))

func isValidEntry(grid [][]int, row int, col int, num int) bool {
	return !(isInCol(grid, col, num) || isInRow(grid, row, num) || isInSubGrid(grid, row, col, num))
}

func isInCol(grid [][]int, col int, num int) bool {
	for row := 0; row < N; row++ {
		if grid[row][col] == num {
			return true
		}
	}

	return false
}

func isInRow(grid [][]int, row int, num int) bool {
	for col := 0; col < N; col++ {
		if grid[row][col] == num {
			return true
		}
	}

	return false
}

func isInSubGrid(grid [][]int, row int, col int, num int) bool {

	// Define top left cell coordinate of sub-grid
	rowStart := row / rootN * rootN
	colStart := col / rootN * rootN

	for rowIndex := 0; rowIndex < rootN; rowIndex++ {
		for colIndex := 0; colIndex < rootN; colIndex++ {
			if grid[rowStart+rowIndex][colStart+colIndex] == num {
				return true
			}
		}
	}

	return false
}

func SolveSudokuHelper(grid [][]int, row int, col int) bool {

	// End of row (move to beginning of next row)
	if col == N {
		row++
		col = 0
	}

	// End of grid
	if row == N && col == 0 {
		return true
	}

	// Skip pre-filled cell
	if grid[row][col] != 0 {
		return SolveSudokuHelper(grid, row, col+1)
	}

	// Attempt to fill empty cell
	for num := 1; num <= N; num++ {
		if isValidEntry(grid, row, col, num) {

			grid[row][col] = num

			if SolveSudokuHelper(grid, row, col+1) {
				return true
			}

			grid[row][col] = 0
		}
	}

	return false
}

func SolveSudoku(grid [][]int) [][]int {
	if SolveSudokuHelper(grid, 0, 0) {
		return grid

	}

	return nil
}
