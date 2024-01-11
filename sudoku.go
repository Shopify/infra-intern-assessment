package main

const G = 9  // Grid dimensions (9x9)
const SG = 3 // Sub-grid dimensions (3x3)

// isValidPlacement receives a 9x9 sudoku board, row, col, and num, and verifies if the number can replace a 0 on the sudoku board
func isValidPlacement(board [][]int, row int, col int, num int) bool {

	// logic for validating cell and surrounding cells

	return true
}

// SolveSudoku recieves a 9x9 sudoku board and returns a solved version of the board.
func SolveSudoku(board [][]int) [][]int {
	row, col := 0, 0

	for {
		// Exit while-loop when board coordinate (9,9) reached
		if row >= G {
			break
		}

		if board[row][col] == 0 {
			for num := 1; num <= G; num++ {
				// Validate num placement on board
				if isValidPlacement(board, row, col, num) {
					board[row][col] = num

					// Return updated board with new num added
					if SolveSudoku(board) != nil {
						return board
					}

					board[row][col] = 0 // Backtrack num placement if new num XXX
				}
			}
			return nil // Cannot place num in this cell
		}

		col++

		// Move to next row when last column reached
		if col >= G {
			col = 0
			row++
		}
	}

	return board
}
