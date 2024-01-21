package main

func SolveSudoku(input [][]int) [][]int {
	// Make a copy of the input 2d array so we can pass by reference
	grid := make([][]int, 9)
	for i := range input {
		grid[i] = make([]int, 9)
		copy(grid, input)
	}

	// Call the recursive sudoku solver and return the grid
	recursiveSolver(grid, 0, 0)
	return grid
}

// We will use backtracking to solve this sudoku, at each cell that needs to be filled in we will test all possible values
func recursiveSolver(grid [][]int, row int, col int) bool {
	if row == 9 {
		// If there are no more rows to fill in we are done
		return true
	} else if col == 9 {
		// If we have gone past the last column of the row, move to the next row
		return recursiveSolver(grid, row+1, 0)
	} else if grid[row][col] != 0 {
		// If the grid slot is not 0 we do not have to fill it in, move to the next column
		return recursiveSolver(grid, row, col+1)
	} else {
		// If the grid slot is 0 we will test all values from 1 to 9 until the recursive function returns true (meaning a complete solution was found)
		for i := 1; i <= 9; i++ {
			// First check if it would be a valid cell (i is not used in the box, row, or column)
			if validCell(grid, row, col, i) {
				grid[row][col] = i
				// Now check if there is a complete solution
				if recursiveSolver(grid, row, col+1) {
					return true
				}
				// If there is not a complete solution, set the grid slot back to 0
				grid[row][col] = 0
			}
		}
	}

	return false
}

// This function checks to see if num is already used in the 3x3 box, row, or column
func validCell(grid [][]int, row int, col int, num int) bool {
	// Check the box
	for i := row - row%3; i < row-row%3+3; i++ {
		for j := col - col%3; j < col-col%3+3; j++ {
			if grid[i][j] == num {
				return false
			}
		}
	}

	// Check the row
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	// Check the column
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}

	return true
}
