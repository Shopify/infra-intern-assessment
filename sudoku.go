package main

import (
	"fmt"
	"slices"
)

// solves a sudoku puzzle using optimized backtracking
func SolveSudoku(board [][]int) [][]int {
	solution := runOptimizedBacktracking(board)
	printBoard(solution)
	return solution
}

// recursively solves a sudoku puzzle using backtracking
// uses naked single and hidden single calculation to reduce the number of recursive calls
// see https://www.sudopedia.org/wiki/Naked_Single
// see https://www.sudopedia.org/wiki/Hidden_Single
func runOptimizedBacktracking(board [][]int) [][]int {
	// make a copy of the board so we don't modify the original (to be able to backtrack easily)
	// slightly hurts performance, but allows naked and hidden singles to be solved in place
	copy := copy2DArray(board)

	// solve any naked and hidden singles before starting backtracking
	solveNakedAndHiddenSingles(copy)

	// find the next empty square
	row, col := findEmptySquare(copy)

	// if there are no empty squares, the puzzle is solved
	if row == -1 && col == -1 {
		return copy
	}

	// get the possible numbers that can be placed in the square
	candidates := getCandidates(copy, row, col)

	// try each candidate recursively
	for _, candidate := range candidates {
		// place the candidate in the square
		copy[row][col] = candidate

		// move on to the next square
		solution := runOptimizedBacktracking(copy)

		// if the puzzle is solved, return the solution
		if solution != nil {
			return solution
		}
	}

	// if none of the candidates worked, backtrack
	return nil
}

// solves any naked and hidden singles in the board
// see https://www.sudopedia.org/wiki/Naked_Single
// see https://www.sudopedia.org/wiki/Hidden_Single
func solveNakedAndHiddenSingles(board [][]int) [][]int {
	for {
		// flag to recompute the candidate tensor and solve naked singles
		recompute := false

		// compute the candidate tensor and solve naked singles
		candidateTensor := computeCandidateTensorAndSolveNakedSingles(board)
		// iterate over the board
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				// if the square is already solved, skip
				if board[i][j] != 0 {
					continue
				}

				// check if any of the candidates are hidden singles
				candidates := candidateTensor[i][j]
				for _, candidate := range candidates {
					if isHiddenSingle(candidate, candidateTensor, i, j) {
						// if a hidden single is found, place it in the square
						board[i][j] = candidate
						// recompute the candidate tensor and solve naked singles again
						recompute = true
						break
					}
				}
			}
		}

		if !recompute {
			break
		}
	}

	return board
}

// finds the next empty square in the board
func findEmptySquare(board [][]int) (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

// creates a 3d array of the possible numbers that can be placed in each empty square, and solves any naked singles.
// these functions are combined to avoid having to iterate over the board twice.
func computeCandidateTensorAndSolveNakedSingles(board [][]int) [][][]int {
	// build a 3d array of possible numbers for each square
	candidateTensor := make([][][]int, 9)

	// iterate over the board
	for row := 0; row < 9; row++ {
		candidateTensor[row] = make([][]int, 9)
		for col := 0; col < 9; col++ {
			// if the square is already solved, skip
			if board[row][col] != 0 {
				candidateTensor[row][col] = []int{}
				continue
			}

			// compute the possible numbers that can be placed in the square
			candidates := getCandidates(board, row, col)
			candidateTensor[row][col] = candidates

			// if there is only one possible number, it's a naked single, so place it in the square
			if len(candidates) == 1 {
				board[row][col] = candidates[0]
			}
		}
	}

	return candidateTensor
}

// checks if a number is a "hidden single" for a given square.
// a hidden single is a number that can only be placed in one square in a row, column, or 3x3 block.
// see https://www.sudopedia.org/wiki/Hidden_Single
func isHiddenSingle(candidate int, candidateTensor [][][]int, row int, col int) bool {
	// check if the number is also a candidate for another square in the row, column, or 3x3 block
	for i := 0; i < 9; i++ {
		if i != col && slices.Contains(candidateTensor[row][i], candidate) {
			return false
		}
		if i != row && slices.Contains(candidateTensor[i][col], candidate) {
			return false
		}
		blockRow := 3*(row/3) + i/3
		blockCol := 3*(col/3) + i%3
		if !((blockRow) == row && (blockCol) == col) && slices.Contains(candidateTensor[blockRow][blockCol], candidate) {
			return false
		}
	}

	return true
}

// computes the possible numbers that can be placed in a given square
func getCandidates(board [][]int, row int, col int) []int {
	candidates := []int{}

	// try numbers 1-9, and add them to the list of candidates if they are legal
	for num := 1; num <= 9; num++ {
		if isLegalMove(board, row, col, num) {
			candidates = append(candidates, num)
		}
	}

	return candidates
}

// checks if a number can be placed in a given square while adhering to the rules of sudoku
func isLegalMove(board [][]int, row int, col int, num int) bool {
	// check if the number is already in the row, column, or 3x3 block
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}

		if board[i][col] == num {
			return false
		}

		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}

	return true
}

// prints a 2D array
func printBoard(board [][]int) {
	for i := range board {
		fmt.Println(board[i])
	}
	fmt.Println()
}

// makes a deep copy of a 2D array
func copy2DArray(arr [][]int) [][]int {
	newArr := make([][]int, len(arr))

	for i := range arr {
		newArr[i] = make([]int, len(arr[i]))
		copy(newArr[i], arr[i])
	}

	return newArr
}
