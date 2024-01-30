// By Chris Lozinski January 29, 2024
package main

//import all necessary packages
import (
	"fmt" // for i/o
)

/*
test main function for basic input testing without sudoku_test.go
*/
func main() {
	puzzle := [][]int{ // puzzle to be solved, all zeros are empty cells
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

	fmt.Println("Unsolved Puzzle:")
	printGrid(puzzle) // print the unsolved puzzle in terminal

	solvedPuzzle := SolveSudoku(puzzle) // send input 2D array directly to puzzleSolver

	if solvedPuzzle == nil { // checks to see if the input is solvable
		fmt.Println("\nThere is no solution to this puzzle!")
	} else {
		fmt.Println("\nSolved Puzzle:")
		printGrid(solvedPuzzle) // print solved puzzle
	}
}

/*
print the sudoku board
*/
func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

/*
solve the sudoku board
*/
func SolveSudoku(puzzle [][]int) [][]int {
	emptyCells := isCellEmpty(puzzle)

	if len(emptyCells) == 0 { // if theres no empty cells the puzzle has been solved
		return puzzle
	}

	row := emptyCells[0][0]
	col := emptyCells[0][1]

	for num := 1; num <= 9; num++ {
		if checkMove(puzzle, row, col, num) {
			puzzle[row][col] = num
			if result := SolveSudoku(puzzle); result != nil {
				return result
			}
			puzzle[row][col] = 0 // Backtrack
		}
	}

	return nil // if there is no solution to the
}

/*
look for empty cells
*/
func isCellEmpty(puzzle [][]int) [][]int {
	emptyCells := make([][]int, 0) // initialize array to collect empty cells
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if puzzle[i][j] == 0 {
				emptyCells = append(emptyCells, []int{i, j})
			}
		}
	}
	return emptyCells // returns empty cells
}

/*
checks to see if the number is already in one of the rows or columns and if the move is not allowed
*/
func checkMove(puzzle [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ { // check all nine cells
		if puzzle[row][i] == num || puzzle[i][col] == num { // if either is true move is not valid
			return false // if move is not valid
		}
	}

	startingRow, startingCol := (row/3)*3, (col/3)*3

	for i := 0; i < 3; i++ { // to iterate through rows of grid
		for j := 0; j < 3; j++ { // for columns of grid
			if puzzle[startingRow+i][startingCol+j] == num { // check if number is already present
				return false
			}
		}
	}

	return true // if number is not present !
}
