package main

import "fmt"

func SolveSudoku(board [][]int) [][]int {
	// Check base case if grid is full (i.e fails to find empty cell)
	row, col := findEmptyCell(board)
	if row == -1 {
		// Print as per instructions
		printBoard(board)
		return board
	}

	for num := 1; num <= 9; num++ {
		// Place num in empty cell if valid
		if isValidNum(board, row, col, num) {
			board[row][col] = num

			// Recursive call to solve remainder of grid
			result := SolveSudoku(board)
			if result != nil {
				return result
			}

			// If the number placement didnt result in an answer, it backtracks, resetting tile to 0
			board[row][col] = 0
		}
	}
	return nil
}

func isValidNum(board [][]int, row int, col int, num int) bool {
	// check row or col
	for i := 0; i < len(board); i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// check subgrid (3x3)
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func findEmptyCell(board [][]int) (int, int) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			if board[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1 // Fails if grid is complete
}

func printBoard(board [][]int) {
	for _, row := range board {
		// Iterate over the columns in each row
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println() // New line after each row
	}
}
