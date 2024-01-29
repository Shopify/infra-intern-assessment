package sudoku

const (
	_size        = 9
	_emptyCell   = 0
	_maxDigit    = 9
	_minDigit    = 1
	_subGridSize = 3
	_bufSize     = 1000
)

type boardData struct {
	board [][]int
	row   int
	col   int
}

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
		boardsToSolve: make(chan boardData),
		solvedBoard:   make(chan [][]int, _bufSize),
	}
}

// Solve returns a solved sudoku board
func (s *Solver) Solve(board [][]int) ([][]int, error) {

	return nil, nil
}

// isValidBoard checks if the board is valid once num is inserted into the specified row and col
func (s *Solver) isValidBoard(board [][]int, row, col, num int) bool {
	//for i := 0; i < _size; i++ {
	//
	//}
	return false
}
