// Code Written By - Bhavesh Motiramani
// Modified on: 29th January 2024

// Sudoku Rules:
// 	1. Each row must contain the digits 1-9 without repetition.
//  2. Each column must contain the digits 1-9 without repetition.
//  3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

// This problem can be solved in the following ways:
//
// 1. Brute force approach:
//    The brute force approach of solving this problem involves traversing the matrix and trying out all possible combinations of numbers in empty cells until a valid solution is found.
//    Time Complexity: O(9^(n*n))
//    Space Complexity: O(n*n)
//
// 2. Backtracking:
// I have used backtracking to solve this problem, which is a more efficient solution than the brute force approach.
//
// Intuition:
// The problem can be divided into subproblems.
//
// Approach:
// When an empty cell is found, attempt to put all available numbers (1 – 9) in that cell. If it's not valid (doesn't satisfy three rules), try the next number; else, call the function recursively with updated state.
// Recursively fill the remaining cells. If unable to fill cells with numbers (1-9), return nil.
// This signals to the parent call that the board is not valid, allowing it to try the next possible number.
//
// A more efficient solution would be to solve this using Bits as demonstrated here (https://www.facebook.com/photo/?fbid=905828379479869&set=a.344710778924968.83425.125845680811480), planned for future study. For simplicity, a backtracking solution is used.

package main

import "fmt"

// SolveSudoku finds a solution to the given Sudoku board using backtracking.
// It modifies the input board in-place.
// Parameters:
// - board: The Sudoku board represented as a 2D array.
// Returns:
// - [][]int: The solved Sudoku board. Returns nil if no solution is possible.
// Time Complexity: O(9^(n*n)) where n is the size of the board (in this case, n = 9)
// Space Complexity: O(1) since we are refilling the given board itself; there is no extra space.

func SolveSudoku(board [][]int) [][]int {
	// Traversing the matrix to find the first cell which has 0 value.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// If the cell is empty: try to fill it with values from [1-9].
			if board[i][j] == 0 {
				for ele := 1; ele <= 9; ele++ {
					// Check if it's a valid value to be placed in the cell.
					if isValid(board, i, j, ele) {
						board[i][j] = ele
						// Recursively call to fill the board.
						// If every recursive call returns true, that means it's a valid sudoku board.
						// We get one valid solution.
						solvedBoard := SolveSudoku(board)
						if solvedBoard != nil {
							return solvedBoard
						}
						// When we don't get a valid solution to a sub-problem, remove it from the solution.
						board[i][j] = 0
					}
				}
				// When we cannot place any value in a particular cell, return nil.
				return nil
			}
		}
	}
	// Return the solved board.
	return board
}

// isValid checks if placing a value 'val' at position (row, col) in the board is valid.
// It checks for the validity of the value in the corresponding row, column, and subgrid.
// Parameters:
// - board: The Sudoku board represented as a 2D array.
// - row: The row index of the position.
// - col: The column index of the position.
// - val: The value to be checked for validity.
// Returns:
// - bool: True if the placement is valid, false otherwise.
func isValid(board [][]int, row, col, val int) bool {
	for i := 0; i < 9; i++ {
		// Check for a particular column, row value changes.
		if board[i][col] == val {
			return false
		}
		// Check for a particular row, column value changes.
		if board[row][i] == val {
			return false
		}
		// Efficient way to check for grid in a single for loop.
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == val {
			return false
		}
	}
	return true
}

func main() {
	board := [][]int{
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

	solvedBoard := SolveSudoku(board)

	for i := 0; i < 9; i++ {
		fmt.Println(solvedBoard[i])
	}
}
