package main

func SolveSudoku(grid [][]int) (solution [][]int) {
	backTrackAlgo(grid, 0, 0)
	return grid
}

func isValidCandidate(grid [][]int, r int, c int, n int) bool {
	// check row & column for duplicates
	currRow := grid[r]
	for i := 0; i < len(currRow); i++ {
		if currRow[i] == n || grid[i][c] == n {
			return false
		}
	}

	// check subgrid for duplicates
	startI := (r / 3) * 3
	startJ := (c / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startI][j+startJ] == n {
				return false
			}
		}
	}
	return true
}

func backTrackAlgo(grid [][]int, r int, c int) bool {
	lenGrid := len(grid)
	if r == lenGrid || c == lenGrid {
		// reached end of grid
		return true
	}

	if grid[r][c] != 0 {
		// cell is already filled with a constant
		newC := (c + 1) % lenGrid                // next cell or cycle back to 0
		newR := r + (c + 1) / lenGrid            // same row unless column cycled back to 0
		return backTrackAlgo(grid, newR, newC)
	}

	for candidate := 1; candidate <= 9; candidate++ {
		if isValidCandidate(grid, r, c, candidate) {
			grid[r][c] = candidate
			newC := (c + 1) % lenGrid            // next cell or cycle back to 0
			newR := r + (c + 1) / lenGrid        // same row unless column cycled back to 0
			if backTrackAlgo(grid, newR, newC) {
				return true
			}
			grid[r][c] = 0
		}
	}
	return false
}
