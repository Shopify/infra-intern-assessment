package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
func SolveSudoku(board [][]int) [][]int {
	computeSudoku(board)
	return board
}

func computeSudoku(board [][]int) bool {
	//nested for loop that checks for an empty cell
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 0 {
				//trying all digit possibilities on an empty cell
				for k := 1; k <= 9; k++ {
					//if the digit respects the validity of sudoku grid then update cell with digit
					if validSudoku(board, i, j, k) {
						board[i][j] = k
						//recursive call function again until there is no more empty cells
						if computeSudoku(board) {
							return true
						} else {
							board[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func validSudoku(board [][]int, row, col int, k int) bool {
	for i := 0; i < 9; i++ {
		// checks row validity
		if board[i][col] != 0 && board[i][col] == k {
			return false
		}
		// checks column
		if board[row][i] != 0 && board[row][i] == k {
			return false
		}
		// checks 3*3 grid using modulo and integer division
		if board[3*(row/3)+i/3][3*(col/3)+i%3] != 0 && board[3*(row/3)+i/3][3*(col/3)+i%3] == k {
			return false
		}
	}
	return true
}
