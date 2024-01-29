package main

func SolveSudoku(board [][]int) [][]int { // Returns the solved board if there is a solution
	if solve(board) {
		return board
	}
	return nil
}

// helper function
func solve(board [][]int) bool { // Solves the sudoku board using backtracking and if correct number adds it to the grid
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						if solve(board) {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// helper function
func isValid(board [][]int, row, col, num int) bool { //Checks whether the number is present in column row or sub-grid
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
		if board[row][i] == num {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}
