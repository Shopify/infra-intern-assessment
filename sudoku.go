package main

import (
	"fmt"
)

func SolveSudoku(grid [][]int) [][]int {
	solvedGrid := CopyGrid(grid)
	return SolveGridNaive(solvedGrid, 0, 0)
}

func PrintSudoku(grid [][]int) {
	for i := range grid {
		fmt.Printf("%v\n", grid[i])
	}
}

func CopyGrid(grid [][]int) [][]int {
	tempGrid := make([][]int, len(grid))
	for i := range grid {
		tempGrid[i] = make([]int, len(grid[i]))
		copy(tempGrid[i], grid[i])
	}
	return tempGrid
}

func SolveGridNaive(grid [][]int, row int, col int) [][]int {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}
	return grid
}