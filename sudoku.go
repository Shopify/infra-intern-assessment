package main
import (
	"fmt"
)

func SolveSudoku(grid [][]int) [][]int {
	if solve(&grid) {
		return grid
	}
	return nil
}

func solve(grid *[][]int) bool {
	row, col := findUnassignedLocation(grid)
	if row == -1 && col == -1 {
		return true // All cells are assigned
	}

	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			(*grid)[row][col] = num // Assign the number

			// If the assignment is successful, recursively solve
			if solve(grid) {
				return true
			}

			// If the assigned number does not lead to a solution, reset it
			(*grid)[row][col] = 0
		}
	}
	return false // No valid number found
}

func findUnassignedLocation(grid *[][]int) (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if (*grid)[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1 // If no unassigned location is found
}

func isSafe(grid *[][]int, row, col, num int) bool {
	return !usedInRow(grid, row, num) &&
		!usedInCol(grid, col, num) &&
		!usedInBox(grid, row-row%3, col-col%3, num)
}

func usedInRow(grid *[][]int, row, num int) bool {
	for col := 0; col < 9; col++ {
		if (*grid)[row][col] == num {
			return true
		}
	}
	return false
}

func usedInCol(grid *[][]int, col, num int) bool {
	for row := 0; row < 9; row++ {
		if (*grid)[row][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(grid *[][]int, boxStartRow, boxStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if (*grid)[row+boxStartRow][col+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}
