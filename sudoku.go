package main

const gridSize = 9

func SolveSudoku(board [][]int) [][]int {
	solveSudokuRecursive(board)
	return board
}

func solveSudokuRecursive(board [][]int) bool {
	empty := findEmptyCell(board)
	if empty[0] == -1 && empty[1] == -1 {
		return true // No empty cells, puzzle solved
	}

	row, col := empty[0], empty[1]

	for num := 1; num <= 9; num++ {
		if isValidMove(board, row, col, num) {
			board[row][col] = num

			if solveSudokuRecursive(board) {
				return true
			}

			board[row][col] = 0 // Backtrack if placement is invalid
		}
	}

	return false
}

func findEmptyCell(board [][]int) [2]int {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if board[i][j] == 0 {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1} // No empty cell found
}

func isValidMove(board [][]int, row, col, num int) bool {
	// Check if the number is not present in the current row and column
	for i := 0; i < gridSize; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Check if the number is not present in the current 3x3 subgrid
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}
