package main

import (
	"fmt"
	"math"
)

// Scans board to determine if assigning currNum to position board[row][col] will conflict with an existing assignment.
func isSafeToAssignCurrNum(board [][]int, row int, col int, currNum int) bool {

	var boardLength = len(board)

	// Scans row to see if currNum already exists
	for j := 0; j < boardLength; j++ {
		if board[row][j] == currNum {
			return false
		}
	}

	// Scans column to see if currNum already exists
	for i := 0; i < boardLength; i++ {

		if board[i][col] == currNum {
			return false
		}
	}

	// modulo is used to determine the box boundaries at current row/column position
	sqrt := int(math.Sqrt(float64(boardLength)))
	boxRowStart := row - row%sqrt
	boxColStart := col - col%sqrt

	// scans current box boundary to determine if currNum already exists
	for i := boxRowStart; i < boxRowStart+sqrt; i++ {
		for j := boxColStart; j < boxColStart+sqrt; j++ {
			if board[i][j] == currNum {
				return false
			}
		}
	}

	// if scans never failed, then it is safe to assign currNum
	return true
}

// A recursive backtracking function used to solve a sudoku board, by assigning a number to an empty square,
// and recursively calling this function to assign a number to a different square. If at any point this fails, a square is
// reset to 0 and is then assigned a different currNum. This continues until all square have been assigned.
func solveBoard(board [][]int, n int) bool {
	row := -1
	col := -1
	isFullyFilled := true

	// find an empty square to be filled and assign indices to 'row' and 'col'
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 0 {
				row = i
				col = j

				isFullyFilled = false
				break
			}

		}
		// if not fully filled, break out of for loop and continue to assigning numbers step
		if !isFullyFilled {
			break
		}
	}

	// if fully filled, return as a correct solution has been found
	if isFullyFilled {
		return true
	}

	// from 1 to 9, try assigning a number to the empty square, backtracking if necessary
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

// SolveSudoku is the master function that calls solveBoard to solve the sudoku puzzle, populating 'board' with
// the correct values. The solved board is then printed out to the console.
func SolveSudoku(board [][]int) [][]int {
	solveBoard(board, len(board))

	fmt.Println("Solved Sudoku Board Output:")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%v,\n", board[i])
	}
	//fmt.Println(board)
	//printBoard(board, len(board))
	return board
}
