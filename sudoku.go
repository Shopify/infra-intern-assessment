/*
Shopify Intern Assessment Infrastructure Engineering - Vitaly Pastukhov

This program implement Backtracking algorithm to solve Sudoku 9x9 puzzle by filling the empty cells.
Function called SolveSudoku takes a 9x9 grid as input and returns the solved Sudoku grid.
Function SolveSudoku also prints the solved grid to the console.

Backtracking approach reduce substantially the complexity of the algorithm in comparison to the Brute Force solution.
Brute force generate all possible ways to fill the empty cells with the numbers from 1 to 9, which means 9^81 operations in the worst case.

Constraints:
1. The input grid is a 9x9 two-dimensional array of integers.
2. The input grid have exactly one solution.
3. The input grid may contain zeros (0) to represent empty cells.

General overview of the Algorithm:
1. Create a dictionary for each row, column and box(3x3 sub-grid) to keep track of the numbers (1-9) used.
2. Initialize dictionaries with the numbers from the input sudoku board.
3. Start from the upper left element (row = 0, col = 0),
   iterate from left to right through each row until we reach an empty cell (0)
4. Iterate over the numbers from 1 to 9
5. If the number does not exist in the current row, column and box, place the number.
   Placing the number includes:
   - updating the row, column and box dictionaries
   - adding the number to the board
6. Go to step 3 and start from the next cell from just the filled one
7. If there is no numbers left to fill the empty cell, we have to backtrack and remove the number from the previous filled cell
   Removing the number includes:
   - removing the number from row, column and box dictionaries
   - removing the number from the board
   Then go to step 4 to try place next number to the cell.
*/

package main

import "fmt"

// Row and column size
const BOARD = 9

/*
Main function
Input: 9x9 two-dimensional array of integers, input Sudoku grid
Output: Solved Sudoku 9x9 grid
*/
func SolveSudoku(board [][]int) [][]int {

	// create empty rows dictionary
	rows := make([][]int, BOARD)
	for i := range rows {
		rows[i] = make([]int, BOARD+1)
	}

	// create empty columns dictionary
	cols := make([][]int, BOARD)
	for i := range cols {
		cols[i] = make([]int, BOARD+1)
	}

	// create empty boxes dictionary
	boxes := make([][]int, BOARD)
	for i := range boxes {
		boxes[i] = make([]int, BOARD+1)
	}

	//Iterate through each element of the input board and add existing not zero numbers to dictionaries (rows, cols, boxes)
	for i := 0; i < BOARD; i++ {
		for j := 0; j < BOARD; j++ {
			if board[i][j] < 0 || board[i][j] > BOARD {
				panic("Invalid board")
			}
			if board[i][j] != 0 {
				num := board[i][j]
				PlaceNum(board, num, i, j, rows, cols, boxes)
			}
		}
	}

	// run Backtrack algo, print the board to the console if the solution found
	if Backtrack(board, 0, 0, rows, cols, boxes) {
		fmt.Print("[\n")
		for row := 0; row < BOARD; row++ {
			fmt.Print("   [")
			for col := 0; col < BOARD; col++ {
				if col == BOARD-1 {
					fmt.Print(board[row][col])
				} else {
					fmt.Print(board[row][col], ", ")
				}
			}
			if row == BOARD-1 {
				fmt.Print("]\n")
			} else {
				fmt.Print("],\n")
			}
		}
		fmt.Print("]\n")
	}
	return board
}

func Backtrack(board [][]int, row int, col int, rows [][]int, cols [][]int, boxes [][]int) bool {

	// iterate until we found an empty cell
	for row < BOARD && (col == BOARD || board[row][col] != 0) {
		if col == BOARD { //move to the next row
			col = 0
			row++
		} else { // move to the next column
			col++
		}
	}

	//reach the end of the board, found the solution
	if row == BOARD {
		return true
	}

	//iterate over all potential options
	for num := 1; num < 10; num++ {
		if CanPlaceNum(board, num, row, col, rows, cols, boxes) { //num does not present in current row, column and box
			//add number to the board and dictionaries
			PlaceNum(board, num, row, col, rows, cols, boxes)
			//move to the next empty cell
			if Backtrack(board, row, col+1, rows, cols, boxes) {
				return true
			} else { //all options did not work out, backtrack and remove the number
				RemoveNum(board, num, row, col, rows, cols, boxes)
			}
		}
	}
	//all options did not work out
	return false
}

//function check if the number present in the current row, column and box.
func CanPlaceNum(board [][]int, num int, row int, col int, rows [][]int, cols [][]int, boxes [][]int) bool {
	return !(rows[row][num] == 1 || cols[col][num] == 1 || boxes[BoxIndex(row, col)][num] == 1)
}

func PlaceNum(board [][]int, num int, row int, col int, rows [][]int, cols [][]int, boxes [][]int) {
	rows[row][num] = 1
	cols[col][num] = 1
	boxes[BoxIndex(row, col)][num] = 1
	board[row][col] = num
}

func RemoveNum(board [][]int, num int, row int, col int, rows [][]int, cols [][]int, boxes [][]int) {
	rows[row][num] = 0
	cols[col][num] = 0
	boxes[BoxIndex(row, col)][num] = 0
	board[row][col] = 0
}

//function BoxIndex return box index based on row and column number
func BoxIndex(row int, col int) int {
	return (row/3)*3 + col/3
}
