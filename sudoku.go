package main

import ("fmt")

// SolveSudoku solves the Sudoku puzzle using backtracking
func SolveSudoku(puzzle [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if puzzle[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isSafe(puzzle, row, col, num) {
						puzzle[row][col] = num
						if SolveSudoku(puzzle) {
							return true
						}
						puzzle[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// isSafe checks if it's safe to place 'num' at puzzle[row][col]
func isSafe(puzzle [][]int, row, col, num int) bool {
	// Check if 'num' is not present in the current row and column
	for i := 0; i < 9; i++ {
		if puzzle[row][i] == num || puzzle[i][col] == num {
			return false
		}
	}

	// Check if 'num' is not present in the current 3x3 subgrid
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if puzzle[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

// Modify puzzleResp.Puzzle to compute the solution
func computeSolution(puzzle [][]int) {
	if SolveSudoku(puzzle) {
		fmt.Println("Computed Solution:")
		printPuzzle(puzzle)
	} else {
		fmt.Println("No solution found.")
	}
}
