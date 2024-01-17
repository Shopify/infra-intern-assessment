package main

// dfs flattens the 2D board array into a 1D array and recursively fills in numbers.
// n: The index in the flattened 1D array.
// rows: The set of numbers already used in each row when examining position n.
// cols: The set of numbers already used in each column when examining position n.
// grids: The set of numbers already used in each grid (3x3 grid) when examining position n.
func dfs(board [][]int, n int, rows []int, cols []int, grids []int) bool {
	// Recursive end condition: if n is equal to the total number of cells, all positions are filled.
	if n == len(board)*len(board) {
		return true
	}
	// Locate the row and column based on n.
	var i, j = n / len(board), n % len(board)
	// If the current position is already filled, skip to the next position.
	if board[i][j] != 0 {
		return dfs(board, n+1, rows, cols, grids)
	}
	// Attempt to fill the current position with numbers 1 to 9.
	for k := 1; k <= 9; k++ {
		var bit = 1 << k
		// Check if the number is not yet used in the corresponding row, column, and grid.
		if unused := rows[i]&bit == 0 && cols[j]&bit == 0 && grids[i/3*3+j/3]&bit == 0; unused {
			// Mark the number as used in the row, column, and grid.
			rows[i] |= bit
			cols[j] |= bit
			grids[i/3*3+j/3] |= bit
			// Continue to fill the next position.
			if dfs(board, n+1, rows, cols, grids) {
				// If the recursion returns true, it means all subsequent positions are correctly filled,
				// so we can fill the current position.
				board[i][j] = k
				return true
			} else {
				// If the recursion returns false, it means subsequent positions failed to fill,
				// so the current number is invalid and we need to try the next number.
				// Before trying the next number, we need to unmark the current number in the row, column, and grid.
				rows[i] &= ^bit
				cols[j] &= ^bit
				grids[i/3*3+j/3] &= ^bit
			}
		}
	}
	// If none of the numbers 1 to 9 can be filled in the current position, the Sudoku puzzle is unsolvable at this configuration.
	return false
}
