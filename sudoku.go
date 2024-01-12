package main

import "fmt"

// this class represents a sudoku board. It is simply a 2D slice of integers
type SudokuBoard struct{ board [][]int }

// creat new sudoku board from a 2D slice of integers
func NewSudokuBoard(board [][]int) *SudokuBoard { 
	return &SudokuBoard{board: board} 
}

// helper to print board
func (s *SudokuBoard) PrintBoard() {
	for _, row := range s.board {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
}

func (s *SudokuBoard) Solve() {
	// there are two main ways this can be solved. The first
	// method is to brute force all different possible combinations
	// of the board which is exponential in time (about 9^(n^2)) where
	// n is the number of 0's on the board

	// another solution is to use backtracking, which for every index,
	// tries a number and then recursively moves on to the next unfilled
	// index. once it is solved (or not solved) it will revert the change
	// and (depending on if this way worked) try the other numbers.

	// interestingly this has the same runtime since at worst you have to try
	// all possible combinations, but since this will "short-circuit" the moment
	// an incorrect move is made, the real run-time should be a lot faster.

	// short-circuit means pruning the search tree so we don't explore a move
	// that we know won't actually work
}

func SolveSudoku(board [][]int) [][]int {
	sudoku := NewSudokuBoard(board)
	// sudoku.PrintBoard()

	sudoku.Solve()

	return board
}
