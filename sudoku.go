package main
import "fmt"

func solveHelper(board [][]int, row int, col int) bool{

	//we check if we have reached the end of the board
	if row == 8 && col == 9 {
		return true
	}

	//check if reached end of column
	if col == 9 {
		row = row + 1
		col = 0
	}

	//check if the current cell is already filled
  if (board[row][col] != 0){
		return solveHelper(board, row, col + 1)
	}
 
	//a dfs type approach to solve the sudoku
	//iteratively try values from 1 to 9 and check if we can progress
  for num := 1; num < 10; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solveHelper(board, row, col + 1) {
				return true
			}
		}
		//resetting if invalid solve
		board[row][col] = 0
	}

	//if we reach here it means we have not found a valid solution
	return false
}

func isValid(board [][]int, row int, col int, num int) bool {

	// check row and column simultaneously
	for i:= 0; i < 9; i++ {

		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}
	}

	//check box
	startRow := row - row % 3
	startCol := col - col % 3

	for i:= 0; i < 3; i++ {
		for j:= 0; j < 3; j++ {
			if board[startRow + i][startCol + j] == num {
				return false
			}
		}
	}
	return true
}


func SolveSudoku(board [][]int) [][]int {

	//if we find a valid solution return
	if solveHelper(board, 0, 0){
		fmt.Println(board)
		return board

	}
	//else return an empty board
	return [][]int{}
}