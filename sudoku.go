package main

// Function Name:   SolveSudoku
// Purpose:         The main function used to solve a given sudoku board
// Parameters:      board - 2D integer array representing the sudoku board
// Returns:         [][]int
func SolveSudoku(board [][]int) [][]int {
	// Get the dimensions of the board
	rowLength := len(board[0])
	colLength := len(board)

	// Find the "0" spots that need to be filled
	zeroesPosition := make([][]int, 0)
	for row := 0; row < rowLength; row++ {
		for col := 0; col < colLength; col++ {
			if board[row][col] == 0 {
				zeroesPosition = append(zeroesPosition, []int{row, col})
			}
		}
	}

	// Call the findSolution function to solve the Sudoku puzzle
	solution := findSolution(board, zeroesPosition, 0)

	return solution
}

// Function Name:   findSolution
// Purpose:         Recursively attempts to solve the board by finding the next number to fill
//                  Backtracks as needed if the filled number does not satisfy the game's restrictions
//
// Parameters:      board          - the current state of the sudoku board
//                  zeroesPosition - 2D integer array representing coordinates (row, column) of the empty spots to be filled
//                  position       - the current position
//
// Returns:         [][]int
func findSolution(board [][]int, zeroesPosition [][]int, position int) [][]int {
	// Base case: return if all the zeroes spots are filled
	if position == len(zeroesPosition) {
		return board
	}

	// Get the indices/position of the current spot to be filled
	row := zeroesPosition[position][0]
	col := zeroesPosition[position][1]

	// Fill the current spot on the sudoku grid via cycling through 1-9
	for num := 1; num <= 9; num++ {
		if checkValidity(board, row, col, num) {
			board[row][col] = num

			// Recursively solve the puzzle with the updated board
			result := findSolution(board, zeroesPosition, position + 1)
			if result != nil {
				return result
			}

			// Backtrack by resetting current spot if no solution is found
			board[row][col] = 0
		}
	}

	return nil
}

// Function name:   isInSquare
// Purpose:         Checks if given number is present in the 3x3 sub-grid of the sudoku board
//
// Parameters:      board - the current sudoku board
//                  row   - the row index of the current cell
//                  col   - the column index of the current cell
//                  num   - the number to check for
//
// Returns:         boolean
func isInSquare(board [][]int, row, col, num int) bool {
	topRow := (row / 3) * 3
	leftCol := (col / 3) * 3

	for i := 0; i < 9; i++ {
		if board[topRow + i / 3][leftCol + i % 3] == num {
			return true
		}
	}
	return false
}

// Function name:   checkValidity
// Purpose:         Checks the validity of placing the number in a specific spot against the current row, current column, and current 3x3 subgrid
//
// Parameters:      board - the current sudoku board
//                  row   - the row index of the current cell
//                  col   - the column index of the current cell
//                  num   - the number to check for
//
// Returns:         boolean
func checkValidity(board [][]int, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num || isInSquare(board, row, col, num) {
			return false
		}
	}
	return true
}
