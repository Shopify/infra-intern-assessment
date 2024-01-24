package main

// SolveSudoku takes 9x9 grid of integers
// 0s indicate empty cells
// 1-9s representing non empty cels
func SolveSudoku(grid [][]int) [][]int {
	Solve(grid)
	return grid
}

// Helper function
// Solve seeks to find empty cells and uses backtracking The function returns
// true if the puzzle is solved and false if no valid solution exists.
func Solve(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// search for empty cell
			if grid[i][j] == 0 {
				for k := 1; k <= 9; k++ {
					if IsValid(grid, i, j, k) {
						grid[i][j] = k
						if Solve(grid) {
							return true
						} else {
							grid[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

// Helper function
// IsValid checks if placing the num in grid[row][col] is valid
func IsValid(grid [][]int, row, col, num int) bool {
	// check if the row's valid
	for x := 0; x < 9; x++ {
		if grid[row][x] == num {
			return false
		}
	}

	// check if col is valid
	for x := 0; x < 9; x++ {
		if grid[x][col] == num {
			return false
		}
	}

	// check if the 3x3 square num is contained in is valid
	startRow := row - row%3
	startCol := col - col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if grid[i][j] == num {
				return false
			}
		}
	}
	return true
}
