package main

func isValid(grid [][]int, row int, col int, num int) bool {
	// Check if the number is already in the row
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	// Check if the number is already in the column
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}

	// Check if the number is already in the 3x3 box
	boxRow := row - row%3
	boxCol := col - col%3
	for i := boxRow; i < boxRow+3; i++ {
		for j := boxCol; j < boxCol+3; j++ {
			if grid[i][j] == num {
				return false
			}
		}
	}

	return true
}

func SolveSudoku(grid [][]int) [][]int {

	for i := range grid {
		for j, k := range grid[i] {
			// If the cell is empty, try to fill it with a number
			if k != 0 {
				continue
			}

			// Try all numbers from 1 to 9
			for n := 1; n <= 9; n++ {
				if isValid(grid, i, j, n) {
					grid[i][j] = n
					// If the number is valid, try to solve the rest of the grid
					if SolveSudoku(grid) != nil {
						return grid
					} else {
						// If the number is not valid, reset it and try the next one
						grid[i][j] = 0
					}
				} else {
					// If the number is not valid, reset it and try the next one
					grid[i][j] = 0
				}
			}

			// If no number is valid, return nil
			return nil
		}

	}

	return grid

}
