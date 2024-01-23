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

func findNextEmptyCell(grid [][]int) (int, int, bool) {
	for i:=0; i<BOARD_SIZE; i++ {
		for j:=0; j<BOARD_SIZE; j++ {
			if grid[i][j] == 0{
				return i, j, true
			}
		}
	}
	return -1, -1, false
}


func isSolved(grid [][]int) bool {
	row, col, found := findNextEmptyCell(grid)
	if !found {
		return true
	}
	
	// try numbers from 1 - BOARD_SIZE and back track if bad choices have been made 
	for n := 1; n <= BOARD_SIZE; n++ {
		if isPossible(grid, row, col, n) {
			grid[row][col] = n
			if isSolved(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}


func SolveSudoku(grid[][]int) [][]int {
	// Assuming that we always get a valid sudoku board
	isSolved(grid)
	return grid
}