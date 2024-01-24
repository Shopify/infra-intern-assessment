package main

import (
	"fmt"
)

const N = 9

// checks if a number can be placed at that row
func CanPlaceNumber(num int, row, col int, rows, cols, squares []map[int]bool) bool {
	return !rows[row][num] && !cols[col][num] && !squares[3*(row/3)+col/3][num]
}

// initializes a board (preload already filled values into the maps)
// we have maps to check if a number is already in a row, column, or 3x3 square
func InitializeBoard(board [][]int) ([]map[int]bool, []map[int]bool, []map[int]bool) {
	rows := make([]map[int]bool, N)
	cols := make([]map[int]bool, N)
	squares := make([]map[int]bool, N)

	for i := 0; i < N; i++ {
		rows[i] = make(map[int]bool)
		cols[i] = make(map[int]bool)
		squares[i] = make(map[int]bool)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			num := board[i][j]
			if num != 0 {
				rows[i][num] = true
				cols[j][num] = true
				squares[(i/3)*3+j/3][num] = true
			}
		}
	}

	return rows, cols, squares
}

// we use a backtracking approach

func SolveSudoku(board [][]int) [][]int {
	// create copy
	boardCopy := make([][]int, len(board))
	for i := range board {
		boardCopy[i] = make([]int, len(board[i]))
		copy(boardCopy[i], board[i])
	}

	rows, cols, squares := InitializeBoard(boardCopy)

	var solve func(row, col int) bool
	solve = func(row, col int) bool {

		// if we're out of the bounds of the board, we are done
		if row == N {
			return true
		}

		nextRow, nextCol := row, col+1

		if col == N-1 {
			nextRow, nextCol = row+1, 0
		}

		if boardCopy[row][col] != 0 {
			return solve(nextRow, nextCol)
		}

		// backtracking step

		for num := 1; num <= N; num++ {
			if CanPlaceNumber(num, row, col, rows, cols, squares) {
				boardCopy[row][col] = num
				rows[row][num] = true
				cols[col][num] = true
				squares[3*(row/3)+col/3][num] = true

				if solve(nextRow, nextCol) {
					return true
				}

				boardCopy[row][col] = 0
				delete(rows[row], num)
				delete(cols[col], num)
				delete(squares[3*(row/3)+col/3], num)
			}
		}

		return false
	}

	if solve(0, 0) {
		return boardCopy
	} else {
		return nil // return nil if the puzzle cannot be solved
	}
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
