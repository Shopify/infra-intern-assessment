package main

func CheckSubSudoku(board [][]int, row, col int, availableNums map[int]bool) {
	var top, bottom, left, right int

	if col <= 2 {
		left, right = 0, 2
	} else if col <= 5 {
		left, right = 3, 5
	} else {
		left, right = 6, 8
	}

	if row <= 2 {
		top, bottom = 0, 2
	} else if row <= 5 {
		top, bottom = 3, 5
	} else {
		top, bottom = 6, 8
	}

	for i := top; i <= bottom; i++ {
		for j := left; j <= right; j++ {
			if board[i][j] != 0 {
				delete(availableNums, board[i][j])
			}
		}
	}
}

func GetAvailableNumbers(board [][]int, row, col int) map[int]bool {
	availableNums := make(map[int]bool)
	for i := 1; i <= 9; i++ {
		availableNums[i] = true
	}

	for i := 0; i < len(board[0]); i++ {
		if board[row][i] != 0 {
			delete(availableNums, board[row][i])
		}
	}

	for j := 0; j < len(board); j++ {
		if board[j][col] != 0 {
			delete(availableNums, board[j][col])
		}
	}

	CheckSubSudoku(board, row, col, availableNums)
	return availableNums
}

func FillBoard(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				availableNumbers := GetAvailableNumbers(board, i, j)
				for num := range availableNumbers {
					board[i][j] = num
					if FillBoard(board) {
						return true
					}
					board[i][j] = 0
				}
				return false
			}
		}
	}
	return true
}

func SolveSudoku(board [][]int) [][]int {
	FillBoard(board)
	return board
}
