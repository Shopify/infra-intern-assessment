package main

// common keywords
const empty int = 0
const length int = 9

func IsValidRow(sudoku [][]int, row int, col int, value int) bool{
	// if intended value exists in row, return false
	for rowPos := 0; rowPos < length; rowPos++ {
		if sudoku[row][rowPos] == value {
			return false
		}
	}
	return true
}

func IsValidColumn(sudoku [][]int, row int, col int, value int) bool{
	// if intended value exists in column, return false
	for colPos := 0; colPos < length; colPos++ {
		if sudoku[colPos][col] == value {
			return false
		}
	}
	return true
}

func IsValidBox(sudoku [][]int, row int, col int, value int) bool{
	// if intended value exists in box, return false
	boxRow := row - row%3
	boxCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[i+boxRow][j+boxCol] == value {
				return false
			}
		}
	}
	return true
}

func findEmpty(sudoku [][]int) (int, int){
	// finds the next 0 in the sudoku board and returns positions
	// returns -1, -1 if no 0's found (board is solved)
	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			if sudoku[row][col] == empty {
				return row, col
			}
		}
	}
	return -1, -1
}

func SolveSudoku(sudoku [][]int) [][]int{
	// backtracking solution for sudoku
	// no checks for size and contents of sudoku arrays, assuming constraints are followed
	// if board is unsolvable, return null value

	emptyRow, emptyColumn := findEmpty(sudoku)

	if emptyRow == -1 && emptyColumn == -1 {
		// return true since sudoku is solved
		return sudoku
	}

	for num := 1; num <= length; num++ {
		// check if current number is a valid choice
		if IsValidRow(sudoku, emptyRow, emptyColumn, num) &&
		IsValidColumn(sudoku, emptyRow, emptyColumn, num) &&
		IsValidBox(sudoku, emptyRow, emptyColumn, num) {

			// valid choice
			sudoku[emptyRow][emptyColumn] = num

			// if backtracking leads to solved board
			if SolveSudoku(sudoku) != nil {
				return sudoku
			}

			// backtracking step, undo number filled in
			sudoku[emptyRow][emptyColumn] = empty
		}
	}

	// return null value if unsolved
	return nil
}