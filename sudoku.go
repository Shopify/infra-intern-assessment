package main

import "fmt"

/**
 * Solves a sudoku puzzle using backtracking
 * @param puzzle The sudoku puzzle to solve
 * @return The solved sudoku puzzle
 **/
func SolveSudoku(puzzle [][]int) [][]int {
	// Iterate over each row
	for row := 0; row < 9; row++ {
		// Iterate over each column
		for col := 0; col < 9; col++ {
			if puzzle[row][col] == 0 {
				for val := 1; val <= 9; val++ {
					if isPossible(row, col, val, puzzle) {
						// We try with this value since it is possible
						puzzle[row][col] = val
						if SolveSudoku(puzzle) != nil {
							return puzzle
						}
						// We backtrack since it is not possible
						puzzle[row][col] = 0
					}
				}
				return nil
			}
		}
	}
	return puzzle
}

/**
 * Checks if a value is possible in a given position.
 * We check if 
 * @param y The y coordinate / row
 * @param x The x coordinate / column
 * @param val The value to check
 * @param puzzle The sudoku puzzle
 * @return True if the value is possible, false otherwise
 **/
func isPossible(row int, col int, val int, puzzle [][]int) bool {

	// Check for row
	for i := 0; i < 9; i++ {
		if puzzle[row][i] == val {
			return false
		}
	}

	// Check for column
	for i := 0; i < 9; i++ {
		if puzzle[i][col] == val {
			return false
		}
	}

	// Check for box
	boxX := (col / 3) * 3
	boxY := (row / 3) * 3
	for i := boxY; i < boxY+3; i++ {
		for j := boxX; j < boxX+3; j++ {
			if puzzle[i][j] == val {
				return false
			}
		}
	}
	return true
}

/**
 * Prints a sudoku puzzle
 * @param puzzle The sudoku puzzle to print
 **/
func print(puzzle [][]int) {
	for _, row := range puzzle {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// The below code can be uncommented to print the solution. Replace the input with the puzzle you want to solve.
/*
func main() {
	input := [][]int {
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

	print(SolveSudoku(input))

}
*/
