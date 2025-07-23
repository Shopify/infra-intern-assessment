package main

const BoardLength = 9
const BlockLength = 3

func CheckBlocks(board [][]int) bool {
	for row := 0; row < BoardLength; row += BlockLength {
		for col := 0; col < BoardLength; col += BlockLength {
			// create matrix and fill in values of the matrix
			matrix := make([][]int, BlockLength)
			for rows := range matrix {
				matrix[rows] = make([]int, BlockLength)
			}
			for r := 0; r < BlockLength; r ++ {
				for c := 0; c < BlockLength; c ++ {
					matrix[r][c] = board[row + r][col + c]
				}
			}

			// check values of the matrix
			check := make(map[int]bool)
			for i := 0; i < BlockLength; i++ {
				for j := 0; j < BlockLength; j++ {
					if (matrix[i][j] == 0) {
						return false;
					}
					_, exists := check[matrix[i][j]]
					if exists {
						return false
					} else {
						check[matrix[i][j]] = true
					}
				}
			}
		}
	}
	return true
}

func CheckRow(board [][]int) bool {
	for row := 0; row < BoardLength; row++ {
		check := make(map[int]bool)
		for i := 0; i < BoardLength; i++ {
			if (board[row][i] == 0) {
				return false;
			}
			_, exists := check[board[row][i]]
			if exists {
				return false
			} else {
				check[board[row][i]] = true
			}
		} 
	}
	return true
}

func getColumn(board[][] int, index int) []int {
	column := make([]int, BoardLength)
	for i:= 0; i < BoardLength; i++ {
		column[i] = board[i][index]
	}
	return column
}

func CheckCol(board [][]int) bool {
	for col := 0; col < BoardLength; col++ {
		column := getColumn(board, col)
		check := make(map[int]bool)
		for i := 0; i < BoardLength; i++ {
			if (column[i] == 0) {
				return false;
			}
			_, exists := check[column[i]]
			if exists {
				return false
			} else {
				check[column[i]] = true
			}
		} 
	}
	return true
}

// returns whether or not a sukdoku board is 
// solved or not
func CheckSudoku(board [][]int) bool {
	var blockCheck bool = CheckBlocks(board)
	var rowCheck bool = CheckRow(board)
	var colCheck bool = CheckCol(board)
	return blockCheck && rowCheck && colCheck
} 	


func checkTransition(board [][]int, row int, col int, val int) bool {
	// check row
	for i := 0; i < BoardLength; i++ {
		if (board[row][i] == val) && (i != col) {
			// value should be forfieted
			return false
		}
	}
	// check col
	for i := 0; i < BoardLength; i++ {
		if (board[i][col] == val) && (i != row) {
			// value should be forfieted
			return false
		}
	}
	// check 3x3 block 
	blockRow := (row / 3) * 3
	blockCol := (col / 3) * 3
	for i := blockRow; i < (blockRow + BlockLength); i++ {
		for j := blockCol; j < (blockCol + BlockLength); j++ {
			if (board[i][j] == val) && ((i != row) || (j != col)) {
				// value should be forfieted
				return false
			}
		}
	}
	return true
}


func getTransitions(board[][]int) [][][]int {
	transitions := make([][][]int, 0)

	// locate first empty section
	row := 0
	col := 0
	findEmpty:
	for r := 0; r < BoardLength; r ++ {
		for c := 0; c < BoardLength; c ++ {
			if (board[r][c] == 0) {
				row = r
				col = c
				break findEmpty
			}
		}
	}

	// Test Values that can possibly go in the 0
	for val := 1; val <= BoardLength; val ++ {
		if (checkTransition(board, row, col, val)) {
			matrix := make([][]int, len(board))
			for i := range board {
				matrix[i] = make([]int, len(board[i]))
				copy(matrix[i], board[i])
			}
			matrix[row][col] = val
			transitions = append(transitions, matrix)
		}
	}
	return transitions
}

// This function takes a board and returns the solved sudoku board
func SolveSudoku(board [][]int) [][]int {
	// if it's solved return it
	if (CheckSudoku(board)) {
		return board
	} else {
		matrix := make([][]int, len(board))
		// fill in a possible turn on a 0 
		transitions := getTransitions(board)
		for i := 0; i < len(transitions); i++ {
			res := SolveSudoku(transitions[i])
			// either we get no solved board and try a new transition
			// or we return the solved board
			if (len(res[0]) == 0) {
				continue
			} else if (CheckSudoku(res)) {
				for i := range res {
					matrix[i] = make([]int, len(res[i]))
					copy(matrix[i], res[i])
				}
				break
			}
		}
		return matrix
	}
}