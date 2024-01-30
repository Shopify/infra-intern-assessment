package main

func nextSquare(row int, col int) (int, int) {
	// Determines the next square to be filled in the board.

	pos := row*9 + col + 1
	if pos == 81 {
		// All squares are filled. There is no next square.
		return -1, -1
	}
	return pos / 9, pos % 9
}

func checkCandidate(board [][]int, row int, col int, candidate int) bool {
	// Check if the candidate is valid for the given board and square.

	// Check if the candidate is valid for the row.
	for i := 0; i < 9; i++ {
		if board[row][i] == candidate {
			return false
		}
	}

	// Check if the candidate is valid for the column.
	for i := 0; i < 9; i++ {
		if board[i][col] == candidate {
			return false
		}
	}

	// Check if the candidate is valid for the 3x3 square.
	for i := 0; i < 9; i++ {
		if board[row/3*3+i/3][col/3*3+i%3] == candidate {
			return false
		}
	}

	return true
}

func solveBoard(board [][]int, row int, col int) bool {
	// Solve the board recursively.
	
	// Find the next square to be filled.
	for board[row][col] != 0 {
		row, col = nextSquare(row, col)
		if row == -1 {
			// All squares are filled.
			return true
		}
	}

	// Check all candidates for the square on the board.
	for candidate := 1; candidate <= 9; candidate++ {
		if checkCandidate(board, row, col, candidate) {
			// Candidate is valid.
			board[row][col] = candidate
			if solveBoard(board, row, col) {
				return true
			}
			board[row][col] = 0
		}
	}

	return false
}

func SolveSudoku(board [][]int) [][]int {
	// Copy board to avoid modifying the original board.
	new_board := make([][]int, 9)
	for i := range new_board {
		new_board[i] = make([]int, 9)
		copy(new_board[i], board[i])
	}

	// Call the recursive function.
	if solveBoard(new_board, 0, 0) {
		return new_board
	} else {
		panic("No solution.")
	}
}
