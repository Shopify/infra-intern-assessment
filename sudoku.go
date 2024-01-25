package main

func SolveSudoku(grid [][]int) [][]int {
	sodoku := NewSodoku(grid, GRID_ROW_SIZE, GRID_COL_SIZE)
	sodoku.Solve()
	if sodoku.solved {
		return sodoku.solvedGrid
	}
	return sodoku.grid
}
