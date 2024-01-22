package main

// Function to be called to solve a sudoku board of size 9 by 9
func SolveSudoku(board [][]int) [][]int {
	solve(board, 0, 0)
	return board
}

//recursive solution to solve input board
func solve(board [][]int, row, col int) bool {
	for i := row; i < 9; i, col = i+1, 0 {
		for j := col; j < 9; j++ {
			// assume inputs in place are valid
			if board[i][j] != 0 {
				continue
			}
			for num := 1; num <= 9; num++ {
				//check with backtracking if the current value is valid
				if isValid(board, i, j, int(num)) {
					board[i][j] = int(num)
					//if valid result, check the next column recursively
					if solve(board, i, j+1) {
						return true
					}
					// if previous result did not return, then the next column check was false
					// and we must reset the current attempt
					board[i][j] = 0
				}
			}
			return false
		}
	}

	return true
}
//this function checks rows/cols/and 3 by 3 blocks for a given board

func isValid(board [][]int, row, col int, num int) bool {
	boxRow, boxCol := (row/3)*3, (col/3)*3
	for i := 0; i < 9; i++ {
		//if any of the col/row/box tests fails, then return False
		if board[i][col] == num || board[row][i] == num ||
			board[boxRow+i/3][boxCol+i%3] == num {
			return false
		}
	}
	//if this point is reached, then the input was valid and "True" can be returned
	return true
}