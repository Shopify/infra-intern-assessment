package main

// Define N as the size of the sudoku grid
const N = 9 

// isInCol checks if 'num' is present in the current 'col'
func isInCol(grid [][]int, col, num int) bool {
	for row := 0; row < N; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// isInRow checks if 'num' is present in the current 'row'
func isInRow(grid [][]int, row, num int) bool {
	for col := 0; col < N; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// isInSubgrid checks if 'num' is present in the 3x3 subgrid
func isInSubgrid(grid [][]int, subgridStartRow, subgridStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[row + subgridStartRow][col + subgridStartCol] == num {
				return true
			}
		}
	}
	return false
}

// isValid checks if 'num' is valid for the given cell (validate placement)
func isValid(grid [][]int, row int, col int, num int) bool {
	// Checks for presence of 'num' in current row, column, or subgrid
	return !(isInRow(grid, row, num) || isInCol(grid, col, num) || isInSubgrid(grid, row-row%3, col-col%3, num))
}


// solveSudokuHelper, a recursive helper function, uses backtracking to solve the puzzle
func solveSudokuHelper(grid [][]int, row int, col int) bool {
	// End of grid is reached (base case)
	if row == N-1 && col == N {
		return true
	}
	// At the end of the column, the function moves to the next row
	if col == N {
		row++
		col = 0
	}
	// Filled/occupied cells are skipped, 'num' cannot be placed here
	if grid[row][col] != 0 {
		return solveSudokuHelper(grid, row, col+1)
	}
	// Placing number 1 through N, backtrack if necessary 
	for i := 1; i <= N; i++ {
		if isValid(grid, row, col, i) {
			grid[row][col] = i

			if solveSudokuHelper(grid, row, col+1) {
				return true
			}
	
			grid[row][col] = 0
		}
	}
	return false
}

// SolveSudoku solves the Sudoku puzzle
func SolveSudoku(grid [][]int) [][]int {
	if solveSudokuHelper(grid, 0, 0) {
		return grid 
	}
	return nil
}
