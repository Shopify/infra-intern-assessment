package main

/*
Written by Zaheer MacDonald, updated as recently as January 26, 2024, for the purposes of
the Shopify Infra Intern Assessment.
*/
func SolveSudoku(sudoku [][]int) [][]int {
	/*
		Using backtracking, this function takes a sudoku board as an argument, and then
		with the help of other verification functions such as isComplete, mutates the board
		in-place and returns true once the sudoku is solved.
	*/

	ExecuteSolution(sudoku)
	return sudoku

}

func ExecuteSolution(sudoku [][]int) bool {
	var col, row int

	if IsComplete(sudoku, &row, &col) {
		return true // function terminates if sudoku is complete.
	}
	for potential := 1; potential < 10; potential++ {
		if IsValidSudoku(sudoku, potential, row, col) {
			sudoku[row][col] = potential // tentatively stores a grid value to check if it is valid.
			if ExecuteSolution(sudoku) {
				return true
			}
			sudoku[row][col] = 0 // restores grid position to zero if no valid sudoku exists with previous potential value.
		}
	}

	return false
}

func IsComplete(sudoku [][]int, row *int, col *int) bool {
	/*
		This function checks to see if the sudoku board still has any remaining spaces.
		Passes row and column positions by reference rather than value in order to improve efficiency.
	*/
	for *row = 0; *row < 9; *row++ {
		for *col = 0; *col < 9; *col++ {
			if sudoku[*row][*col] == 0 { // checks if there is a position in the board that is empty.
				return false
			}
		}
	}
	return true
}

func IsValidSudoku(sudoku [][]int, potential int, pRow int, pCol int) bool {
	/*
		This function takes in:
		1) a sudoku
		2) a potential value to be added to the board
		3) the coordinates of the board for which the potential value may be added
		and returns true if there are no duplicates in any of the squares, rows, or columns
		when the potential value to be added is taken into account. If duplicates are
		present, it returns false.
	*/
	return ValidCol(sudoku, pCol, potential) && ValidRow(sudoku,
		pRow, potential) && (ValidSquare(sudoku, pCol-pCol%3, pRow-pRow%3, potential))

}

func ValidRow(sudoku [][]int, pRow int, potential int) bool {
	/*
		This function takes as argument a sudoku grid, possible row number, and possible value to be inserted in the grid.
		It returns true if the row does not yet contain the possible value, and otherwise returns false.
	*/
	for col := 0; col < 9; col++ {
		if sudoku[pRow][col] == potential {
			return false
		}
	}
	return true
}

func ValidCol(sudoku [][]int, pCol int, potential int) bool {
	/*
		This function takes as argument a sudoku grid, possible column number, and possible value to be inserted in the grid.
		It returns true if the column does not yet contain the possible value, and otherwise returns false.
	*/
	for row := 0; row < 9; row++ {
		if sudoku[row][pCol] == potential {
			return false
		}
	}
	return true
}

func ValidSquare(sudoku [][]int, pColSquare int, pRowSquare int, potential int) bool {
	/*
		This function receives as argument a sudoku grid, potential value to insert into the grid,
		and row & column positions where that value may go. If the 3x3 square the potential value would
		be in already contains that value, the function returns false. Otherwise, it returns true.
	*/
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if sudoku[pRowSquare+row][pColSquare+col] == potential {
				return false
			}
		}
	}
	return true
}
