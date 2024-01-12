// Alexander Carvalho [01-11-2024]
//
// This is a solution to the Shopify Intern Assessment Production Engineering
// Sudoku problem using backtracking.
//
// The solution is based on the following assumptions:
// - The input grid will be a 9x9 two-dimensional array of integers.
// - The input grid will have exactly one solution.
// - The input grid may contain zeros (0) to represent empty cells.

package main

const GridDims, SubGridDims = 9, 3

// `isValidPlacement` receives a 9x9 sudoku board, row, col, and num,
// and verifies if the number can replace a 0 on the sudoku board.
func isValidPlacement(board [][]int, row int, col int, num int) bool {
	// Check if `num` exists in the current row or column
	for idx := 0; idx < GridDims; idx++ {
		if board[row][idx] == num || board[idx][col] == num {
			return false
		}
	}

	// Check if `num` exists in the 3x3 sub-grid
	rowInit, colInit := row-(row%3), col-(col%3) // get coordinates of sub-grid starting cell
	for rowIdx := 0; rowIdx < SubGridDims; rowIdx++ {
		for colIdx := 0; colIdx < SubGridDims; colIdx++ {
			if board[rowInit+rowIdx][colInit+colIdx] == num {
				return false
			}
		}
	}

	// Can place `num` in this cell
	return true
}

// `SolveSudoku` recieves a 9x9 sudoku board and returns a solved version of the board.
// If no solution is found, the function returns nil.
func SolveSudoku(board [][]int) [][]int {
	row, col := 0, 0

	// Iterate through board cells until empty cell is found
	for {
		// Increment `row` when last `col` reached, reset `col` index
		if col >= GridDims {
			col = 0
			row++
		}

		// Exit while-loop when board coordinate (9,9) reached
		if row >= GridDims {
			break
		}

		// Check if cell is empty
		if board[row][col] == 0 {
			// Iterate through possible numbers to place in cell
			for num := 1; num <= GridDims; num++ {
				// Validate `num` placement on board
				if isValidPlacement(board, row, col, num) {
					board[row][col] = num

					// Validate next board coordinate
					if SolveSudoku(board) != nil {
						return board
					}

					board[row][col] = 0 // Backtrack `num` placement if no solution found
				}
			}

			return nil // Cannot place `num` in this cell
		}

		col++
	}

	return board
}
