package main

import "fmt"

const dim = 9
const emptyCell = 0

//  solves the Sudoku puzzle using backtracking
func solve(grid [][]int) bool {
	row, col, found := findEmptyCell(grid)

	// check if successful
	if !found {
		return true
	}

	// Filling empty cells
	for num := 1; num <= dim; num++ {
		if checkIfSafe(grid, row, col, num) {
			grid[row][col] = num
			// recursive step to solve puzzle
			if solve(grid) {
				return true
			}

			// If solution is invalid, backtrack
			grid[row][col] = emptyCell
		}
	}
	return false
}

// finds the position of empty cells
func findEmptyCell(grid [][]int) (int, int, bool) {
	for row := 0; row < dim; row++ {
		for col := 0; col < dim; col++ {
			if grid[row][col] == emptyCell {
				return row, col, true
			}
		}
	}
	return -1, -1, false
}

// checks if specified value is present in the given row
func checkRow(grid [][]int, row int, val int) bool {
	for columnIndex := 0; columnIndex < dim; columnIndex++ {
		if grid[row][columnIndex] == val {
			return true
		}
	}
	return false
}

// checks if specified value is present in the given column
func checkCol(grid [][]int, col int, val int) bool {
	for rowIndex := 0; rowIndex < dim; rowIndex++ {
		if grid[rowIndex][col] == val {
			return true
		}
	}
	return false
}

// checks if specified value is present in the given section
func checkSection(grid [][]int, row int, col int, val int) bool {
	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		for colIndex := 0; colIndex < 3; colIndex++ {
			if grid[rowIndex+row][colIndex+col] == val {
				return true
			}
		}
	}
	return false
}

// checks whether the given value is able to be placed in the given position
func checkIfSafe(grid [][]int, row int, col int, val int) bool {
    return grid[row][col] == emptyCell && !checkRow(grid, row, val) && !checkCol(grid, col, val) && !checkSection(grid, row-row%3, col-col%3, val)
}

// prints the grid
func printGrid(grid [][]int) {
	for rowIndex := 0; rowIndex < dim; rowIndex++ {
		for colIndex := 0; colIndex < dim; colIndex++ {
			fmt.Printf("%d ", grid[rowIndex][colIndex])
		}
		fmt.Println()
	}
}

// solves the puzzle
func SolveSudoku(input [][]int) [][]int {
	if !solve(input) {
		return nil
	}

	return input
}

// main function, used for debugging
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
	if solved != nil {
		printGrid(solved)
		return
	}
	fmt.Println("Cannot be solved")
}
