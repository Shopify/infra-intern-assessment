package main

const GRID_SIZE = 9
const SUBGRID_SIZE = 3

// SolveSudoku attempts to solve the Sudoku puzzle using backtracking
// and will succeed since it is guaranteed by the problem statement
// It uses SolveBoard with parameters 0,0 to indicate that it is starting with
// the very first block of the grid

// Runtime of O(n^m) where n is the GRID_SIZE and m is the number of 0s in the grid
func SolveSudoku(grid [][]int) [][]int {
	SolveBoard(grid, 0, 0)
	return grid
}

// SubGridContainsNum checks if the number has already been used in the
// smaller 3x3 subgrid based on the row and col we want to check
func SubGridContainsNum(grid [][]int, row int, col int, num int) bool {

	subgrid_start_row := row - (row % SUBGRID_SIZE)
	subgrid_start_col := col - (col % SUBGRID_SIZE)
	for i := 0; i < SUBGRID_SIZE; i++ {
		for j := 0; j < SUBGRID_SIZE; j++ {
			if num == grid[subgrid_start_row+i][subgrid_start_col+j] {
				return true
			}
		}
	}
	return false
}

// IsValidPosition checks if we are allowed to place a number in the position
// in the grid by checking it against the rules of sudoku
func IsValidPosition(grid [][]int, row int, col int, num int) bool {
	for i := 0; i < GRID_SIZE; i++ {
		row_contains_num := grid[row][i] == num
		col_contains_num := grid[i][col] == num
		grid_contains_num := SubGridContainsNum(grid, row, col, num)
		if row_contains_num || col_contains_num || grid_contains_num {
			return false
		}
	}
	return true
}

// SolveBoard uses backtracking to solve the sudoku puzzle by
// going through every block column by column, then row by row,
// trying every combination for those blocks if it is "0". If SolveBoard cannot solve
// the sudoku puzzle with its current guess, it "backtracks" and continues
// the next possible number. Otherwise, if the block is not "0",
// we move onto the next block
func SolveBoard(grid [][]int, row int, col int) bool {
	if row == GRID_SIZE {
		// we solved every square in the board
		return true
	}
	if col == GRID_SIZE {
		// go to the next row
		return SolveBoard(grid, row+1, 0)
	}

	if grid[row][col] == 0 {
		for num := 1; num <= GRID_SIZE; num++ {
			if IsValidPosition(grid, row, col, num) {
				// try to put this number in this part of the grid
				grid[row][col] = num
				if SolveBoard(grid, row, col+1) {
					// eventually solved, so return
					return true
				} else {
					// undo since it failed
					grid[row][col] = 0
				}
			}
		}
		return false
	} else {
		return SolveBoard(grid, row, col+1)
	}
}
