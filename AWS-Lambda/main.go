package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

// SudokuEvent represents the input event for the Lambda function.
type SudokuEvent struct {
	Board [][]int `json:"board"`
}

// SudokuResponse represents the output from the Lambda function.
type SudokuResponse struct {
	SolvedBoard [][]int `json:"solvedBoard,omitempty"`
	Error       string  `json:"error,omitempty"`
}

// CanPlace checks if a number can be placed at the given row and column in the Sudoku board.
func CanPlace(board [][]int, row, col, num int) bool {
	for x := 0; x < 9; x++ {
		// Check if the number is already in the row or column.
		if board[row][x] == num || board[x][col] == num {
			return false
		}
		// Check if the number is in the 3x3 subgrid.
		if board[3*(row/3)+x/3][3*(col/3)+x%3] == num {
			return false
		}
	}
	return true
}

// solve is a recursive helper function that attempts to solve the Sudoku puzzle.
func solve(board [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Look for an empty cell (denoted by 0).
			if board[i][j] == 0 {
				// Try placing numbers 1 through 9 in the empty cell.
				for num := 1; num <= 9; num++ {
					if CanPlace(board, i, j, num) {
						board[i][j] = num
						// Recursively attempt to solve the rest of the board.
						if solve(board) {
							return true
						}
						// Backtrack if placing num doesn't lead to a solution.
						board[i][j] = 0
					}
				}
				// Return false if no number can be placed in the empty cell.
				return false
			}
		}
	}
	// Return true if all cells are filled, meaning the puzzle is solved.
	return true
}

// SolveSudoku attempts to solve the given Sudoku puzzle and returns the solved board.
func SolveSudoku(board [][]int) [][]int {
	if solve(board) {
		// If the puzzle is solvable, return the solved board.
		return board
	}
	// Return nil if no solution exists for the puzzle.
	return nil
}

// Handler is the Lambda function handler.
func Handler(ctx context.Context, event SudokuEvent) (SudokuResponse, error) {
	// Attempt to solve the Sudoku puzzle.
	solved := SolveSudoku(event.Board)
	if solved != nil {
		// If a solution exists, return the solved board.
		return SudokuResponse{SolvedBoard: solved}, nil
	} else {
		// If no solution exists, return an error message.
		return SudokuResponse{Error: "No solution exists"}, nil
	}
}

func main() {
	// Start the Lambda function handler.
	lambda.Start(Handler)
}

