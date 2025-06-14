package main

// Check the sub array of the sudoku board (3x3) to fetch the available numbers for the current cell.
//
// * Input requirements:
//   - Each element in the board array must have value from 0-9, where 0 represents an empty (unfilled) cell. The element with a value from 1-9 represents a filled cell.
//   - The sudoku board must have only one single solution.
//   - row and col must be integers in range [0, 8]
//   - availableNums must be a map of integers from 1-9
//
// Return nil: we adjust the availableNums map in-place
func CheckSubSudoku(board [][]int, row, col int, availableNums map[int]bool) {
	// top, bottom, left, and right represent the boundary of the current sub-sudoku (3x3)
	var top, bottom, left, right int

	if col <= 2 {
		left, right = 0, 2
	} else if col <= 5 {
		left, right = 3, 5
	} else {
		left, right = 6, 8
	}

	if row <= 2 {
		top, bottom = 0, 2
	} else if row <= 5 {
		top, bottom = 3, 5
	} else {
		top, bottom = 6, 8
	}

	// Loop to see if there is used numbers within our current sub-sudoku (3x3)
	// If yes, remove that number from our availableNums
	for i := top; i <= bottom; i++ {
		for j := left; j <= right; j++ {
			if board[i][j] != 0 {
				delete(availableNums, board[i][j])
			}
		}
	}
}

// Get the available numbers to fill the current cell
//
// * Input requirements:
//   - Each element in the board array must have value from 0-9, where 0 represents an empty (unfilled) cell. The element with a value from 1-9 represents a filled cell.
//   - The sudoku board must have only one single solution.
//   - row and col must be integers in range [0, 8]
//
// Return a map of integers from 1-9, where the key represents the available number to fill the current cell.
// The returned map can be empty (no available number to fill the current cell)
func GetAvailableNumbers(board [][]int, row, col int) map[int]bool {
	availableNums := make(map[int]bool)
	for i := 1; i <= 9; i++ {
		availableNums[i] = true
	}
	// Check if any numbers have been used in the same row
	for i := 0; i < len(board[0]); i++ {
		if board[row][i] != 0 {
			delete(availableNums, board[row][i])
		}
	}
	// Check if any numbers have been used in the same column
	for j := 0; j < len(board); j++ {
		if board[j][col] != 0 {
			delete(availableNums, board[j][col])
		}
	}
	// Check the sub-sudoku (3x3) to scope down the available numbers to fill
	CheckSubSudoku(board, row, col, availableNums)

	return availableNums
}

// Fill the sudoku board recursively using backtracking
//
// * Input requirements:
//   - Each element in the board array must have value from 0-9, where 0 represents an empty (unfilled) cell. The element with a value from 1-9 represents a filled cell
//   - The sudoku board must have only one single solution.
//
// Return a boolean value: true if the sudoku board is filled, false otherwise
func FillBoard(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				// if the current cell is not filled, get the available digits from 0 -> 9 to fill that cell.
				availableNumbers := GetAvailableNumbers(board, i, j)
				for num := range availableNumbers {
					// Fill each available number to the board & recursively call FillBoard(board) to fill the next empty cell if possible.
					board[i][j] = num
					if FillBoard(board) {
						// We have found a valid way to solve our sudoku board => terminate the recursion immediately.
						return true
					}
					// Backtrack the cell to the previous unfilled state (0)
					board[i][j] = 0
				}
				return false
			}
		}
	}
	return true
}

// Given a 2D 9x9 word that represents a sudoku board
//
// * Input requirements:
//   - Each element in the board array must have value from 0-9, where 0 represents an empty (unfilled) cell. The element with a value from 1-9 represents a filled cell
//   - The sudoku board must have only one single solution.
//
// Return a 2D 9x9 array with all the cells filled (no elements is equal to 0)
func SolveSudoku(board [][]int) [][]int {
	FillBoard(board)
	return board
}
