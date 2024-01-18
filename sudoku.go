package main

// Backtracking algorithm
// Find what numbers I can place at a given square
// Place each candidate number
// If there is an empty cell that I cannot place a number in -> backtrack

// Finds the next cell that needs to be solved
func findNextCell(input [][]int, prevCol int, prevRow int) (col int, row int) {
	for col := prevCol; col < len(input); col++ {
		for row := prevRow; row < len(input); row++ {
			if input[row][col] == 0 {
				return col, row
			}
		}
	}
	return -1, -1
}

// Finds the set of possible placable numbers for a given cell
func findPossibleNumbers(input [][]int, col int, row int) (possibleNumbers map[int]struct{}) {
	// initialize a set of [0,9]
	possibleNumbers = map[int]struct{}{}
	for i := 1; i <= 9; i++ {
		possibleNumbers[i] = struct{}{}
	}

	// remove numbers in rows and columns
	for i := 0; i < 9; i++ {
		delete(possibleNumbers, input[row][i])
		delete(possibleNumbers, input[i][col])
	}

	// remove numbers in a subgrid (3x3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			delete(possibleNumbers, input[row/3*3+i][col/3*3+j])
		}
	}
	return possibleNumbers
}

func backtrack(input [][]int, col int, row int) (solved bool) {
	col, row = findNextCell(input[:], col, row)
	if col == -1 && row == -1 { // if the next cell is unreachable,
		return true //  we have found a solution
	}

	// for every possible number
	for possibleNumber := range findPossibleNumbers(input[:], col, row) {
		input[row][col] = possibleNumber // place
		if backtrack(input[:], 0, 0) {   // try to solve the rest of the board
			return true
		}
		input[row][col] = 0 // remove, and retry
	}
	return false // no solution
}

func SolveSudoku(input [][]int) (output [][]int) {
	backtrack(input[:], 0, 0) // = true if solved, else false
	return input
}
