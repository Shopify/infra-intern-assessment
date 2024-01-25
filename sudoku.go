package main

// SolveSudoku solves the provided sudoku puzzle and returns the
// solved version. If the provided puzzle does not have a valid
// solution or is, return nil
func SolveSudoku(sudoku [][]int) [][]int {

	if !checkValidSudoku(sudoku) {
		return nil
	}

	table := deepCopy(sudoku)

	// solves each sudoku cell left to right, top to bottom
	var solveCell func(r, c int) bool
	solveCell = func(r, c int) bool {
		if r == 9 { // r=9 only after solving bottom right grid
			return true
		} else if c == 9 {
			return solveCell(r+1, 0)
		} else if table[r][c] != 0 {
			return solveCell(r, c+1)
		} else {
			for i := 1; i < 10; i++ {
				if isValid(table, r, c, i) {
					table[r][c] = i
					if solveCell(r, c+1) {
						return true
					}
					table[r][c] = 0
				}
			}
		}
		return false
	}

	if !solveCell(0, 0) {
		return nil
	}
	return table
}

// checkValidSudoku returns true if sudoku is a 9x9 grid
func checkValidSudoku(sudoku [][]int) bool {
	// check number of rows
	if len(sudoku) != 9 {
		return false
	}

	// check number of columns
	for _, row := range sudoku {
		if len(row) != 9 {
			return false
		}
	}
	return true
}

// deepCopy returns a deep copy of matrix
func deepCopy(matrix [][]int) [][]int {
	deepCopy := make([][]int, len(matrix))
	for i := range deepCopy {
		deepCopy[i] = make([]int, len(matrix[i]))
		copy(deepCopy[i], matrix[i])
	}

	return deepCopy
}

// isValid checks if num can be placed on table[r][c] without
// violating sudoku rules. Returns true if possible and false
// otherwise
func isValid(table [][]int, r int, c int, num int) bool {
	// check row
	for _, elem := range table[r] {
		if elem == num {
			return false
		}
	}

	// check column
	for i := 0; i < len(table); i++ {
		if table[i][c] == num {
			return false
		}
	}

	// check grids
	gridCol := (c / 3) * 3
	gridRow := (r / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if table[i+gridRow][j+gridCol] == num {
				return false
			}
		}
	}
	return true
}
