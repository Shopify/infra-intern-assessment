package sudoku

type Solver interface {
	// Solve will return a solved board
	Solve() [][]int
}
