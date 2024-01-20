package sudoku

import (
	"github.com/pkg/errors"
)

const boardSize = 9

// SudokuSolver inherits from Solver
type SudokuSolver struct {
	board [][]int
}

// NewSudokuSolver creates a sudoku solver object.
func NewSudokuSolver(board [][]int) *SudokuSolver {
	return &SudokuSolver{
		board: board,
	}
}

// Solve will solve the board in a sequential fashion
func (s *SudokuSolver) Solve() ([][]int, error) {
	// Check for invalid boards
	if s.board == nil {
		return nil, errors.New("board is nil")
	}

	if !isValidBoardSize(s.board) {
		return nil, errors.New("invalid board size")
	}

	// Need to create a deep copy since solve will modify the board
	solved := deepCopy(s.board)
	s.solve(solved, 0, 0)
	return solved, nil
}

// solve will solve the sudoku board
//
// Algorithm:
//
// Start on the top left cell and check if it's empty.
// If so, try filling in all numbers from 1 to 9 (checking if it's a valid number before trying it). Recursively solve the board with the new number filled in
// If not, move onto the next cell
// Do this for all cells until a solution has been reached
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
		for i := 9; i >= 1; i-- {
			if s.isValid(board, row, col, i) { // check if i is a possibility for board[row][col]
				board[row][col] = i

				// try the number, if it works, return true
				// else set the slot back to empty so we can try the next number
				if s.solve(board, row, col+1) {
					return true
				}

				board[row][col] = 0
			}
		}
		return false
	} else {
		// continue solving next slot
		return s.solve(board, row, col+1)
	}
}

// ParallelSolve will solve the board with parallelism
func (s *SudokuSolver) ParallelSolve() ([][]int, error) {
	// Check for invalid boards
	if s.board == nil {
		return nil, errors.New("board is nil")
	}

	if !isValidBoardSize(s.board) {
		return nil, errors.New("invalid board size")
	}

	// create a channel to receive solutions
	resultChannel := make(chan [][]int)

	// Need to create a deep copy since solve will modify the board
	board := deepCopy(s.board)

	// start solving in parallel
	go s.parallelSolve(board, 0, 0, resultChannel)

	// wait for the solution
	solved := <-resultChannel

	return solved, nil
}

// parallelSolve will solve the sudoku board in parallel
//
// Algorithm:
//
// Start on the top left cell and check if it's empty.
// If so, try filling in all numbers from 1 to 9 (checking if it's a valid number before trying it). Recursively solve the board with the new number filled in
// For each number, it spins up a new goroutine to try and solve
// If not, move onto the next cell
// Do this for all cells until a solution has been reached
func (s *SudokuSolver) parallelSolve(board [][]int, row, col int, resultChannel chan<- [][]int) {
	// if whole board finished, send the solution to the channel
	if row == boardSize {
		// send final answer to resultChannel
		resultChannel <- board
		return
	}

	// if at end of row, go to the next one
	if col == boardSize {
		s.parallelSolve(board, row+1, 0, resultChannel)
		return
	}

	// only solve if its empty slot
	if board[row][col] == 0 {
		// start a goroutine for each possibility
		for i := 9; i >= 1; i-- {
			if s.isValid(board, row, col, i) {
				// create a new board for the goroutine
				newBoard := deepCopy(board)
				newBoard[row][col] = i

				go s.parallelSolve(newBoard, row, col+1, resultChannel)
			}
		}
	} else {
		// continue solving next slot
		s.parallelSolve(board, row, col+1, resultChannel)
	}
}

// checks if num is input into board[row][col], will the sudoku board be valid
func (s *SudokuSolver) isValid(board [][]int, row, col int, num int) bool {
	for i := 0; i < boardSize; i++ {
		// based on row and col, grab the elements (in logical order)
		inCol := board[i][col]
		inRow := board[row][i]
		inGrid := board[3*(row/3)+i/3][3*(col/3)+i%3]
		// inGrid does math to get the ith cell in the sub-grid
		// 0 1 2
		// 3 4 5
		// 6 7 8

		if inCol == num || inRow == num || inGrid == num {
			return false
		}
	}

	return true
}

// deep copy's a 2D matrix (so we don't modify the original board in place as 2D matrix gets passed by reference)
func deepCopy(matrix [][]int) [][]int {
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
