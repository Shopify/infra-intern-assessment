package main

import "fmt"

// width and height of the square sudoku board
const BoardSize int = 9

// struct representing coordinates of a space on the sudoku board
// pointers are used so that we can indicates "no space exists" with nil fields
type coords struct {
	rowIndex *int
	colIndex *int
}

// solves the given sudoku board
func SolveSudoku(board [][]int) [][]int {

	// solve the sudoku board
	Solve(board)

	// print the solved board so that it is structured nicely in the terminal
	fmt.Println("Solved Sodoku:")
	for row := 0; row < BoardSize; row++ {
		for column := 0; column < BoardSize; column++ {
			fmt.Print(board[row][column], "  ")
		}
		fmt.Println()
	}

	// return the solved board
	return board
}

// recursive sudoku solving function, returns true if board is solveable, otherwise false
func Solve(board [][]int) bool {

	// find the next available empty space
	emptySpace := FindEmptySpace(board)

	// if there are no empty spaces, the board must be solved (because every input board has EXACTLY ONE solution)
	if emptySpace.rowIndex == nil && emptySpace.colIndex == nil {
		return true
	}

	row, col := *emptySpace.rowIndex, *emptySpace.colIndex

	// try digits 1-9 until we find a valid one
	for digit := 0; digit <= 9; digit++ {
		if VerifyDigitFitsInSpace(board, row, col, digit) {
			board[row][col] = digit

			if Solve(board) {
				return true
			}

			board[row][col] = 0
		}
	}

	// board is unsolveable so return false and try different digits
	return false
}

// finds the next empty space if it exists and returns its coordinates, otherwise returns nil for both coords
func FindEmptySpace(board [][]int) coords {

	// go through every space in the board until we find an empty one
	for row := 0; row < BoardSize; row++ {
		for column := 0; column < BoardSize; column++ {

			// check if the current space is empty and return its coords if it is
			if board[row][column] == 0 {
				return coords{rowIndex: &row, colIndex: &column}
			}
		}
	}

	// there are no empty spaces in the board
	return coords{rowIndex: nil, colIndex: nil}
}

// determines whether the given digit is valid in a specific space on the board, returns true if it is, otherwise false
func VerifyDigitFitsInSpace(board [][]int, rowIndex int, colIndex int, digit int) bool {

	// check if the digit can fit in the given row
	for i := 0; i < BoardSize; i++ {
		if board[rowIndex][i] == digit {
			return false
		}
	}

	// check if the digit can fit in the given column
	for i := 0; i < BoardSize; i++ {
		if board[i][colIndex] == digit {
			return false
		}
	}

	// check if the digit can fit in the given box
	startRow, startCol := rowIndex-(rowIndex%3), colIndex-(colIndex%3)
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if board[r][c] == digit {
				return false
			}
		}
	}

	// if the digit is valid in the given row, column and box, return true
	return true
}
