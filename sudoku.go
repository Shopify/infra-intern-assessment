package main

func SolveSudoku(board [][]int) [][]int {
	out := make([][]int, len(board))
	for i := range board {
		out[i] = make([]int, len(board[i]))
		copy(out[i], board[i])
	}
	solve(&out)
	return out
}

func solve(board *[][]int) bool {
	empty := []int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*board)[i][j] == 0 {
				empty = append(empty, i, j)
			}
		}
	}
	if len(empty) == 0 {
		//finished solving sudoku, no empty cells
		return true
	}

	r, c := empty[0], empty[1]

	for num := 1; num <= 9; num++ {
		if isValid(*board, r, c, num) {
			(*board)[r][c] = num
			//recursively determine if num is valid for board[r][c]
			if solve(board) {
				return true
			}
			//if board[r][c] = num is invalid, then reset the cell
			(*board)[r][c] = 0
		}
	}

	return false // No valid number found for the current cell
}

func isValid(board [][]int, row, col, num int) bool {
	// validate row and col
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// valid 3x3 subgrid for duplicates
	x, y := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[x+i][y+j] == num {
				return false
			}
		}
	}
	//return true if grid is valid
	return true
}
