package main

import (
	"fmt"
)

func isValid(board [][]int, row, col, num int) bool {
	// Check the line
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// Check the column
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check the box
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func solved(board [][]int) bool {
	// Check if each row, column, and 3x3 subgrid contains numbers from 1 to 9 without repetition
	return checkRows(board) && checkCols(board) && checkSubgrids(board)
}

func checkRows(board [][]int) bool {
	for i := 0; i < 9; i++ {
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != 0 && seen[num] {
				return false // Duplicate in row
			}
			seen[num] = true
		}
	}
	return true
}

func checkCols(board [][]int) bool {
	for j := 0; j < 9; j++ {
		seen := make(map[int]bool)
		for i := 0; i < 9; i++ {
			num := board[i][j]
			if num != 0 && seen[num] {
				return false // Duplicate in column
			}
			seen[num] = true
		}
	}
	return true
}

func checkSubgrids(board [][]int) bool {
	for startRow := 0; startRow < 9; startRow += 3 {
		for startCol := 0; startCol < 9; startCol += 3 {
			if !checkSubgrid(board, startRow, startCol) {
				return false
			}
		}
	}
	return true
}

func checkSubgrid(board [][]int, startRow, startCol int) bool {
	seen := make(map[int]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := board[i+startRow][j+startCol]
			if num != 0 && seen[num] {
				return false // Duplicate in subgrid
			}
			seen[num] = true
		}
	}
	return true
}

func deepCopy(board [][]int) [][]int {
	copyBoard := make([][]int, len(board))
	for i := range board {
		copyBoard[i] = make([]int, len(board[i]))
		copy(copyBoard[i], board[i])
	}
	return copyBoard
}

func SolveSudoku(board [][]int) [][]int {
	fmt.Println("Entering SolveSudoku")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(board, i, j, num) {
						boardCopy := deepCopy(board) // Create a copy of the board
						boardCopy[i][j] = num

						fmt.Printf("Trying %d at (%d, %d)\n", num, i, j)
						fmt.Println("Current board:")
						printBoard(boardCopy)

						if solved(boardCopy) {
							fmt.Println("Board solved:")
							printBoard(boardCopy)
							return boardCopy
						} else {
							fmt.Println("Recursive call...")
							result := SolveSudoku(boardCopy)
							fmt.Println("Back from recursive call")
							printBoard(result)
						}
					}
				}
				fmt.Printf("No valid number found at (%d, %d)\n", i, j)
				return board // Return the original board if no valid number is found
			}
		}
	}
	fmt.Println("Exiting SolveSudoku")
	return board
}

func printBoard(board [][]int) {
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println()
}
