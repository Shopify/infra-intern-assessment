// File: helper.go
// Created Date: Jan 29 2024
// Author: Indranil Palit
// Description: This file contains the helper functions for the Sudoku solver.

package main

// This function checks if the given digit can be placed at the given row
func inRow(grid [][]int, row, num int) bool {
	for i := 0; i < N; i++ {
		if grid[row][i] == num {
			return true
		}
	}
	return false
}

// This function checks if the given digit can be placed at the given column
func inCol(grid [][]int, col, num int) bool {
	for i := 0; i < N; i++ {
		if grid[i][col] == num {
			return true
		}
	}
	return false
}

// This function checks if the given digit can be placed in the (3 x 3) box
func inBox(grid [][]int, row, col, num int) bool {
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return true
			}
		}
	}
	return false
}

// Helper function to check if a digit can be placed at a given position.
// This basically calls the above three functions to check if the digit can be placed at the given row, column and box
func isSafe(grid [][]int, row, col, num int) bool {
	return !inRow(grid, row, num) && !inCol(grid, col, num) && !inBox(grid, row, col, num)
}
