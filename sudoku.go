package main

import "sudoku/sudoku"

func SolveSudoku(board [][]int) [][]int {
	s := sudoku.NewSudokuSolver(board)
	return s.Solve()
}
