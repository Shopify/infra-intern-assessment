package main

const SIZE = 9

// Bitmasks for each row/column/box
var row, col, box [SIZE]int

// SolveSudoku Creates a deep copy of the grid, initializes the board, and starts solving the Sudoku puzzle.
func SolveSudoku(grid [][]int) [][]int {
	res := make([][]int, len(grid))
	for row := range grid {
		res[row] = make([]int, len(grid[row]))
		copy(res[row], grid[row])
	}
	initializeBoard(res)
	found := backtrack(res, 0, 0)
	if found {
		return res
	} else {
		return nil
	}
}

func backtrack(grid [][]int, i, j int) bool {
	// Base case: If reached the end of the grid, the Sudoku is solved.
	if i == SIZE-1 && j == SIZE {
		return true
	}
	// Move to the next row when reaching the end of a row.
	if j == SIZE {
		j = 0
		i++
	}
	// If the cell is already filled, move to the next cell.
	if grid[i][j] > 0 {
		return backtrack(grid, i, j+1)
	}
	// Try filling the cell with numbers 1 to SIZE.
	for currVal := 1; currVal <= SIZE; currVal++ {
		if isSafe(i, j, currVal) {
			// Update masks and fill the cell.
			grid[i][j] = currVal
			updateMasks(i, j, currVal, true)

			// Continue solving the rest of the puzzle.
			if backtrack(grid, i, j+1) {
				return true
			}

			// If the current placement doesn't lead to a solution, backtrack.
			updateMasks(i, j, currVal, false)
		}

		// Reset the cell if the current number doesn't lead to a solution.
		grid[i][j] = 0
	}

	// No solution found for the given board.
	return false
}

// updateMasks updates the masks (row, col, box) based on the given parameters.
// It either sets or unsets the corresponding bit in each mask based on the 'set' parameter.
// - set: A boolean flag indicating whether to set (true) or unset (false) the bit in the masks.
func updateMasks(i, j, currVal int, set bool) {
	mask := 1 << currVal
	boxIdx := i/3*3 + j/3
	if set {
		row[i] |= mask
		col[j] |= mask
		box[boxIdx] |= mask
	} else {
		row[i] &= ^mask
		col[j] &= ^mask
		box[boxIdx] &= ^mask
	}
}

// checks if placing 'number' at cell (i, j) is safe based on current mask state.
func isSafe(i, j, number int) bool {
	mask := 1 << number
	return (row[i]&mask) == 0 &&
		(col[j]&mask) == 0 &&
		(box[i/3*3+j/3]&mask) == 0
}

// initializes the masks (row, col, box) based on the pre-filled grid values.
func initializeBoard(grid [][]int) {
	row = [SIZE]int{}
	col = [SIZE]int{}
	box = [SIZE]int{}
	for i, rowValues := range grid {
		for j, value := range rowValues {
			mask := 1 << value
			row[i] |= mask
			col[j] |= mask
			box[i/3*3+j/3] |= mask
		}
	}
}
