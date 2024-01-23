package main

import "fmt"

// PS: I think I applied to backend but i liked this challenge better! :D


// Hello so I implemented a back tracking solution to solve this sudoku puzzle.
// It use a variety of helper functions that I will go through with you all!
// I also added a print function to print out the sudoku puzzle.
func main() {
	// here is an example of the sudoku puzzle on the test go file
	sudoku := [][]int{
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
	// here is the function call to solve the sudoku puzzle
	solveSudoku(sudoku)
	// here is the function call to print the sudoku puzzle
	printSudoku(sudoku)
	}

// this is the main function that solves the sudoku puzzle
func solveSudoku(sudoku [][]int) bool {
	// Since the sudoku puzzle is 9x9, we can use 9 as the base case for the row and column
	// we are going to use pointers to keep track of the row and column
	var row, col int

	// This is an edge case that tests if the cells are not empty. Which basically means that the sudoku is solved.
	if !Empty(sudoku, &row, &col) {
		return true
	}
	// This is the main for loop that will iterate through the sudoku puzzle
	for num := 1; num <= 9; num++ {
		// This is the helper function that checks if the number is safe to place in the cell
		if Safe(sudoku, row, col, num) {
			// we then would place it into the cell
			sudoku[row][col] = num
			// then we would recursively call the function again to check if the sudoku puzzle is solved
			// if it is solved, then we would return true
			// if it is not solved, then we would backtrack and try again
			if solveSudoku(sudoku) {
				return true
			}
			// this is the backtracking part
			// we start by setting the cell back to 0 every time we backtrack
			sudoku[row][col] = 0
		}
	}
	// if the sudoku puzzle is not solved, then we would return false
	return false
}
// this is the helper function that checks if the cell is empty
func Empty(sudoku [][]int, row, col *int) bool {
	// basically we are iterating through the sudoku puzzle
	// we use 2 for loops to iterate through the rows and columns
	for *row = 0; *row < 9; *row++ {
		for *col = 0; *col < 9; *col++ {
			if sudoku[*row][*col] == 0 {
				return true
			}
		}
	}
	// if the cell is not empty, then we would return false
	return false
}
// this is the helper function that checks if the number is safe to place in the cell
func Safe(sudoku [][]int, row, col, num int) bool {
	// we are going to use 3 helper functions to check if the number is safe to place in the cell
	// we are going to check if the number is in the row
	// we are going to check if the number is in the column
	// we are going to check if the number is in the box
	return !InRow(sudoku, row, num) &&
		!InCol(sudoku, col, num) &&
		!InBox(sudoku, row-row%3, col-col%3, num)
}
// this is the helper function that checks if the number is in the row
func InRow(sudoku [][]int, row, num int) bool {
	// we are going to iterate through the row
	for i := 0; i < 9; i++ {
		// if the number is in the row, then we would return true
		if sudoku[row][i] == num {
			return true
		}
	}
	// if not false
	return false
}
// this is the helper function that checks if the number is in the column
func InCol(sudoku [][]int, col, num int) bool {
	// we are going to iterate through the column
	for i := 0; i < 9; i++ {
		// if the number is in the column, then we would return true
		if sudoku[i][col] == num {
			return true
		}
	}
	// if not false
	return false
}
// finally, this is the helper function that checks if the number is in the box
func InBox(sudoku [][]int, startRow, startCol, num int) bool {
	// we are going to iterate through the box
	for i := 0; i < 3; i++ {
		// we are going to use 2 for loops to iterate through the rows and columns
		for j := 0; j < 3; j++ {
			// if the number is in the box, then we would return true
			if sudoku[startRow+i][startCol+j] == num {
				return true
			}
		}
	}
	// if not false
	return false
}
// this is the helper function that prints the sudoku puzzle
func printSudoku(sudoku [][]int) {
	for _, row := range sudoku {
		fmt.Println(row)
	}
}

// ok so that is it for the sudoku puzzle
// IF you have any questions, please feel free to ask me!
// Contact me anywhere!
// Hope to hear back from you soon! :)