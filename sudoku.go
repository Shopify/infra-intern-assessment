package main

import (
	"fmt"
)


const gridSize = 9


func SolveSudoku(grid [][]int) bool {

	row, col := findEmptyCell(grid)

	if row == -1 && col == -1 {
		return true
	}

	for num := 1; num <= gridSize; num++ {
		if isSafe(grid, row, col, num) {
			
			grid[row][col] = num

			
			if SolveSudoku(grid) {
				return true
			}

			grid[row][col] = 0
		}
	}

	return false
}

func findEmptyCell(grid [][]int) (int, int) {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if grid[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

func isSafe(grid [][]int, row, col, num int) bool {
	for i := 0; i < gridSize; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

func PrintGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func main() {
	inputGrid := [][]int{
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

	fmt.Println("Input Sudoku:")
	PrintGrid(inputGrid)

	if SolveSudoku(inputGrid) {
		fmt.Println("\nSolved Sudoku:")
		PrintGrid(inputGrid)
	} else {
		fmt.Println("\nNo solution found.")
	}
}
