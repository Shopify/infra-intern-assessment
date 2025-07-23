package main

import (
	"errors"
)

const (
	// BOARD_SIZE represents the size of the Sudoku board.
	BOARD_SIZE = 9
	// SMALL_SQUARE_SIZE represents the size of each small square in the Sudoku board.
	SMALL_SQUARE_SIZE = 3
)

// Position stores the row and column of a position in the board.
type Position struct {
	row int
	col int
}

// SolveSudoku solves a Sudoku puzzle and returns the solved board.
func SolveSudoku(board [][]int) [][]int {
	board, _ = SolveSudokuHelper(board, Position{0, 0})
	return board
}

// SolveSudokuHelper is a recursive helper function for solving Sudoku.
// It explores possible values for each unfilled position and backtracks if necessary.
func SolveSudokuHelper(board [][]int, curPosition Position) ([][]int, error) {
	nextUnfilled, err := nextUnfilled(board, curPosition)
	if err != nil {
		return board, nil
	}
	for value := 1; value <= 9; value++ {
		validValue := isValuePlaceable(board, nextUnfilled, value)
		if validValue {
			boardWithValue := deepCopy2DSlice(board)
			boardWithValue[nextUnfilled.row][nextUnfilled.col] = value
			newBoard, err := SolveSudokuHelper(boardWithValue, nextUnfilled)
			if err == nil {
				return newBoard, nil
			}
		}
	}
	return board, errors.New("no valid board states found")
}

// deepCopy2DSlice creates a deep copy of a 2D slice.
func deepCopy2DSlice(board [][]int) [][]int {
	boardWithValue := make([][]int, len(board))
	for i := range board {
		boardWithValue[i] = make([]int, len(board[i]))
		copy(boardWithValue[i], board[i])
	}
	return boardWithValue
}

// nextUnfilled finds the next unfilled position in the Sudoku board.
func nextUnfilled(board [][]int, curPosition Position) (Position, error) {
	for col := curPosition.col; col < BOARD_SIZE; col++ {
		if board[curPosition.row][col] == 0 {
			return Position{curPosition.row, col}, nil
		}
	}
	curPosition.row++
	for row := curPosition.row; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			if board[row][col] == 0 {
				return Position{row, col}, nil
			}
		}
	}
	return Position{}, errors.New("no next unfilled")
}

// isValuePlaceable checks if a value can be placed at a given position in the Sudoku board.
func isValuePlaceable(board [][]int, testingPos Position, value int) bool {
	// Check if the column of the position already has the value.
	for row := 0; row < BOARD_SIZE; row++ {
		if board[row][testingPos.col] == value {
			return false
		}
	}
	// Check if the row of the position already has the value.
	for col := 0; col < BOARD_SIZE; col++ {
		if board[testingPos.row][col] == value {
			return false
		}
	}
	// Check if the small square of the position already has the value.
	startRow := testingPos.row - testingPos.row%SMALL_SQUARE_SIZE
	startCol := testingPos.col - testingPos.col%SMALL_SQUARE_SIZE
	for row := startRow; row < startRow+SMALL_SQUARE_SIZE; row++ {
		for col := startCol; col < startCol+SMALL_SQUARE_SIZE; col++ {
			if board[row][col] == value {
				return false
			}
		}
	}
	// If it's in none of them, then this placement is valid.
	return true
}
