package main

const (
	Shape = 9
	Grid  = 3
	Empty = 0
)

var digits map[int]bool

// Creates a set of digits from 1-9
func init() {
	digits = make(map[int]bool)
	for i := 1; i <= 9; i++ {
		digits[i] = true
	}
}

func SolveSudoku(board [][]int) [][]int {
	boardCopy := matrixDeepCopy(board)
	search(boardCopy)
	return boardCopy
}

// Backtracking algorithm that recursively tries every unused digit until board is solved or no solution found
func search(board [][]int) bool {
	if isValidState(board) {
		return true
	}

	for rowIdx, row := range board {
		for colIdx, cell := range row {
			if cell == Empty {
				for candidate := range getCandidates(board, rowIdx, colIdx) {
					board[rowIdx][colIdx] = candidate
					isSolved := search(board)
					if isSolved {
						return true
					}
					board[rowIdx][colIdx] = Empty
				}
				return false
			}
		}
	}

	return true
}

// Checks if the board is solved
func isValidState(board [][]int) bool {
	for _, row := range board {
		rowSet := makeSet(row)
		if !setIsEqual(rowSet, digits) {
			return false
		}
	}

	for _, col := range getCols(board) {
		colSet := makeSet(col)
		if !setIsEqual(colSet, digits) {
			return false
		}
	}

	for _, grid := range getGrids(board) {
		gridSet := makeSet(grid)
		if !setIsEqual(gridSet, digits) {
			return false
		}
	}
	return true
}

// Gets digits that are unused
func getCandidates(board [][]int, rowIdx, colIdx int) map[int]bool {
	usedDigits := make(map[int]bool)
	usedDigits = setUnion(usedDigits, makeSet(getKthRow(board, rowIdx)))
	usedDigits = setUnion(usedDigits, makeSet(getKthCol(board, colIdx)))
	usedDigits = setUnion(usedDigits, makeSet(getGridAtRowCol(board, rowIdx, colIdx)))
	delete(usedDigits, Empty) // Remove Empty from set

	candidates := setDifference(digits, usedDigits)
	return candidates
}

// Helper functions

func matrixDeepCopy(board [][]int) [][]int {
	boardCopy := make([][]int, len(board))
	for i := range board {
		boardCopy[i] = make([]int, len(board[i]))
		copy(boardCopy[i], board[i])
	}
	return boardCopy
}

// Gets column slices from matrix
func getCols(board [][]int) [][]int {
	cols := make([][]int, len(board))
	for col := range cols {
		cols[col] = make([]int, len(board[col]))
		for row := range board {
			cols[col][row] = board[row][col]
		}
	}
	return cols
}

// Gets grid slices from matrix. For every 3x3 matrix in a board, returns a flattened slice.
func getGrids(board [][]int) [][]int {
	grids := make([][]int, 0)
	for row := 0; row < Shape; row += Grid {
		for col := 0; col < Shape; col += Grid {
			grid := make([]int, 0)
			for r := row; r < row+Grid; r++ {
				for c := col; c < col+Grid; c++ {
					grid = append(grid, board[r][c])
				}
			}
			grids = append(grids, grid)
		}
	}
	return grids
}

func getKthRow(board [][]int, row int) []int {
	return board[row]
}

func getKthCol(board [][]int, col int) []int {
	kthCol := make([]int, Shape)
	for row := 0; row < Shape; row++ {
		kthCol[row] = board[row][col]
	}
	return kthCol
}

// Gets 3x3 subgrid as a flattened slice.
func getGridAtRowCol(board [][]int, row, col int) []int {
	row = row / Grid * Grid
	col = col / Grid * Grid

	grid := make([]int, 0)
	for r := row; r < row+Grid; r++ {
		for c := col; c < col+Grid; c++ {
			grid = append(grid, board[r][c])
		}
	}
	return grid
}

// func printMatrix(board [][]int) {
// 	for _, row := range board {
// 		fmt.Println(row)
// 	}
// }

// Set helper functions

// Makes a set from a slice
func makeSet(arr []int) map[int]bool {
	set := make(map[int]bool, 0)
	for _, element := range arr {
		set[element] = true
	}
	return set
}

func setIsEqual(set1, set2 map[int]bool) bool {
	if len(set1) != len(set2) {
		return false
	}

	for key := range set1 {
		if _, exists := set2[key]; !exists {
			return false
		}
	}
	return true
}

func setDifference(set1, set2 map[int]bool) map[int]bool {
	res := make(map[int]bool)
	for key := range set1 {
		if _, exists := set2[key]; !exists {
			res[key] = true
		}
	}
	return res
}

func setUnion(setA, setB map[int]bool) map[int]bool {
	setUnion := make(map[int]bool)
	for k := range setA {
		setUnion[k] = true
	}
	for k := range setB {
		setUnion[k] = true
	}
	return setUnion
}
