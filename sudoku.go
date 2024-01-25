package main

import "fmt"

// SolveSudoku(board [][]int) [][]int is the entry point for the Sudoku solver.
// It calls the solve function to solve the puzzle from the first cell.
//
// Parameters
// board [][]int - the Sudoku board
//
// Returns the solved Sudoku board - which is modified in-place
func SolveSudoku(board [][]int) [][]int {
	if !solve(board, 0) {
		fmt.Println("unsolvable puzzle")
	}
	return board;
}


// solve(board [][]int, s int) produces the solution of the Sudoku puzzle using backtracking.
// It fills in the board cell by cell using each possible number and checks for the validity of each placement.
//
// Parameters
// board [][]int - the Sudoku board
// s int - the linear index of the current cell
//
// Returns true if the puzzle is solved, false otherwise
func solve(board [][]int, s int) bool {
	// If we have filled all cells, the puzzle is solved.
	if s == 81 {
		return true
	}

	// Calculate the row and column indices from the linear index.
	i := s / 9
	j := s % 9

	// If the current cell is not empty, move on to the next cell.
	if board[i][j] != 0 {
		return solve(board, s+1)
	}

	// Try all possible digits 1-9 in current cell
	for c := 1; c <= 9; c++ {
		if isValid(board, i, j, c) {
			// Place the digit and recursively attempt to solve the puzzle.
			board[i][j] = c
			if solve(board, s+1) {
				return true
			}
			// If placing the digit does not lead to a solution, backtrack and try the next digit.
			board[i][j] = 0
		}
	}

	// No valid digit found for the current cell, trigger backtracking.
	return false
}


// isValid(board [][]int, row, col, num int) checks if the placement of num in the cell (row, col) 
// results in a valid Sudoku puzzle.
// 
// Parameters
// board [][]int - the Sudoku board
// row int - the row index of the cell
// col int - the column index of the cell
// num int - the digit to be placed in the cell
//
// Returns true if the placement is valid, false otherwise
func isValid(board [][]int, row, col, num int) bool {
	// Check the current row and column for the presence of the digit.
	for i := 0; i < 9; i++ {
		if board[i][col] == num || board[row][i] == num {
			return false
		}
	}

	// Check the 3x3 subgrid for the presence of the digit.
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	// The placement is valid.
	return true
}