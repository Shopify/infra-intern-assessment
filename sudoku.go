package main

// check if there are duplicate numbers within a row
func checkRow(game [][]int) bool {
	for i := 0; i < 9; i++ {
		visited := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if visited[game[i][j]] == true {
				return false
			} else {
				if game[i][j] != 0 {
					visited[game[i][j]] = true
				}
			}
		}
	}
	return true
}

// check if there are duplicate numbers within a column
func checkCol(game [][]int) bool {
	for j := 0; j < 9; j++ {
		visited := make(map[int]bool)
		for i := 0; i < 9; i++ {
			if visited[game[i][j]] == true {
				return false
			} else {
				if game[i][j] != 0 {
					visited[game[i][j]] = true
				}
			}
		}
	}
	return true
}

// check if there are duplicate numbers within 3x3 grid
func checkGrid(game [][]int) bool {
	for x := 0; x < 9; x += 3 {
		for y := 0; y < 9; y += 3 {
			visited := make(map[int]bool)
			for i := x; i < x+3; i++ {
				for j := y; j < y+3; j++ {
					if visited[game[i][j]] == true {
						return false
					} else {
						if game[i][j] != 0 {
							visited[game[i][j]] = true
						}
					}
				}
			}
		}
	}
	return true
}

func checkGame(game [][]int) bool {
	return checkGrid(game) && checkCol(game) && checkRow(game)
}

// solve sudoku by backtracking: place a number, check validity of the game after each placement
// if the placement makes the game invalid, undo it and try another number, repeat till the board is filled and return
// the solved game, return nil if there's no solution
func SolveSudoku(game [][]int) [][]int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if game[i][j] == 0 {
				for n := 1; n <= 9; n++ {
					game[i][j] = n
					if checkGame(game) && SolveSudoku(game) != nil {
						return game
					} else {
						game[i][j] = 0
					}
				}
				return nil
			}
		}
	}
	return game
}

