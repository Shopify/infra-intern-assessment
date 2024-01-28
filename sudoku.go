package main

// dimension of the sudoku puzzle
const N = 9

// solves a sudoku puzzle and returns the solved puzzle
func SolveSudoku(sudoku [][]int) [][]int {
	// check if sudoku is valid
	if !IsValidSudoku(sudoku) {
		return nil
	}

	// create a deep copy of sudoku
	sudokuCopy := DeepCopy(sudoku)

	// start solving from top left
	if !solve(0, 0, sudokuCopy) {
		return nil
	}

	return sudokuCopy
}

// helper function that solves a sudoku puzzle and returns true if solved and false otherwise
func solve(r int, c int, sudoku [][]int) bool {
	// check if we are done
	if r == 9 {
		return true
	}

	// check if we need to move to next row
	if c == 9 {
		return solve(r+1, 0, sudoku)
	}

	// check if we need to solve this cell
	if sudoku[r][c] != 0 {
		return solve(r, c+1, sudoku)
	}

	// try all possible numbers
	for i := 1; i <= 9; i++ {
		if IsValidCell(sudoku, r, c, i) {
			sudoku[r][c] = i
			if solve(r, c+1, sudoku) {
				return true
			}
			sudoku[r][c] = 0 // backtrack
		}
	}

	return false
}

// returns true if sudoku is valid (constraints don't check for out of bound integers)
func IsValidSudoku(sudoku [][]int) bool {
	// check valid numbers
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if sudoku[r][c] < 0 || sudoku[r][c] > 9 {
				return false
			}
		}
	}

	return true
}

// returns a deep copy of the sudoku
func DeepCopy(sudoku [][]int) [][]int {
	sudokuCopy := make([][]int, N)
	for i := range sudokuCopy {
		sudokuCopy[i] = make([]int, N)
		// copy row
		copy(sudokuCopy[i], sudoku[i])
	}
	return sudokuCopy
}

// returns true if sudoku[r][c] can hold num
func IsValidCell(sudoku [][]int, r int, c int, num int) bool {
	// check row and column
	for i := 0; i < N; i++ {
		if sudoku[r][i] == num || sudoku[i][c] == num {
			return false
		}
	}

	// check 3x3 grids
	startRow := (r / 3) * 3
	startCol := (c / 3) * 3
	for dr := 0; dr < 3; dr++ {
		for dc := 0; dc < 3; dc++ {
			if sudoku[startRow+dr][startCol+dc] == num {
				return false
			}
		}
	}

	return true
}
