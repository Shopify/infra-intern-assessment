package main

import (
	"errors"
)

const (
	boardsize   = 9
	subgridsize = 3
)

func SolveSudoku(input_board [][]int) [][]int {

	if err := inputvalidator(input_board); err != nil {
		return nil
	}

	output := make([][]int, boardsize)
	for i := range input_board {
		output[i] = make([]int, boardsize)
		copy(output[i], input_board[i])
	}

	if SudokuSolverRecursion(output) {
		return output
	} else {
		return nil
	}

}

func inputvalidator(board [][]int) error {

	if len(board) != boardsize {
		return errors.New("Error: != 9*9 matrix")
	}
	for _, row := range board {
		if len(row) != boardsize {
			return errors.New("Error: != 9*9 matrix")
		}
		for _, val := range row {
			if val < 0 || val > 9 {
				return errors.New("Error: [< 0 or > 9] values are not excepted")
			}
		}
	}

	return nil
}

func SudokuSolverRecursion(board [][]int) bool {

	zeroinBoard := false
	zeroinrow, zeroincol := 0, 0

	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
			if board[i][j] == 0 {
				zeroinrow = i
				zeroincol = j
				zeroinBoard = true
				break
			}
		}
	}

	if !zeroinBoard {
		return true
	}

	for possibleValue := 1; possibleValue <= boardsize; possibleValue++ {
		if isCellValid(board, zeroinrow, zeroincol, possibleValue) {
			board[zeroinrow][zeroincol] = possibleValue

			if SudokuSolverRecursion(board) {
				return true
			}

			board[zeroinrow][zeroincol] = 0
		}
	}
	return false
}

func isCellValid(board [][]int, row, col, possibleValue int) bool {

	for i := 0; i < boardsize; i++ {
		if board[row][i] == possibleValue || board[i][col] == possibleValue {
			return false
		}
	}

	startOfRow, startOfCol := subgridsize*(row/subgridsize), subgridsize*(col/subgridsize)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startOfRow+i][startOfCol+j] == possibleValue {
				return false
			}
		}
	}
	return true
}
