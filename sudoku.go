// Hi there! This is a Sudoku solver written by  David Oche Gideon for the shopify internship application(Infrastucture Engineering).
// I have included the test file for the program below 
// I have also included the original puzzle and the solved puzzle below
// this uses a backtracking algorithm to solve the puzzle due to the fact that it is an NP complete problem
// Backtracking is a technique used in algorithm design and
// is often associated with solving problems through a systematic exploration of possible solutions.

package main

import (
	"fmt"
)

const gridSize = 9

// it displays the Sudoku grid.
func PrintGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

// this function solves the Sudoku puzzle using backtracking algo.
func SolveSudoku(grid [][]int) bool {
	emptyCell := findEmptyCell(grid)

	// If there are no empty cells, the puzzle is solved.
	if emptyCell[0] == -1 {
		return true
	}

	row, col := emptyCell[0], emptyCell[1]

	// Try placing numbers from 1 to 9 in the empty cell.
	for num := 1; num <= gridSize; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num

			// loop attempt to solve the Sudoku with the current choice.
			if SolveSudoku(grid) {
				return true
			}

			// If the current choice doesn't lead to a solution, backtrack.
			grid[row][col] = 0
		}
	}

	// if no solution found for the current state.
	return false
}

// this function returns the coordinates of the first empty cell in the grid.
func findEmptyCell(grid [][]int) [2]int {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 0 {
				return [2]int{i, j}
			}
		}
	}
	// No empty cell found.
	return [2]int{-1, -1}
}

// the below function checks if placing a number in a particular position is safe.
func isSafe(grid [][]int, row, col, num int) bool {
	return !usedInRow(grid, row, num) &&
		!usedInCol(grid, col, num) &&
		!usedInBox(grid, row-row%3, col-col%3, num)
}

// the below function  checks if the number is already used in the row.
func usedInRow(grid [][]int, row, num int) bool {
	for i := 0; i < gridSize; i++ {
		if grid[row][i] == num {
			return true
		}
	}
	return false
}

// the below function checks if the number is already used in the column.
func usedInCol(grid [][]int, col, num int) bool {
	for i := 0; i < gridSize; i++ {
		if grid[i][col] == num {
			return true
		}
	}
	return false
}

// The below function  checks if the number is already used in the 3x3 box.
func usedInBox(grid [][]int, startRow, startCol, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return true
			}
		}
	}
	return false
}

func main() {
	// Sudoku puzzle with 0 represents empty cells
	sudokuGrid := [][]int{
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

	fmt.Println("Original Sudoku Puzzle:")
	PrintGrid(sudokuGrid)

	if SolveSudoku(sudokuGrid) {
		fmt.Println("\nSolved Sudoku Puzzle:")
		PrintGrid(sudokuGrid)
	} else {
		fmt.Println("\nNo solution exists.")
	}
}
