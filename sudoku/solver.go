package sudoku

import "errors"

const (
	_size        = 9
	_emptyCell   = 0
	_maxDigit    = 9
	_minDigit    = 1
	_subGridSize = 3
	_bufSize     = 100
)

// boardData is the information required for a goroutine to perform backtracking
type boardData struct {
	board [][]int
	row   int
	col   int
}

// Solver contains members for solving a sudoku board
type Solver struct {
	numWorkers int

	boardsToSolve chan boardData
	solvedBoard   chan [][]int

	done bool
}

// NewSolver returns an instance of Solver
func NewSolver(numWorkers int) *Solver {
	return &Solver{
		numWorkers:    numWorkers,
		boardsToSolve: make(chan boardData, _bufSize),
		solvedBoard:   make(chan [][]int),
	}
}

// Solve returns a solved sudoku board by spinning up a set number of goroutines which work towards a solution
func (s *Solver) Solve(board [][]int) ([][]int, error) {
	if board == nil {
		return nil, errors.New("board can't be nil")
	}

	for i := 0; i < s.numWorkers; i++ {
		go s.parallelSolve()
	}

	s.boardsToSolve <- boardData{
		board: board,
		row:   0,
		col:   0,
	}

	solvedBoard := <-s.solvedBoard

	return solvedBoard, nil
}

// parallelSolve is the work a single goroutine does to work towards the solution
func (s *Solver) parallelSolve() {
	var board, newBoard boardData

	for !s.done {
		board = <-s.boardsToSolve

		if board.row == _size {
			s.done = true
			s.solvedBoard <- board.board
			return
		}

		if board.col == _size {
			board.row++
			board.col = 0

			s.boardsToSolve <- board

			continue
		}

		if board.board[board.row][board.col] == _emptyCell {
			for digit := _minDigit; digit <= _maxDigit; digit++ {
				if s.isValidBoard(board.board, board.row, board.col, digit) {
					newBoard = boardData{
						board: copyBoard(board.board),
						row:   board.row,
						col:   board.col + 1,
					}

					newBoard.board[board.row][board.col] = digit

					s.boardsToSolve <- newBoard
				}
			}
		} else {
			board.col += 1
			s.boardsToSolve <- board
		}
	}
}

// isValidBoard checks if the board is valid once num is inserted into the specified row and col
func (s *Solver) isValidBoard(board [][]int, row, col, num int) bool {
	for i := 0; i < _size; i++ {
		subGridVal := board[_subGridSize*(row/_subGridSize)+(i/_subGridSize)][_subGridSize*(col/_subGridSize)+i%_subGridSize]
		colVal := board[row][i]
		rowVal := board[i][col]

		if subGridVal == num || colVal == num || rowVal == num {
			return false
		}
	}
	return true
}

// copyBoard returns a complete copy of a board
func copyBoard(board [][]int) [][]int {
	cpy := make([][]int, len(board))

	for i := 0; i < len(board[0]); i++ {
		cpy[i] = make([]int, len(board[0]))
		copy(cpy[i], board[i])
	}

	return cpy
}
