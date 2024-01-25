package main

const (
	GRID_ROW_SIZE    = 9
	GRID_COL_SIZE    = 9
	MAX_POSSIBLE_NUM = 9
	SQUARE_SIZE      = 3
	EMPTY_SPACE      = 0
)

type Solver interface {
	Solve()
}

type Sodoku struct {
	grid         [][]int
	solvedGrid   [][]int
	solved       bool
	NumberOfCols uint8
	NumberOfRows uint8
}

func NewSodoku(grid [][]int, numberOfCols, numberOfRows uint8) *Sodoku {
	// copying the grid
	solvedGrid := make([][]int, numberOfRows)
	for i := 0; i < int(numberOfRows); i++ {
		solvedGrid[i] = make([]int, numberOfCols)
		for j := 0; j < int(numberOfCols); j++ {
			solvedGrid[i][j] = grid[i][j]
		}
	}

	return &Sodoku{
		grid:         grid,
		solvedGrid:   solvedGrid,
		solved:       false,
		NumberOfCols: numberOfCols,
		NumberOfRows: numberOfRows,
	}
}
