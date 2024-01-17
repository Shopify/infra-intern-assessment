package main

// To validate if the number can be placed in the cell
func valid(row, col, num int, arr [][]int) bool {
	// check if number exist on the same row and col
	for i := 0; i < 9; i++ {
		if arr[row][i] == num || arr[i][col] == num {
			return false
		}
	}

	// get the first index of the current block
	r := (row / 3) * 3
	c := (col / 3) * 3

	// check if the number exist in the same block (3x3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if arr[r+i][c+j] == num {
				return false
			}
		}
	}
	return true
}
// check if it contaitns 0's if it does that means solution is not complete
func containsZeros(arr [][]int) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if arr[r][c] == 0 {
				return true
			}
		}
	}
	return false
}

func SolveSudoku(arr [][]int) [][]int {
	// for each cell, check if it's empty and if so, attempt to fill it
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if arr[r][c] == 0 {
				// try every possible number
				for n := 1; n <= 9; n++ {
					if valid(r, c, n, arr) {
						arr[r][c] = n
						result := SolveSudoku(arr)
						// Check if a solution was found by looking for any zeros left
						if !containsZeros(result) {
							return result
						}
						arr[r][c] = 0 // backtrack
					}
				}
				return arr // no valid number found for this cell, backtrack
			}
		}
	}
	return arr // return the final array, solved or not
}
