package main

import (
	"errors"
	"fmt"
	"math"
)

const BoardSize = 9

// solve sudoku board with backtracking
// basic steps are:
// 1. find empty box / check if board is solved
// 2. find a valid number to put in empty box
// 3. repeat 1 - 2, backtracking if chosen number leads to unsolvable board
func SolveSudoku(board [][]int) [][]int {
	solveSudoku(board)
	printSudokuboard(board)
	return board
}

func solveSudoku(board [][]int) bool {
	row, col, err := findEmptyBox(board)
	// board is already solved (no empty boxes)
	if err != nil {
		return true
	}

	// try valid nums for current (empty) box
	for num:= 1; num <= BoardSize; num++ {
		if numValidPlacement(board, row, col, num) {
			board[row][col] = num;
			if solveSudoku(board) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

// finds empty box in sudoku board
func findEmptyBox(board [][]int) (int, int, error) {
	for row := 0; row < BoardSize; row ++ {
		for col := 0; col < BoardSize; col ++ {
			if board[row][col] == 0 {
				return row, col, nil
			}
		}
	}
	return 0, 0, errors.New("board is solved. no empty boxes")
}

func numValidPlacement(board [][]int, row int, col int, num int) bool {
	return numValidInRow(board, row, num) &&
		numValidInCol(board, col, num) &&
		numValidInSection(
			board, 
			row - row % int(math.Sqrt(BoardSize)),
			col - col % int(math.Sqrt(BoardSize)),
			num)
}

func numValidInRow(board [][]int, row int, num int) bool {
	for col := 0; col < BoardSize; col++ {
		if board[row][col] == num {
			return false
		}
	}
	return true
}

func numValidInCol(board [][]int, col int, num int) bool {
	for row := 0; row < BoardSize; row++ {
		if board[row][col] == num {
			return false
		}
	}
	return true
}

func numValidInSection(board [][]int, rowStart int, colStart int, num int) bool {
	for row := 0; row < int(math.Sqrt(BoardSize)); row++ {
		for col := 0; col < int(math.Sqrt(BoardSize)); col++ {
			if board[rowStart + row][colStart + col] == num {
				return false
			}
		}
	}
	return true
}

func printSudokuboard(board [][]int) {
	fmt.Println()
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
