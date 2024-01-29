package main

import (
	"fmt"
)

const N = 9

// convertSliceToArray converts a 2D slice to a 2D array
func convertSliceToArray(slice [][]int) [N][N]int {
	var array [N][N]int
	for i := range slice {
		for j := range slice[i] {
			array[i][j] = slice[i][j]
		}
	}
	return array
}

// convertArrayToSlice converts a 2D array to a 2D slice
func convertArrayToSlice(array [N][N]int) [][]int {
	slice := make([][]int, N)
	for i := range array {
		slice[i] = make([]int, N)
		for j := range array[i] {
			slice[i][j] = array[i][j]
		}
	}
	return slice
}

// To check if it iss safe to place a number in the given position
func isSafe(grid *[N][N]int, row, col, num int) bool {
	// Check if 'num' is not in the given row
	for x := 0; x < N; x++ {
		if grid[row][x] == num {
			return false
		}
	}

	// Check if 'num' is not in the given column
	for x := 0; x < N; x++ {
		if grid[x][col] == num {
			return false
		}
	}

	// Check if 'num' is not in the given 3x3 box
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

// Internal function to solve sudoku
func solveSudoku(grid *[N][N]int) bool {
	row := -1
	col := -1
	isEmpty := true
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				row = i
				col = j
				isEmpty = false
				break
			}
		}
		if !isEmpty {
			break
		}
	}

	if isEmpty {
		return true
	}

	for num := 1; num <= N; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num
			if solveSudoku(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

// SolveSudoku solves the Sudoku puzzle
// It accepts and returns a slice of slices

func SolveSudoku(grid [][]int) [][]int {
	gridArray := convertSliceToArray(grid)
	if solveSudoku(&gridArray) {
		return convertArrayToSlice(gridArray)
	}
	return [][]int{} // return an empty slice if no solution
}

func main() {
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

	solved := SolveSudoku(grid)
	for _, row := range solved {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
