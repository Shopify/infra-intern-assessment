package main

import (
	"fmt"
)

// TODO:
// 	- document functions
// 	- write comments explaining code

var puzzle = [][]int{
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

// Contains positions of all numbers 1-9
// Format:
// {
// 	number: [[row, col], [row, col], ...],
// 	...
// }
// Example:
// {
// 	1: [[0, 0], [3, 1], ...],
// 	2: [[4, 5], [6, 6], ...],
// 	...
// }
var pos = map[int][][]int{}

// Count of the number of remaining occurrences of each number
// Format:
// {
// 	number: remaining count,
// 	...
// }
// Example:
// {
// 	1: 7,
// 	2: 5,
// 	...
// }
var remaining = map[int]int{}

// Graph representing the valid positions for each number
// Format:
// {
// 	number: {
// 		row_i: [col_i, ...],
// 		row_j: [col_j, ...],
// 		...
// 	},
// 	...
// }
// Example:
// {
// 	1: {
// 		0: [0, 1, 2, 6, 7],
// 		3: [0, 1, 2, 6, 7],
// 		...
// 	},
// 	...
// }
var graph = map[int]map[int][]int{}

// SolveSudoku takes a 9x9 Sudoku puzzle and returns the solved puzzle using a backtracking algorithm with crosshatching
func SolveSudoku(puzzle [][]int) [][]int {
	printPuzzle(puzzle)

	// Add logic
	
	return solvedPuzzle
}

func printPuzzle(puzzle [][]int) {
	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[row]); col++ {
			fmt.Printf("%d ", puzzle[row][col])
		}
		fmt.Println()
	}
}

func main() {
	SolveSudoku(puzzle)
}