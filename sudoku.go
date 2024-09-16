package main

const N = 9 // N represents the size of the Sudoku grid (9x9)

// isSafe checks if a number can be placed in the given cell.
// Sudoku rules: the number must not be present in the same row, column, or 3x3 subgrid.
func isSafe(grid [][]int, row, col, num int) bool {
	for i := 0; i < N; i++ {
		// Check for the number in the row and column
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}

		// Check for the number in the 3x3 subgrid
		subgridRow := 3*(row/3) + i/3
		subgridCol := 3*(col/3) + i%3
		if grid[subgridRow][subgridCol] == num {
			return false
		}
	}
	return true
}

// SolveSudoku solves a given Sudoku puzzle.
func SolveSudoku(grid [][]int) [][]int {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if grid[row][col] == 0 {
				for num := 1; num <= N; num++ {
					if isSafe(grid, row, col, num) {
						grid[row][col] = num

						// Recursively try to fill the rest of the grid
						if SolveSudoku(grid) != nil {
							return grid
						}

						// Backtrack if no number leads to a solution
						grid[row][col] = 0
					}
				}
				return nil // Trigger backtracking
			}
		}
	}
	return grid // Return the solved puzzle
}
