package main


const N = 9

// This function checks if the number inputted is valid or not within the whole sudoku grid
func valid(input [][]int, row, col, num int) bool {

	// This makes sure that there is no same number in the same row, else returns false
	for x := 0; x < N; x++ {
		if input[row][x] == num {
			return false
		}
	}

	// This makes sure that there is no same number in the same column, else returns false
	for x := 0; x < N; x++ {
		if input[x][col] == num {
			return false
		}
	}
	// This makes sure that there is no same number in the 3x3 section that the number is suppose to be in, else returns false
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if input[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	// Made it to the end so this inputted number is legal
	return true
}


// This function assigns values to the sudoku grid and sees if it is correct
// Takes in any grid input and the row/column currently on
func solveSudokuHelper(input [][]int, row, col int) bool {

	// Reached the end, returns true
	if row == N-1 && col == N {
		return true
	}

	// Reached the end of column, then move on to next row
	if col == N {
		row++
		col = 0
	}

	// If current index is filled already, move onto next column
	if input[row][col] > 0 {
		return solveSudokuHelper(input, row, col+1)
	}

	// For each number  1-9
	for num := 1; num <= N; num++ {
		// Determine if the number is going to be valid when placed
		if valid(input, row, col, num) {
			// If so, place it
			input[row][col] = num
			// Move on to next column
			if solveSudokuHelper(input, row, col+1) {
				return true
			}
			// The guess/placement was wrong so we remove number
			input[row][col] = 0
		}
	}
	return false
}

// This function is just the one that the sudoku_test calls, as I did not want to change my function parameters or the parameters of the test function
func SolveSudoku(input [][]int) [][]int{
	if solveSudokuHelper(input, 0, 0){
		return input
	}
	return nil
}

