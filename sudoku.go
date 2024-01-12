/*

	Sudoku Solver in Go
	Author: Shawn Zhang
	Date: January 11th 2024

	Description:
		This Go program provides a backtracking algorithm-based Sudoku solver.
		The 'SolveSudoku' function attempts to solve a partially filled 9x9 Sudoku puzzle 
		and returns the solved puzzle.

	Algorithm Overview:
		A recursive backtracking algorithm is used by the solution. The 'solve' function
		iterates through each cell in the puzzle, trying numbers 1 through 9 on empty cells 
		until a valid solution is discovered. The 'isValidMove' function checks for conflicts in 
		the same row, column, and 3x3 block of a specific cell.

*/

package main


func SolveSudoku(puzzle [][]int) [][]int {
	// call helper function 'solve' to start solving from top left corner
	if solve(puzzle, 0, 0) {
		// return puzzle if solved
		return puzzle
	}

	// return input if no solution is found
	return puzzle
}

// solve is the helper function that recursively tries to solve the Sudoku puzzle
func solve(grid [][]int, startRow, startCol int) bool {
	for r := startRow; r < 9; r++ {
		// reset our starting column to 0 for every new row
		startCol = 0
		for c := startCol; c < 9; c++ {

			// skip already filled cells
			if grid[r][c] != 0 {
				continue
			}

			for num := 1; num <= 9; num++ {
				if isValidMove(grid, r, c, num) {
					// if our current number is valid, place it on our grid
					grid[r][c] = num

					// recursively attempt to solve puzzle with updated grid
					if solve(grid, r, c+1) {
						return true
					}
					// if the current number did not work, backtrack and reset cell
					grid[r][c] = 0
				}
			}
			// no valid number is found
			return false
		}
	}
	// all cells are filled and no errors 
	return true
}

// isValidMove checks a number at a specific position in the grid is valid
func isValidMove(grid [][]int, row, col, num int) bool {
	// calculate the starting position of the 3x3 block containing the current cell
	blockRow := (row / 3) * 3
	blockCol := (col / 3) * 3

	// check for conflicts in same row and column
	for i := 0; i < 9; i++ {
		if grid[i][col] == num || grid[row][i] == num {
			return false
		}
	}

	// check for conflicts in the 3x3 block
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[blockRow+i][blockCol+j] == num {
				return false
			}
		}
	}

	return true
}

