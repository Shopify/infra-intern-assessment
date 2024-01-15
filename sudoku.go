package main

func SolveSudoku(grid [][]int) [][]int {
	solve(grid)
	return grid
}

func solve(grid [][]int) bool {
	empty := findEmptyCell(grid)
	if empty == nil {
		return true // Puzzle solved
	}

	row, col := empty[0], empty[1]

	for num := 1; num <= 9; num++ {
		if isValidMove(grid, row, col, num) {
			grid[row][col] = num

			if solve(grid) {
				return true // If puzzle is solved with the current placement, return true
			}

			// If the current placement does not lead to a solution, backtrack
			grid[row][col] = 0
		}
	}

	return false // No valid move found, trigger backtracking
}

func findEmptyCell(grid [][]int) []int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return []int{i, j}
			}
		}
	}
	return nil
}

func isValidMove(grid [][]int, row, col, num int) bool {
	// Check if the number is not present in the same row and column
	for i := 0; i < 9; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	// Check if the number is not present in the 3x3 sub-grid
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}
