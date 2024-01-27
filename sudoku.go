package main

import (
	"fmt"
)

// TODO:
// 	- document functions
// 	- write comments explaining code

// SolveSudoku takes a Sudoku puzzle and returns the solved puzzle
func SolveSudoku(puzzle [][]int) [][]int {
	// The Sudoku puzzle is transformed into convenient data structures (as seen below) to make solving easier

	// Contains positions of all numbers 1-9, in the following format:
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
	var pos map[int][][]int

	// Count of the number of remaining occurrences of each number, in the following format:
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
	var remaining map[int]int

	// Graph representing the potential positions for each number, stored in the following format:
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
	var graph map[int]map[int][]int

	// pos, remaining = populatePosAndRemaining(puzzle)

	// TODO: Sort the remaining map in ascending order by value
	// We want to fill in the numbers with the least remaining occurrences first,
	// which will make the algorithm more efficient

	// graph = populateGraph(puzzle)

	printPuzzle(puzzle)

	solvedPuzzle := fillPuzzle(puzzle, pos, remaining, graph)

	printPuzzle(solvedPuzzle)
	
	return solvedPuzzle
}

// fillPuzzle fills in the Sudoku puzzle with the correct numbers using a backtracking algorithm
// with crosshatching
func fillPuzzle(puzzle [][]int, pos map[int][][]int, remaining map[int]int,
				graph map[int]map[int][]int) [][]int {
	// Add logic
	return puzzle
}

// validPlacement checks if the number in the position puzzle[row][col] is in a valid spot
// according to the rules of Sudoku
func validPlacement(puzzle [][]int, row int, col int) bool {
	num := puzzle[row][col]
	for i := 0; i < len(puzzle); i++ {
		// Check if any numbers in the same row equal num
		if puzzle[row][i] == num && i != col {
			return false
		}
		// Check if any numbers in the same column equal num
		if puzzle[i][col] == num && i != row {
			return false
		}
	}

	// Check if there are any duplicate numbers in the 3x3 box
	boxRowStart := row / 3 * 3
	boxColStart := col / 3 * 3

	for r := boxRowStart; r < boxRowStart + 3; r++ {
		for c := boxColStart; c < boxColStart + 3; c++ {
			if puzzle[r][c] == num && r != row && c != col {
				return false
			}
		}
	}

	return true
}

// createPosAndRem returns maps which contain the positions and remaining counts of each number in
// the puzzle
func createPosAndRem(puzzle [][]int) (map[int][][]int, map[int]int) {
	pos := make(map[int][][]int)
	rem := make(map[int]int)

	// Initialize pos and rem
	for i := 1; i <= len(puzzle); i++ {
		pos[i] = [][]int{}
		rem[i] = 9
	}

	for r := 0; r < len(puzzle); r++ {
		for c := 0; c < len(puzzle); c++ {
			if puzzle[r][c] != 0 {  // Ignore 0s
				// Add the position of the number to the pos map
				pos[puzzle[r][c]] = append(pos[puzzle[r][c]], []int{r, c})
				// Decrement the remaining count of the number
				rem[puzzle[r][c]]--
			}
		}
	}

	return pos, rem
}

// createGraph returns a map which contains potential placements for each number
func createGraph(puzzle [][]int, pos map[int][][]int) map[int]map[int][]int {
	graph := make(map[int]map[int][]int)

	// Iterate over each number and its positions in the puzzle
	for k, v := range pos {
		graph[k] = make(map[int][]int)

		// Initialize slices to keep track of which rows and columns already have the number k
		rows := make([]bool, len(puzzle))
		cols := make([]bool, len(puzzle))

		// Mark the rows and columns which already have the number k
		for _, coord := range v {
			rows[coord[0]] = true
			cols[coord[1]] = true
		}

		// Iterate over each cell in the puzzle
		for r := 0; r < len(puzzle); r++ {
			for c := 0; c < len(puzzle); c++ {
                // If the current row and column don't have the number and the cell is empty
                // (indicated by 0), store the cell in the graph
				if !rows[r] && !cols[c] && puzzle[r][c] == 0 {
					graph[k][r] = append(graph[k][r], c)
				}
			}
		}
	}

	return graph
}

// printPuzzle prints a Sudoku puzzle
func printPuzzle(puzzle [][]int) {
	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[row]); col++ {
			fmt.Printf("%d ", puzzle[row][col])
		}
		fmt.Println()
	}
}

func main() {
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

	SolveSudoku(puzzle)
}