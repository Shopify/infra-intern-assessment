package main

import (
	"fmt"
)

const SudokuSize = 9
const SudokuBoxSize = 3

type Position struct {
	row    int // The value of row index
	column int // The value of column index
}

func SolveSudoku(input [][]int) [][]int {
	isSolved, Solution := SudokuSolver(input)
	if !isSolved {
		fmt.Println("Sudoku can't be solved")
		return nil
	}
	return Solution
}

func main() {
	input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	// input := [][]int{
	// 	{9, 5, 3, 1, 6, 8, 7, 4, 2},
	// 	{8, 6, 2, 7, 3, 4, 9, 5, 1},
	// 	{4, 1, 7, 9, 5, 2, 8, 3, 6},
	// 	{7, 4, 6, 8, 9, 3, 1, 2, 5},
	// 	{2, 8, 1, 6, 4, 5, 3, 9, 7},
	// 	{3, 9, 5, 2, 7, 1, 4, 6, 8},
	// 	{1, 3, 8, 5, 2, 9, 6, 7, 4},
	// 	{5, 7, 4, 3, 8, 6, 2, 1, 9},
	// 	{6, 2, 9, 4, 1, 7, 5, 8, 3},
	// }
	fmt.Println(SolveSudoku(input))
}
