package main

import (
	"fmt"
)

const SudokuSize = 9
const SudokuBoxSize = 3

// Stores index of position on sudoku grid
type Position struct {
	row    int // Row index
	column int // Column index
}

func SolveSudoku(input [][]int) [][]int {
	isSolved, Solution := SudokuSolver(input)
	if !isSolved {
		fmt.Println("Sudoku can't be solved")
		return nil
	}
	return Solution
}
