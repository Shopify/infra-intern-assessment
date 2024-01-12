package main

import (
	"errors"
)

// Defined per constraints of the challenge
const (
	boardSize = 9
	gridSize  = 3
)

// SolveSudoku solves a Sudoku puzzle represented by a 9x9 matrix.
//
// Args:
//
//	inputBoard: A 2D slice representing the Sudoku puzzle.
//
// Returns:
//
//	[][]int: The solved Sudoku board, or nil if the puzzle is unsolvable or the input is invalid.
func SolveSudoku(inputBoard [][]int) [][]int {

	if err := validateInput(inputBoard); err != nil {
		return nil
	}

	// Deepcopy to avoid modifying argument
	outputBoard := make([][]int, boardSize)
	for i := range inputBoard {
		outputBoard[i] = make([]int, boardSize)
		copy(outputBoard[i], inputBoard[i])
	}

	if solveSudokuRecursive(outputBoard) {
		return outputBoard
	} else {
		return nil
	}
}

// validateInput checks if the provided Sudoku board is a valid 9x9 matrix with values
// in the range [0, 9].
//
// Args:
//
//	board: A 2D slice representing the Sudoku puzzle.
//
// Returns:
//
//	error: An error if the board is invalid, or nil if the input is valid.
func validateInput(board [][]int) error {

	if len(board) != boardSize {
		return errors.New("invalid board: not a 9*9 matrix")
	}
	for _, row := range board {
		if len(row) != boardSize {
			return errors.New("invalid board: not a 9*9 matrix")
		}
		for _, val := range row {
			if val < 0 || val > 9 {
				return errors.New("invalid board: values outside the range [0, 9]")
			}
		}
	}

	// Valid input
	return nil
}

// solveSudokuRecursive is a recursive backtracking algorithm. Helper function
// for SolveSudoku. Modifies the board in place.
//
// Args:
//
//	board: A 2D slice(9x9) representing the Sudoku puzzle.
//
// Returns:
//
//	bool: true if the puzzle is solvable, false otherwise.
func solveSudokuRecursive(board [][]int) bool {

	// Iterate till zero is found
	zeroFound := false
	zeroRow, zeroCol := 0, 0

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == 0 {
				zeroRow = i
				zeroCol = j
				zeroFound = true
				break
			}
		}
	}

	// No zeroes => Board is solved
	if !zeroFound {
		return true
	}

	// Recursive Backtracking
	for candidate := 1; candidate <= boardSize; candidate++ {
		if isPlacementValid(board, zeroRow, zeroCol, candidate) {
			board[zeroRow][zeroCol] = candidate

			if solveSudokuRecursive(board) {
				return true
			}

			board[zeroRow][zeroCol] = 0
		}
	}

	return false
}

// isPlacementValid checks if placing a candidate at a specific position (row, col) in
// the Sudoku board is a valid move.
//
// Args:
//
//	board: A 2D slice representing the Sudoku puzzle.
//	row: The row index.
//	col: The column index.
//	candidate: The candidate value to be placed.
//
// Returns:
//
//	bool: true if the placement is valid, false otherwise.
func isPlacementValid(board [][]int, row, col, candidate int) bool {

	// Check validity in row and column
	for i := 0; i < boardSize; i++ {
		if board[row][i] == candidate || board[i][col] == candidate {
			return false
		}
	}

	// Check validity in the 3x3 subgrid
	startRow, startCol := gridSize*(row/gridSize), gridSize*(col/gridSize)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == candidate {
				return false
			}
		}
	}

	return true
}
