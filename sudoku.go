package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isSafe checks if it's valid to place a number in the specified cell.
// The function ensures that the same number does not appear in the same row,
// column, or 3x3 subgrid.
func isSafe(board [][]int, row, col, num int) bool {

	// Check the row and column
	for x := 0; x < 9; x++ {
		if board[row][x] == num || board[x][col] == num {
			return false
		}
	}

	// Check the 3x3 subgrid
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

// solveSudoku attempts to solve the Sudoku puzzle using a backtracking algorithm.
// It returns true if the puzzle is solved and false otherwise.
func solveSudoku(board [][]int) bool {
	row := -1
	col := -1
	isEmpty := true

	// Search for an empty cell
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				row = i
				col = j
				// Empty cell found
				isEmpty = false 	// Mark that we still have work to do
				break
			}
		}
		if !isEmpty {
			break
		}
	}

	// If no empty cell is found, the puzzle is solved
	if isEmpty {
		return true
	}

	// Try placing numbers 1 through 9 in the current empty cell
	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true       	// Recursively solve the rest of the board
			}
			board[row][col] = 0   	// Backtrack if the guess was incorrect
		}
	}

	// If no number can be placed in the current empty cell, return false
	return false
}

// printBoard outputs the Sudoku grid to the console.
func printBoard(board [][]int) {
	for i := 0; i < 9; i++ {
		for j := 
0; j < 9; j++ {
fmt.Printf("%d ", board[i][j])
}
fmt.Println()
}
}

// This main is the entry point of the program.
func main() {
	reader := bufio.NewReader(os.Stdin)
	board := make([][]int, 9) 		// Initialize a 9x9 grid

	// Read each row of the Sudoku grid from the console
	for i := 0; i < 9; i++ {
		fmt.Printf("Enter row %d (9 numbers separated by spaces): ", i+1)
		rowStr, _ := reader.ReadString('\n')
		rowStr = strings.TrimSpace(rowStr)
		nums := strings.Split(rowStr, " ")

		// Validate the input to ensure there are exactly 9 numbers
		if len(nums) != 9 {
			fmt.Println("Invalid input: each row must have exactly 9 numbers")
			os.Exit(1)
		}

		// Parse each number and place it in the board
		board[i] = make([]int, 9)
		for j, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 0 || num > 9 {
				fmt.Println("Invalid number in input: ", numStr)
				os.Exit(1)
			}
			board[i][j] = num
		}
	}

	// Solve the Sudoku and print the result
	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("No solution exists")
	}
}
