package main

import (
	"fmt"
)

func SolveSudoku(grid [][]int) [][]int {
	solvedGrid := CopyGrid(grid)
	SolveGridNaive(solvedGrid, 0, 0)
	PrintSudoku(solvedGrid)
	return solvedGrid
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

func SolveGridNaive(grid [][]int, row int, col int) bool {
	// move to next row if a row is finished
	if col >= len(grid[row]) {
		row++
		col = 0
		// Solved if passed last row
		if (row >= len(grid)) {
			return true
		}
	}
	// move to next box if box is filled 
	if grid[row][col] != 0 {
		return SolveGridNaive(grid, row, col + 1)
	}
	// try values until a solution is found or exhausted
	for i := 1; i <= 9; i++ {
		if IsBoxValid(grid, row, col, i) {
			grid[row][col] = i
			if SolveGridNaive(grid, row, col + 1) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

func IsBoxValid(grid [][]int, row int, col int, value int) bool {
	// check duplicate in row
	for j := range grid[row] {
		if grid[row][j] == value {
			return false
		}
	}
	// check duplicate in column
	for i := range grid {
		if grid[i][col] == value {
			return false
		}
	}
	// get the top left index of the 3x3 subgrid
	subRow := row - row % 3
	subCol := col - col % 3
	// check duplicate in 3x3 subgrid
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[subRow + i][subCol + j] == value {
				return false
			}
		}
	}
	return true
}