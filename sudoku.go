package main

func SolveSudoku(board [9][9]int) [9][9]int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for digit := 9; digit >= 1; digit-- {
					board[i][j] = digit
					if CheckBoardValidity(board) {
						if IsBoardComplete(board) {
							return board
						} else {
							SolveSudoku(board)
						}
					}
					board[i][j] = 0
				}
				return [9][9]int{} // If no solution found, return an empty board
			}
		}
	}
	return board
}

func CheckBoardValidity(board [9][9]int) bool {
	// Check duplicates by row
	for row := 0; row < 9; row++ {
		rowMap := make(map[int]int)
		for col := 0; col < 9; col++ {
			if rowMap[board[row][col]] == 1 {
				return false
			} else {
				rowMap[board[row][col]] = 1
			}
		}
	}

	// Check duplicates by column
	for col := 0; col < 9; col++ {
		colMap := make(map[int]int)
		for row := 0; row < 9; row++ {
			if colMap[board[row][col]] == 1 {
				return false
			} else {
				colMap[board[row][col]] = 1
			}
		}
	}

	// Check duplicates in 3x3 sections
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			threeByThreeMap := make(map[int]int)
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					if board[row][col] != 0 {
						if threeByThreeMap[board[row][col]] == 1 {
							return false
						} else {
							threeByThreeMap[board[row][col]] = 1
						}
					}
				}
			}
		}
	}

	return true
}

func IsBoardComplete(board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
