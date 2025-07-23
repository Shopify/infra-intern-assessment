package main

import "sudoku/sudoku"

// SolveSudoku returns a solved sudoku board
func SolveSudoku(board [][]int) [][]int {

	const _numWorkers = 100

	solver := sudoku.NewSolver(_numWorkers)

	solution, err := solver.Solve(board)
	if err != nil {
		panic(err)
	}

	return solution
}
