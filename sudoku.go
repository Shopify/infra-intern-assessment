package main

func SolveSudoku(board [][]int) [][]int {
	Solver(board, 0, 0)
	return board
}

// determine if there is a valid entry for this cell and all remaining empty cells
// assumption: a valid solution exists for the input board
func Solver(board [][]int, x int, y int) bool {
	// base case: we already checked the last cell (8,8) which was valid
	if x == 0 && y == 9 {
		return true
	}
	// cell content != 0 means it's fixed, so move onto the next cell
	if board[y][x] != 0 {
		x, y = GetNext(x, y)
		return Solver(board, x, y)
	}
	if board[y][x] == 0 {
		// increment through 1-9 for this cell, and return
		for i := 1; i <= 9; i++ {
			board[y][x] = i
			if CheckCell(board, x, y) {
				xn, yn := GetNext(x, y)
				if Solver(board, xn, yn) {
					return true
				}
				// otherwise, the board is not valid with this current value
				// continue incrementing
			}
		}
		// all values 1-9 for this cell are invalid, so reset this cell
		// the 'previous' cell should increment itself
		board[y][x] = 0
	}
	return false
}

func GetNext(x int, y int) (int, int) {
	x++
	// if we are at the end of a row, 'loop around' to the next
	if x > 8 {
		y++
		x = 0
	}
	return x, y
}

// check in this row, column, and 3x3 square that no other value
// is equal to board[y][x]
func CheckCell(board [][]int, x int, y int) bool {
	current := board[y][x]
	// check columns and rows, skipping the current cell
	for i := 0; i <= 8; i++ {
		if i != y && board[i][x] == current {
			return false
		}
		if i != x && board[y][i] == current {
			return false
		}
	}
	// 'round down' to the top left cell in the nearest 3x3 grid
	lowx := x - (x % 3)
	lowy := y - (y % 3)
	// check the current value is not repeated in the local 3x3 grid
	for xi := lowx; xi < lowx+3; xi++ {
		for yi := lowy; yi < lowy+3; yi++ {
			if xi != x && yi != y {
				if board[yi][xi] == current {
					return false
				}
			}
		}
	}
	return true
}
