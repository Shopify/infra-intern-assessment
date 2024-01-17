package main

const (
	LEGAL_SUDOKU_SIZE = 9 // legal size of the input sudoku

)

func SolveSudoku(board [][]int) [][]int {

	// check input size
	if len(board) != LEGAL_SUDOKU_SIZE {
		reportErr(ERR_INVALID_SUDOKU_SIZE)
		return board
	}
	for _, row := range board {
		if len(row) != LEGAL_SUDOKU_SIZE {
			reportErr(ERR_INVALID_SUDOKU_SIZE)
			return board
		}
	}

	// Initialize slices to keep track of the numbers used in each row, column, and zone (3x3 grid),
	// storing this information as bits.
	var rows, cols, grids = make([]int, 9), make([]int, 9), make([]int, 9)
	for i := range board {
		for j := range board[i] {
			if board[i][j] != 0 {
				var bit = 1 << board[i][j]
				rows[i] |= bit          // Mark the bit as used in the row.
				cols[j] |= bit          // Mark the bit as used in the column.
				grids[i/3*3+j/3] |= bit // Mark the bit as used in the zone.
			}
		}
	}
	dfs(board, 0, rows, cols, grids)
	return board
}
