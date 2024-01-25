package main

var SIZE int = 9

func SolveSudoku(sudoku [][]int) [][]int {
	if solve(sudoku, 0, 0) {
		return sudoku
	} else {
		return nil
	}
}

func solve(sudoku [][]int, row int, col int) bool {
	if row == SIZE-1 && col == SIZE {
		return true
	}
	if col == SIZE {
		row++
		col = 0
	}

	if sudoku[row][col] > 0 {
		return solve(sudoku, row, col+1)
	}

	for num := 1; num <= SIZE; num++ {
		if sudokuCheck(sudoku, row, col, num) {
			sudoku[row][col] = num
			if solve(sudoku, row, col+1) {
				return true
			}
		}
		sudoku[row][col] = 0
	}
	return false
}

func sudokuCheck(sudoku [][]int, row int, col int, num int) bool {
	for c := 0; c < SIZE; c++ {
		if sudoku[row][c] == num {
			return false
		}
	}
	for r := 0; r < SIZE; r++ {
		if sudoku[r][col] == num {
			return false
		}
	}

	for i := row - row%3; i < row+3-row%3; i++ {
		for j := col - col%3; j < col+3-col%3; j++ {
			if sudoku[i][j] == num {
				return false
			}
		}
	}
	return true
}
