package main

import "fmt"

/*
input: board yet to be solved
output: solved board
solves the input board (assuming that there is a solution)
*/
func SolveSudoku(board [][]int) [][]int {
	solveHelper(&board)
	fmt.Println(board)
	return board
}

/*
input: board
output: bool
Recursive helper function
*/
func solveHelper(board *[][]int) bool {
	rowCol := findEmptySquare(*board)
	row := rowCol[0]
	col := rowCol[1]

	// checking if board is solved
	if row == -1 && col == -1 {
		return true
	}

	for i := 1; i < 10; i++ {
		if checkRow(*board, row, i) && checkCol(*board, col, i) && checkSquare(*board, col, row, i) {
			(*board)[row][col] = i
			if solveHelper(board) {
				return true
			}
			(*board)[row][col] = 0
		}
	}

	return false
}

/*
input: board, rowNum, check
output: bool
checks if the check number is in the specified board row.
returns false if the number is in the row, true if not
*/
func checkRow(board [][]int, rowNum int, check int) bool {
	for i := 0; i < 9; i++ {
		if board[rowNum][i] == check {
			return false
		}
	}
	return true
}

/*
input: board, colNum, check
output: bool
Checks if the check number is in the specified board column.
Returns false if the number is in the column, true if not.
*/
func checkCol(board [][]int, colNum int, check int) bool {
	for i := 0; i < 9; i++ {
		if board[i][colNum] == check {
			return false
		}
	}
	return true
}

/*
input: board, colNum, rowNum, check
output: bool
Checks if the check number is in the specified board square. A board square is a 3x3 subsquare in the board.
Returns false if the number is in the square, true if not
*/
func checkSquare(board [][]int, colNum int, rowNum int, check int) bool {
	row := (rowNum / 3) * 3
	col := (colNum / 3) * 3

	for row < (rowNum / 3 + 1) * 3 {
		for col < (colNum / 3 + 1) * 3 {
			if board[row][col] == check {
				return false
			}
			col++
		}
		row++
	}
	return true
}

/*
input: board
output: list of integers
Looks for the an empty square (if the value is 0 in the square).
If it is found, then it returns a list of the coordinates for the empty square [row, col].
Otherwise if there are no more empty squares, it returns [-1, -1]
*/
func findEmptySquare(board [][]int) []int {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return []int{row, col}
			}
		}
	}
	return []int{-1, -1}
}