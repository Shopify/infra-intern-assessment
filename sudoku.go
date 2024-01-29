package main

// Checks whether a given value placement is valid (that is, the same number
// exists in neither the block, column, or row).
// row and int numbers start from 1 and in the top left corner, moving down and right.
func isValidCell(row int, col int, board *[][]int) bool {
	// Check row and col
	for i := 0; i < 9; i++ {
		if (col != i && (*board)[row][col] == (*board)[row][i]) ||
			(row != i && (*board)[row][col] == (*board)[i][col]) {
			return false
		}
	}
	// Check block
	// Determine the block indices using some simple math
	block_row := row / 3
	block_col := col / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (i+block_row*3 != row && j+block_col*3 != col) &&
				(*board)[i+block_row*3][j+block_col*3] == (*board)[row][col] {
				return false
			}
		}
	}
	return true
}

// Use a simple backtracking algorithm. Assume there is a solution.
func SolveSudoku(board [][]int) [][]int {
	// Maintain a stack of cells that represent our "path" to filling the board.
	var cell_stack *Stack
	cell_stack = new(Stack)
	// The top cell in the stack is the cell we filled the most recently.
	base := 1
	for i := 0; i < 9; i++ { // go row-by-row
		// Search for an empty cell
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 { // empty cell
				valid := false
				for val := base; val <= 9; val++ {
					board[i][j] = val
					valid = isValidCell(i, j, &board)
					if valid { // valid value found
						break
					}
				}
				if valid {
					// add to stack
					cell := Coords{row: i, col: j}
					cell_stack.Push(cell)
					base = 1
				} else {
					// If none of the values are valid, backtrack using our stack of cells
					board[i][j] = 0
					cell := cell_stack.Pop()

					base = board[cell.row][cell.col] + 1
					i = cell.row
					j = cell.col - 1
					board[cell.row][cell.col] = 0 // Reset the previous cell on our stack
				}
			}
		}
	}
	return board
}
