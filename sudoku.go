package main

import "fmt"

type SudokuBoard [][]int

// Checks whether a given value placement is valid (that is, the same number
// exists in neither the block, column, or row).
// row and int numbers start from 1 and in the top left corner, moving down and right.
func isValidCell(row int, col int, board *SudokuBoard) bool {
	//fmt.Print("DEBUG: row = ", row, ", col = ", col, ", val = ", (*board)[row][col], "\n")
	// Check row and col
	for i := 0; i < 9; i++ {
		//fmt.Print("DEBUG: i = ", i, "\n")
		//fmt.Print("DEBUG: board[row][col] = ", (*board)[row][col], "\n")
		//fmt.Print("DEBUG: board[row][i] = ", (*board)[row][i], "\n")
		//fmt.Print("DEBUG: board[i][col] = ", (*board)[i][col], "\n")
		if (col != i && (*board)[row][col] == (*board)[row][i]) ||
			(row != i && (*board)[row][col] == (*board)[i][col]) {
			return false
		}
	}
	//fmt.Print("DEBUG : MADE IT HERE 1\n")
	// Check block
	// Determine the block indices using some simple math
	block_row := row / 3
	block_col := col / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//fmt.Print("DEBUG: i = ", i, ", j = ", j, "\n")
			if (i+block_row*3 != row && j+block_col*3 != col) &&
				(*board)[i+block_row*3][j+block_col*3] == (*board)[row][col] {
				return false
			}
		}
	}
	//fmt.Print("DEBUG : MADE IT HERE 2\n")
	return true
}

// Use a simple backtracking algorithm. Assume there is a solution.
func SolveSudoku(board SudokuBoard) SudokuBoard {
	// Maintain a stack of cells that represent our "path" to filling the board.
	var cell_stack *Stack
	cell_stack = new(Stack)
	// The top cell in the stack is the cell we filled the most recently.
	base := 1
	for i := 0; i < 9; i++ { // go row-by-row
		// Search for an empty cell
		for j := 0; j < 9; j++ {
			//fmt.Print("DEBUG: i = ", i, ", j = ", j, "-------------------\n")
			if board[i][j] == 0 { // empty cell
				valid := false
				for val := base; val <= 9; val++ {
					board[i][j] = val
					//fmt.Print("DEBUG: board[i][j] = ", board[i][j], "\n")
					valid = isValidCell(i, j, &board)

					if valid { // valid value found

						break
					}
				}
				//fmt.Print("DEBUG : valid = ", valid, "\n")
				if valid {

					// add to stack
					cell := Coords{row: i, col: j}
					cell_stack.Push(cell)
					base = 1
					//fmt.Print(board)
					//fmt.Print("\nVALID, PUSHED ONTO STACK\n")
				} else {
					//fmt.Print(board)
					//fmt.Print("\nINVALID, BACKTRACKING\n")
					// If none of the values are valid, backtrack using our stack of cells
					board[i][j] = 0
					cell := cell_stack.Pop()
					//fmt.Print("cell.row: ", cell.row, " cell.col: ", cell.col, "\n")

					base = board[cell.row][cell.col] + 1
					i = cell.row
					j = cell.col - 1
					board[cell.row][cell.col] = 0 // EXPERIMENTAL
					//fmt.Print("i = ", i, " j = ", j, "\n")
				}
			}
		}
	}
	fmt.Print(board)
	return board
}
