package main
import "fmt"

const size int = 9; // size of board
const sizeGrid int = 3; // size of each subgrid

// check if it is valid to place num at board[row][col]
func isValid(row int, col int, num int, board [][]int) bool{
	// row validity
	for c:=0; c<size; c++{
		if board[row][c] == num {return false}
	}
	// column validity
	for r:=0; r<size; r++{
		if board[r][col] == num {return false}
	}
	// grid validity
	rowStart:=row/sizeGrid*sizeGrid
	colStart:=col/sizeGrid*sizeGrid
	for r:=rowStart; r<rowStart+3; r++ {
		for c:=colStart; c<colStart+3; c++ {
			if board[r][c] == num {return false}
		}
	}
	return true;
}

// recursive solution with backtracking. 
func solve(board [][]int) [][]int {
	for row:=0; row<size; row++ {
		for col:=0; col<size; col++ {
			// find unfilled spaces
			if board[row][col] != 0 {continue}
			// test if a digit can fill the space
			for num:=1; num<=size; num++ {
				if (!isValid(row, col, num, board)) {continue}
				// recursively solve after filling space
				board[row][col] = num
				newBoard := solve(board)
				// if no solution, backtrack and continue.
				if newBoard == nil {
					board[row][col] = 0
				} else {
					return newBoard
				}
			}
			// no solution found (no valid digit)
			return nil;
		}
	}
	return board
}

// Takes a 9x9 unsolved sudoku board with one solution as input in the 
// form of a nested array with and prints the solution. Empty spaces 
// are represented as 0s.
func SolveSudoku(board [][]int) [][]int{
	board = solve(board);

	// print board
	for r, _ := range board {
		for c, _ := range board[r] {
			fmt.Printf("%d ", board[r][c])
		}
		fmt.Println()
	}
	return board
}