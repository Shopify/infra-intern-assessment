package sudoku

type Solver interface {
	// Will return a solved board
	Solve() [][]int
}
