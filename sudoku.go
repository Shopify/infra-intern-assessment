package main

const gridSize = 9

// run dfs on grid to check all the possible valid states of sudoku and backtrack if necessary
func solveSudokuHelper(grid [][]int, row, col int) bool {
	// reached end of the grid
	if row == gridSize-1 && col == gridSize-1 {
		return true
	}

	// move to the next row if we've reached the end of a column
	if col == gridSize {
		row++
		col = 0
	}

	// skip filled cells and go to the next cell
	if grid[row][col] != 0 {
		return solveSudokuHelper(grid, row, col+1)
	}

	for num := 1; num <= gridSize; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num

			if solveSudokuHelper(grid, row, col+1) {
				return true
			}

			grid[row][col] = 0 // backtrack
		}
	}

	return false
}

func isSafe(grid [][]int, row, col, num int) bool {
	// check if the number is not present in the current row and column
	for i := 0; i < gridSize; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	// check if the number is not present in the 3x3 subgrid
	// using integer division to adjust offset the start point in the grid
	sr, sc := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[sr+i][sc+j] == num {
				return false
			}
		}
	}

	return true
}

func SolveSudoku(grid [][]int) [][]int {
	if solveSudokuHelper(grid, 0, 0) {
		return grid
	}
	// return nil if unsolvable
	return nil
}
