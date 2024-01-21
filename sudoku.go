package main

import ( "fmt" )

const SUDOKU_BOARD_SIZE int = 9
const EMPTY_CELL int = 0
const MAX_DIGIT int = 9
const MIN_DIGIT int = 1

// Default value sets to false
var rowSet [9][10]bool
var colSet [9][10]bool
var squareSet [3][3][10]bool

func printSudokuGrid(sudokuGrid *[][]int) {
	fmt.Printf("SOLVED BOARD\n[\n")

	for row := 0; row < SUDOKU_BOARD_SIZE; row++ {

		fmt.Printf("\t[")

		for col := 0; col < SUDOKU_BOARD_SIZE; col++ {
			fmt.Printf("%d", (*sudokuGrid)[row][col] )

			if col < SUDOKU_BOARD_SIZE - 1 { fmt.Printf(", ") }
		} // for

		fmt.Printf("],\n")
	} // for

	fmt.Printf("]\n")
} // printSudokuGrid

func isValidSudokuMove(row int, col int, digit int) bool {
	return !rowSet[row][digit] && !colSet[col][digit] && !squareSet[row/3][col/3][digit]
} // isValidSudokuMove

func SolveSudokuDepthFirstSearch(sudokuGrid *[][]int, row int, col int) bool {
	if row == SUDOKU_BOARD_SIZE {
		return true
	}
	if col == SUDOKU_BOARD_SIZE {
		return SolveSudokuDepthFirstSearch(sudokuGrid, row+1, 0)
	}
	if (*sudokuGrid)[row][col] != EMPTY_CELL {
		return SolveSudokuDepthFirstSearch(sudokuGrid, row, col+1)
	}

	for digit := MIN_DIGIT; digit <= MAX_DIGIT; digit++ {

		if isValidSudokuMove(row, col, digit) {

			// Set the grid and place this choice in our row, col, and square sets so we know to not re-pick this value in future cells
			(*sudokuGrid)[row][col] = digit
			rowSet[row][digit] = true
			colSet[col][digit] = true
			squareSet[row/3][col/3][digit] = true

			/* Follow this depth first search, with the decision that sudokuGrid[row][col] = i.
			If we reach the end of the board and the choices made were all valid, then we found the solution path! */
			if SolveSudokuDepthFirstSearch(sudokuGrid, row, col+1) {
				return true
			}

			// If sudokuGrid[row][col] = i does not result in a valid solution then undo the choice and continue
			(*sudokuGrid)[row][col] = EMPTY_CELL
			rowSet[row][digit] = false
			colSet[col][digit] = false
			squareSet[row/3][col/3][digit] = false

		} // if
	} // for

	return false
} // SolveSudokuDepthFirstSearch

func SolveSudoku(sudokuGrid [][]int) [][]int {
	/* 
	The main idea is to traverse the entire grid and at each empty cell, we try every possible
	valid digit using a depth first search algorithm. As soon as we discover a valid path, we return as that
	is our unique answer as guaranteed by the problem statement.

	Create a rowSet, colSet, and squareSet where:
	rowSet[r][digit] == true iff sudokuGrid at row r contains digit
	colSet[c][digit] == true iff sudokuGrid at column c contains digit
	squareSet[r/3][c/3][digit] == true iff sudokuGrid within the square containing cell (r, c) contains digit 
	
	We use these sets to check what digits are valid in any cell of the sudokuGrid efficiently.
	*/

	for row := 0; row < SUDOKU_BOARD_SIZE; row++ {
		for col := 0; col < SUDOKU_BOARD_SIZE; col++ {

			if sudokuGrid[row][col] != EMPTY_CELL {

				var DIGIT int = sudokuGrid[row][col]

				rowSet[row][DIGIT] = true
				colSet[col][DIGIT] = true
				squareSet[row/3][col/3][DIGIT] = true
			} // if
		} // for
	} // for

	SolveSudokuDepthFirstSearch(&sudokuGrid, 0, 0)

	printSudokuGrid(&sudokuGrid)

	return sudokuGrid
} // SolveSudoku
