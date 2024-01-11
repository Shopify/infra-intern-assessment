package main

// Alexander Carvalho [01-11-2024]
//
//

const G = 9  // Grid dimensions (9x9)
const SG = 3 // Sub-grid dimensions (3x3)

// `isValidPlacement` receives a 9x9 sudoku board, row, col, and num, and verifies if the number can replace a 0 on the sudoku board
func isValidPlacement(board [][]int, row int, col int, num int) bool {
	// Check if `num` exists in the current row or column
	for idx := 0; idx < G; idx++ {
		if board[row][idx] == num || board[idx][col] == num {
			return false
		}
	}

	// Check if `num` exists in the 3x3 sub-grid
	rowInit, colInit := row-(row%3), col-(col%3) // use modulus to get coordinates of sub-grid start
	for rowIdx := 0; rowIdx < SG; rowIdx++ {
		for colIdx := 0; colIdx < SG; colIdx++ {
			if board[rowInit+rowIdx][colInit+colIdx] == num {
				return false
			}
		}
	}

	// Can place `num` in this cell
	return true
}

// `SolveSudoku` recieves a 9x9 sudoku board and returns a solved version of the board.
func SolveSudoku(board [][]int) [][]int {
	row, col := 0, 0

	for {
		// Exit while-loop when board coordinate (9,9) reached
		if row >= G {
			break
		}

		if board[row][col] == 0 {
			for num := 1; num <= G; num++ {
				// Validate `num` placement on board
				if isValidPlacement(board, row, col, num) {
					board[row][col] = num

					// Validate XXX
					if SolveSudoku(board) != nil {
						return board
					}

					board[row][col] = 0 // Backtrack `num`placement if XXX
				}
			}
			return nil // Cannot place `num` in this cell
		}

		col++

		// Increment `row` when last `col` reached, reset `col` index
		if col >= G {
			col = 0
			row++
		}
	}

	return board
}
