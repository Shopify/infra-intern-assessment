package main

func isValidCell(input [][]int, row, col, val int) bool {
	// check row
	for _, v := range input[row] {
		if v == val {
			return false
		}
	}

	// check column
	for _, v := range input {
		if v[col] == val {
			return false
		}
	}

	// check 3x3 subgrid
	start_row := (row / 3) * 3
	start_col := (col / 3) * 3
	for i := start_row; i < start_row+3; i++ {
		for j := start_col; j < start_col+3; j++ {
			if input[i][j] == val {
				return false
			}
		}
	}

	return true
}

func Solve(input [][]int, row, col int) [][]int {
	// base case
	if row == 9 {
		return input
	}

	// move to next row
	if col == 9 {
		return Solve(input, row+1, 0)
	}

	// skip if cell is not empty
	if input[row][col] != 0 {
		return Solve(input, row, col+1)
	}

	// try all possible values
	for i := 1; i < 10; i++ {
		if isValidCell(input, row, col, i) {
			input[row][col] = i
			if Solve(input, row, col+1) != nil {
				return input
			}

			input[row][col] = 0
		}

	}

	return nil
}

func SolveSudoku(input [][]int) [][]int {
	return Solve(input, 0, 0)
}

func main() {
}
