package main

const N = 9
const M = 3

// obtains the coordinates for the top-left coordinate of the corresponding row, col
func getSubgridTopLeft(row, col int) (int, int) {
	subgridRow := row - row%3
	subgridCol := col - col%3
	return subgridRow, subgridCol
}

// checks if num inserted into puzzle [row][col] would make a valid sudoku puzzle
func isValid(puzzle [][]int, row, col, num int) bool {
	// if num appears in the current row or col, then it is invalid
	for i := 0; i < N; i++ {
		if puzzle[row][i] == num || puzzle[i][col] == num {
			return false
		}
	}

	// if num appears in the current subgrid, then it is invalid
	startRow, startCol := getSubgridTopLeft(row, col)

	for i := startRow; i < startRow+M; i++ {
		for j := startCol; j < startCol+M; j++ {
			if puzzle[i][j] == num {
				return false
			}
		}
	}

	return true
}

func solve(puzzle [][]int) bool {
	// search for the first empty cell
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if puzzle[i][j] == 0 {

				// now try numbers 1 to 9 to place in the empty cell
				for k := 1; k <= 9; k++ {

					if isValid(puzzle, i, j, k) {
						puzzle[i][j] = k

						if solve(puzzle) { // if number is valid in the spot, attempt to solve the rest
							return true
						}

						// if the placement of num doesn't lead to a solution, backtrack
						puzzle[i][j] = 0
					}
				}

				// Break out of the loop if a solution is not found for the current empty cell
				return false
			}
		}
	}

	// if not found, then we have exhausted all cells
	return true
}

func SolveSudoku(initialPuzzle [][]int) [][]int {
	var solvedPuzzle [][]int

	// Copy elements from the slice to a new slice
	for i := 0; i < N; i++ {
		row := make([]int, N)
		copy(row, initialPuzzle[i])
		solvedPuzzle = append(solvedPuzzle, row)
	}

	if solve(solvedPuzzle) {
		return solvedPuzzle
	}

	// If no solution is found, return the initial puzzle as is
	return initialPuzzle
}