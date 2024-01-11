package main

import "fmt"

// func main() {
// 	tempInput := [][]int{
// 		{5, 3, 0, 0, 7, 0, 0, 0, 0},
// 		{6, 0, 0, 1, 9, 5, 0, 0, 0},
// 		{0, 9, 8, 0, 0, 0, 0, 6, 0},
// 		{8, 0, 0, 0, 6, 0, 0, 0, 3},
// 		{4, 0, 0, 8, 0, 3, 0, 0, 1},
// 		{7, 0, 0, 0, 2, 0, 0, 0, 6},
// 		{0, 6, 0, 0, 0, 0, 2, 8, 0},
// 		{0, 0, 0, 4, 1, 9, 0, 0, 5},
// 		{0, 0, 0, 0, 8, 0, 0, 7, 9},
// 	}

// 	solution := SolveSudoku(tempInput)

// 	printGrid(tempInput)
// 	printGrid(solution)
// }

func SolveSudoku(grid [][]int) [][]int {
	// Create a duplicate grid to not overwrite inputted grid
	duplicateGrid := make([][]int, len(grid))
	for i := range grid {
		duplicateGrid[i] = make([]int, len(grid[i]))
		copy(duplicateGrid[i], grid[i])
	}

	// Solve sudoku using recursive backtracking
	solveSudokuRecursive(duplicateGrid, 0, 0)

	return duplicateGrid
}

func isValid(grid [][]int, row int, col int, num int) bool {
	// Check if the number exists in the row
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	// Check if the number exists in the column
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}

	// Check if the number exists in the current 3x3 cell
	cellRowCorner := row - (row % 3)
	cellColCorner := col - (col % 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[cellRowCorner+i][cellColCorner+j] == num {
				return false
			}
		}
	}

	return true
}

func solveSudokuRecursive(grid [][]int, row int, col int) bool {
	// Check if we went beyond 9 columns
	if col == 9 {
		// if we are at the last row, exit, sudoku was solved
		if row == 8 {
			return true
		}
		// otherwise, move on to the beginning of the next row
		row++
		col = 0
	}

	// Move past cells with an existing number
	if grid[row][col] != 0 {
		return solveSudokuRecursive(grid, row, col+1)
	}

	// Backtracking step, iterate through every possible number at every grid cell until a solution is found
	for i := 1; i < 10; i++ {
		if isValid(grid, row, col, i) {
			grid[row][col] = i

			if solveSudokuRecursive(grid, row, col+1) {
				return true
			}
			grid[row][col] = 0
		}

	}

	return false
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}
