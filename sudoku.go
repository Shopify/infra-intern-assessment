package main

import "sudoku/sudoku"

func SolveSudoku(board [][]int) [][]int {
	s := sudoku.NewSudokuSolver(board)
	solvedBoard, err := s.Solve()
	if err != nil {
		return nil
	}

	return solvedBoard
}
