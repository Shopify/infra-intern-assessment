package main

import "fmt"

// BoardSize represents the size of the Sudoku board.
const BoardSize = 9

// SubgridSize represents the size of the subgrids within the Sudoku board.
const SubgridSize = 3

// EmptyCellValue represents the value of an empty cell in the Sudoku board.
const EmptyCellValue = 0

// isValidBoard checks if the input Sudoku board is valid.
func isValidBoard(board [][]int) bool {
	if len(board) != BoardSize || len(board[0]) != BoardSize {
		return false
	}

	for _, row := range board {
		for _, num := range row {
			if num < 0 || num > BoardSize {
				return false
			}
		}
	}
	return true
}

// isPlacementValid checks if it's safe to place a number in a specific cell.
func isPlacementValid(board [][]int, row, col, num int) bool {
	for x := 0; x < BoardSize; x++ {
		if board[row][x] == num || board[x][col] == num {
			return false
		}
	}

	startRow, startCol := row-row%SubgridSize, col-col%SubgridSize
	for i := 0; i < SubgridSize; i++ {
		for j := 0; j < SubgridSize; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

// findEmptyCell finds the first empty cell in the board.
func findEmptyCell(board [][]int) (int, int, bool) {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if board[i][j] == EmptyCellValue {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

// placeNumber places a number in a specific cell on the board.
func placeNumber(board [][]int, row, col, num int) {
	board[row][col] = num
}

// resetCell resets a cell to 0 on the board.
func resetCell(board [][]int, row, col int) {
	board[row][col] = EmptyCellValue
}

// SolveSudoku is a function that solves the Sudoku puzzle using a recursive backtracking algorithm.
// It takes the Sudoku board as input and returns the solved board. If the input board is invalid or
// there is no solution, it returns nil.
func SolveSudoku(board [][]int) [][]int {
	var solve func() bool
	solve = func() bool {
		row, col, isEmpty := findEmptyCell(board)
		if !isEmpty {
			return true
		}
		for num := 1; num <= BoardSize; num++ {
			if isPlacementValid(board, row, col, num) {
				placeNumber(board, row, col, num)

				if solve() {
					return true
				}

				resetCell(board, row, col)
			}
		}
		return false
	}

	if !isValidBoard(board) {
		return nil
	}

	if solve() {
		return board
	}

	return nil
}

// printBoard prints the Sudoku board in a readable format.
func printBoard(board [][]int) {
	fmt.Println("[")
	for i, row := range board {
		fmt.Print("  [")
		for j, num := range row {
			fmt.Printf("%d", num)
			if j < len(row)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("]")
		if i < len(board)-1 {
			fmt.Print(",")
		}
		fmt.Println()
	}
	fmt.Println("]")
}

func main() {
	puzzle := [][]int{
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

	solved := SolveSudoku(puzzle)
	printBoard(solved)
}
