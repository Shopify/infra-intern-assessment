package sudoku


const BOARD_SIZE = 9
const SMALL_GRID_SIZE  = 3


func isPossible(grid [][]int, row int, col int, n int) bool {
	// Check to see if number is unique for row and column 
	for i:=0; i<BOARD_SIZE; i++ {
		if grid[row][i] == n || grid[i][col] == n {
			return false
		}
	}
	
	offsetRow := (row/SMALL_GRID_SIZE)*SMALL_GRID_SIZE
	offsetCol := (col/SMALL_GRID_SIZE)*SMALL_GRID_SIZE

	// Check to see if number is unique for 3x3 subgrid
	for i:=0; i<SMALL_GRID_SIZE; i++ {
		for j:=0; j<SMALL_GRID_SIZE; j++ {
			if grid[i+offsetRow][j+offsetCol] == n {
				return false
			}
		}
	}

	return true
}

func isSolved(grid [][]int, row int, col int) bool {
	// Move along the board
	if row == BOARD_SIZE {
		row = 0
		if col++; col == BOARD_SIZE {
			return true
		}
	}

	if grid[row][col] != 0 {
		return isSolved(grid, row+1, col)
	}

	// try numbers from 1 - BOARD_SIZE and back track if bad choices have been made 
	for n := 1; n <= BOARD_SIZE; n++ {
		if isPossible(grid, row, col, n) {
			grid[row][col] = n
			if isSolved(grid, row+1, col) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}


func SolveSudoku(grid[][]int) [][]int {
	// Assuming that we always get a valid sudoku board
	isSolved(grid, 0, 0) 
	return grid
}