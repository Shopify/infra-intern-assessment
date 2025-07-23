package main

import (
	"fmt"
)

func CheckGrid(row int, col int, val int, board [][]int) bool {
	// Check the row, column and grid for conflicts simultaneously
	for i := 0; i < 9; i++ {
		// row and column
		if board[row][i] == val || board[i][col] == val {
			return false
		}
		// Find the top left of the current 3x3 and add offset
		row_i := 3*(row/3) + i/3
		col_i := 3*(col/3) + i%3
		if board[row_i][col_i] == val {
			return false
		}
	}
	return true
}

func BackTrack(board [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// continue in undesirable case to reduce nesting
			if board[i][j] != 0 {
				continue
			}
			for val := 1; val <= 9; val++ {
				// continue in undesirable case to reduce nesting
				if !CheckGrid(i, j, val, board) {
					continue
				}
				board[i][j] = val
				if BackTrack(board) {
					return true
				}
				// if val doesn't result in a solved solution,
				// reset the value of the current cell
				board[i][j] = 0
			}
			// if all 9 values result in a conflict,
			// stop searching here and backtrack.
			return false
		}
	}
	return true
}

func PrintBoard(board [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func SolveSudoku(board [][]int) [][]int {
	BackTrack(board)
	PrintBoard(board)
	return board
}
