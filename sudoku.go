package main

const N = 9 // N represents the size of the Sudoku grid (9x9)

// usedInRow checks if 'num' is present in the specified 'row'.
func usedInRow(grid [][]int, row, num int) bool {
	for col := 0; col < N; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// usedInCol checks if 'num' is present in the specified 'col'.
func usedInCol(grid [][]int, col, num int) bool {
	for row := 0; row < N; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// usedInSubgrid checks if 'num' is present in the 3x3 subgrid starting at (subgridStartRow, subgridStartCol).
func usedInSubgrid(grid [][]int, subgridStartRow, subgridStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[row + subgridStartRow][col + subgridStartCol] == num {
				return true
			}
		}
	}
	return false
}

// isSafe checks if a number can be placed in the given cell.
// Sudoku rules: the number must not be present in the same row, column, or 3x3 subgrid.
func isSafe(grid [][]int, row int, col int, num int) bool {
	return !(usedInRow(grid, row, num) || usedInCol(grid, col, num) || usedInSubgrid(grid, row-row%3, col-col%3, num))
}

// solveHelper uses backtracking to solve grid.
// It recursively fills cells with numbers and backtracks when it encounters an unsolvable state.
func solveHelper(grid [][]int, row int, col int) bool {
	// Base case: If we've reached the end of the grid, the puzzle is solved
	if row == N-1 && col == N {
		return true
	}

	// Move to the next row if we've reached the end of a column
	if col == N {
		row++
		col = 0
	}
	
	// Skip filled cells and go to the next cell
	if grid[row][col] != 0 {
		return solveHelper(grid, row, col+1)
	}

	// Try all possible numbers for this empty cell
	for i := 1; i <= N; i++ {
		// Checks if number is valid
		if isSafe(grid, row, col, i) {
			grid[row][col] = i

			// Recursively try to fill the rest of the grid
			if solveHelper(grid, row, col+1) {
				return true
			}
	
			// Backtrack if no number leads to a solution
			grid[row][col] = 0
		}
	}

	// Trigger backtracking
	return false
}

// SolveSudoku solves a given Sudoku puzzle
func SolveSudoku(grid [][]int) [][]int {
	if solveHelper(grid, 0, 0) {
		return grid // Return the solved puzzle
	}
	return nil // Return nil if unsolvable (but since constraints state grid will always have solution, should never happen)
}