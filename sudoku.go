package main

import (
	"fmt"
	"sudoku/sudoku"
	"time"
)

// SolveSudoku takes in a board and returns the solved version of the board
func SolveSudoku(board [][]int) [][]int {
	s := sudoku.NewSudokuSolver(board)
	solvedBoard, err := s.Solve()
	if err != nil {
		return nil
	}

	return solvedBoard
}

// SolveSudokuParallel takes in a board and returns the solved version of the board using the parallel solve
func SolveSudokuParallel(board [][]int) [][]int {
	s := sudoku.NewSudokuSolver(board)
	solvedBoard, err := s.ParallelSolve()
	if err != nil {
		return nil
	}

	return solvedBoard
}

// TestTime tests the timing of the sequential solve vs the parallelized one
//
// # Prints the elapsed time of Solve and ParallelSolve
//
// It shows that the sequential is much faster, however, that is most likely due to the parallel solve spinning up
// too many threads. A way to solve this could be limiting the amount of workers (to n workers) and then
// for each worker, have it take in a board from a buffered channel and fill in 1 cell (so like worst case, would
// need to push 9 boards to the buffered channel, 1 for each possible number). Each worker would do that and then
// read in a new board from the buffered channel and repeat. The only issue with this is that the buffer would
// need to be extremely large.
func TestTime(board [][]int) [][]int {
	s := sudoku.NewSudokuSolver(board)
	startTime := time.Now()
	solvedBoard, err := s.Solve()
	endTime := time.Now()
	elapsedTime1 := endTime.Sub(startTime)
	if err != nil {
		return nil
	}

	startTime = time.Now()
	_, err = s.ParallelSolve()
	endTime = time.Now()
	elapsedTime2 := endTime.Sub(startTime)
	if err != nil {
		return nil
	}

	fmt.Println("Sequential Solve time:", elapsedTime1)
	fmt.Println("Parallel Solve time:", elapsedTime2)

	return solvedBoard
}
