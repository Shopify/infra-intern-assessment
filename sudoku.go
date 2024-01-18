package main

// SolveSudoku solves a provide sudoku grid
// Precondition: given grid has is valid and has only one solution
// Input grid([][] int) of size 9 * 9 is guaranteed having only one solution
// Returns solution in a grid([][] int) of size 9 * 9
func SolveSudoku(input [][]int) [][]int {
	backtracking(&input, 0, 0)
	return input
}

func backtracking(grid *[][]int, row int, col int) bool {
	rowMax, colMax := len(*grid), len((*grid)[0])

	// Base case: if we have proceeded all rows, sudoku solved. Return true
	if row == rowMax {
		return true
	}

	// If the end of a column is reached, go the start of next row
	if col == colMax {
		return backtracking(grid, row+1, 0)
	}

	// if it's a non zero cell, proceed to next cell
	if (*grid)[row][col] != 0 {
		return backtracking(grid, row, col+1)
	}

	// Try all values (1 to 9) for current cell
	for value := 1; value <= 9; value++ {
		// If value is not valid for current cell, try next value
		if !validate(value, grid, row, col) {
			continue
		}

		// Assign validated value
		(*grid)[row][col] = value

		// If true, sudoku solution is found, return true
		if backtracking(grid, row, col+1) {
			return true
		}

		// Otherwise, this value is not correct proceed backtrack
		(*grid)[row][col] = 0
	}

	// No solution is found in this call and its recursive calls, return false and backtrack
	return false
}

// validate checks if the given value is a valid input in given grid
// row, col are the coordinates of the value in the grid
// Returns true if value is valid, false otherwise
func validate(value int, grid *[][]int, row int, col int) bool {
	// Check if value is duplicated in current row and column
	for i := 0; i < 9; i++ {
		if (*grid)[i][col] == value || (*grid)[row][i] == value {
			return false
		}
	}

	// blockConv is the conversion factor that helps to find which 3*3 block the value is located
	/* grid[blockRow][blockCol] represent the actual indices of the top left cell of a 3*3 block
	shown as below:
		[X][ ][ ]
		[ ][ ][ ]
		[ ][ ][ ]
	*/
	const blockConv int = 3
	blockRow := row - row%blockConv // Actual row index
	blockCol := col - col%blockConv // Actual column index

	// Check if value is duplicated in the 3*3 block
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*grid)[blockRow+i][blockCol+j] == value {
				return false
			}
		}
	}

	// If value is not duplicated, this value is valid
	return true
}
