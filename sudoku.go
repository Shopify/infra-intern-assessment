package main
import "fmt"

const size int = 9;
const sizeGrid int = 3;

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

//find first open slot, put num that works and recursively call
//if recursive ends in no solution, get rid and find next thing to pu

func SolveSudoku(board [][]int){
	board = solve(board);

	for r, _ := range board {
		for c, _ := range board[r] {
			fmt.Printf("%d ", board[r][c])
		}
		fmt.Println()
	}
}

func solve(board [][]int) [][]int {
	for row:=0; row<size; row++ {
		for col:=0; col<size; col++ {
			if board[row][col] != 0 {continue}
			for num:=1; num<=size; num++ {
				if (!isValid(row, col, num, board)) {continue} //notvalid. has to have some valid as there is solution
				board[row][col] = num;

				newBoard := solve(board)
				if newBoard == nil {
					board[row][col] = 0
				} else {
					return newBoard
				}
			}
			return nil;
		}
	}
	return board
}

func main() {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	SolveSudoku(board)
}