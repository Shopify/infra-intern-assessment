package main


// backtracking solution: 
// input: 2d array representing the sudoku board
// output: 2d array representing the solved sudoku board
// time complexity: O(9^(n*n)) where n is the size of the board
// space complexity: O(n*n) where n is the size of the board
/*
The Sudoku is solved using a backtracking algorithm. Optimizations could be made to the algorithm
by choosing the cell with the least number of possible values to try first. This would reduce the
number of recursive calls made, but not have any effect on the time complexity and will only result in a 
slight improvement for the size of the board.
*/
func SolveSudoku(board [][]int) [][]int {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for value := 1; value <= 9; value++ {
					// check if we can put the value in that cell
					if checkRow(board, row, value) && checkCol(board, col, value) && checkBox(board, row-row%3, col-col%3, value) {
						board[row][col] = value
						if SolveSudoku(board) != nil {
							return board
						}
						board[row][col] = 0
					}
				}
				return nil
			}
		}
	}
	// fully solved
	return board
}

func checkRow(board [][]int, row int, value int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == value {
			return false
		}
	}
	return true
}

func checkCol(board [][]int, col int, value int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == value {
			return false
		}
	}
	return true
}

func checkBox(board [][]int, startRow int, startCol int, value int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == value {
				return false
			}
		}
	}
	return true
}
