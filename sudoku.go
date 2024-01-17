package main

import "math"

func isSafeToAssignCurrNum(board [][]int, row int, col int, currNum int) bool {

	for j := 0; j < len(board); j++ {
		if board[row][j] == currNum {
			return false
		}
	}
	for i := 0; i < len(board); i++ {

		if board[i][col] == currNum {
			return false
		}
	}

	sqrt := int(math.Sqrt(float64(len(board))))

	boxRowStart := row - row%sqrt
	boxColStart := col - col%sqrt

	for i := boxRowStart; i < boxRowStart+sqrt; i++ {
		for j := boxColStart; j < boxColStart+sqrt; j++ {
			if board[i][j] == currNum {
				return false
			}
		}
	}

	return true
}

func solveBoard(board [][]int, n int) bool {
	row := -1
	col := -1
	isBoardFilled := true

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 0 {
				row = i
				col = j

				isBoardFilled = false
				break
			}

		}
		if !isBoardFilled {
			break
		}
	}

	if isBoardFilled {
		return true
	}

	for num := 1; num <= n; num++ {
		if isSafeToAssignCurrNum(board, row, col, num) {
			board[row][col] = num

			if solveBoard(board, n) {
				return true
			} else {
				board[row][col] = 0
			}
		}
	}
	return false
}

func SolveSudoku(board [][]int) [][]int {
	solveBoard(board, len(board))
	return board
}
