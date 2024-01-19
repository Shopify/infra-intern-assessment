package sudoku

import "github.com/pkg/errors"

const boardSize = 9

// inherits from Solver
type SudokuSolver struct {
	board [][]int
}

// NewSudokuSolver creates a sudoku solver object.
func NewSudokuSolver(board [][]int) *SudokuSolver {
	return &SudokuSolver{
		board: board,
	}
}

// checks if num is input into board[row][col], will the sudoku board be valid
func (s *SudokuSolver) isValid(board [][]int, row, col int, num int) bool {
	for i := 0; i < boardSize; i++ {
		if board[i][col] == num || board[row][i] == num || board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

// solves the board
func (s *SudokuSolver) solve(board [][]int, row, col int) bool {
	// if whole board finished, return true
	if row == boardSize {
		return true
	}
	// if at end of row, go to the next one
	if col == boardSize {
		return s.solve(board, row+1, 0)
	}

	// only solve if its empty slot
	if board[row][col] == 0 {
		for i := 1; i <= 9; i++ {
			if s.isValid(board, row, col, i) { // check if i is a possibility for board[row][col]
				board[row][col] = i

				// try the number, if it works, return true
				// else set the slot back to empty so we can try the next number
				if s.solve(board, row, col+1) {
					return true
				} else {
					board[row][col] = 0
				}
			}
		}
		return false
	} else {
		return s.solve(board, row, col+1)
	}
}

// deep copy's a 2D matrix (so we don't modify the original board in place as 2D matrix gets passed by reference)
func (s *SudokuSolver) deepCopy(matrix [][]int) [][]int {
	copiedMatrix := make([][]int, len(matrix))

	for i := range matrix {
		copiedMatrix[i] = make([]int, len(matrix[i]))
		copy(copiedMatrix[i], matrix[i])
	}

	return copiedMatrix
}

// isValidBoardSize checks if the board has a valid size
func isValidBoardSize(board [][]int) bool {
	if len(board) != boardSize {
		return false
	}
	for _, row := range board {
		if len(row) != boardSize {
			return false
		}
	}
	return true
}

// Solve the board
func (s *SudokuSolver) Solve() ([][]int, error) {
	// Check for invalid boards
	if s.board == nil {
		return nil, errors.New("board is nil")
	}

	if !isValidBoardSize(s.board) {
		return nil, errors.New("invalid board size")
	}

	solved := s.deepCopy(s.board)
	s.solve(solved, 0, 0)
	return solved, nil
}
