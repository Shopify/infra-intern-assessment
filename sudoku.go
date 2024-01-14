/*Shopify Infrastructure Intern Assessment - Sudoku Solver 
Rick Zhang

This Go program provides a backtracking algorithm-based Sudoku solver.
The 'SolveSudoku' function attempts to solve a 9x9 Sudoku puzzle. 
We assume that there exists a solution to the sudoku puzzle. When successful, 
we will print the board out.

This program creates constraints as the board is filled and uses 
backtracking in order to reduce the number of combinations to consider.

Preconditions:
1. The input grid will be a 9x9 two-dimensional array of integers.
2. The input grid will have exactly one solution.
3. The input grid may contain zeros (0) to represent empty cells.

Postconditions:
The solved sudoku puzzle will be printed.
The solved sudoku puzzle will be returned.


Here is a general overview of the algorithm:
1. Create a counter for the numbers used in each row, column and box.
2. Starting from (0,0), we will iterate through each row from left to right, 
	as if we were reading the board like a book.
3. If the cell is not already set, try placing every feasible number (a feasible number is one that does not 
	violate the rules of sudoku) in the cell. 
	Placing a number in an empty cell includes:
	 - adding that number to the corresponding row, column and box counters
	 - adding the number to the board
4. After placing a number, go back to step 2, starting at the newly filled cell.
5. If there are no feasible numbers, we have made a wrong move. We will backtrack.
	Backtracking includes:
	 - removing the number from the cell
	 - removing the number from its corresponding row, column and box counters.
	Now, go back to step 3 to try a different number.
*/
package main

import "fmt"


const BOARD_SIZE = 9

func printBoard(board [][]int) {
	// iterate through the entire board and print each number, with some borders.
	for y, row := range board {
		for x, item := range row {
			if x % 3 == 2 && x < 8{
				fmt.Printf("%d | ", item)
			} else {
				fmt.Printf("%d ", item)
			}
			
		}
		fmt.Println()
		if y % 3 == 2 && y < 8{
			fmt.Println("---------------------")
		} 
	}
}

func SolveSudoku(board [][]int) [][]int{

	//Create row, column and box counters.

	// row[i][j] is 1 if the number j is in row i and 0 otherwise.
	row := make([][]int, BOARD_SIZE)
	for i := range row {
		row[i] = make([]int, BOARD_SIZE+1)
	}
		
	// column[i][j] is 1 if the number j is in column i and 0 otherwise.
	column := make([][]int, BOARD_SIZE)
	for i := range column {
		column[i] = make([]int, BOARD_SIZE+1)
	}

	// Consider the board to be a 3x3 grid of boxes. We can label each box
	/*
		| 0 | 1 | 2 |
		| 3 | 4 | 5 |
		| 6 | 7 | 8 |
	*/
	// Note that we can find the box number using x and y. Let box number be b.
	// b = (y/3) * 3 + x/3. Then,
	// box[i][j] is 1 if the number j is in box i, and 0 otherwise.
	box := make([][]int, BOARD_SIZE)
	for i := range box {
		box[i] = make([]int, BOARD_SIZE+1)
	}

	//go through the input board and update the row, column and box counters.
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			item := board[y][x]
			if item != 0 {
				boxNumber := (y/3) * 3 + x/3
				row[y][item] = 1
				column[x][item] = 1
				box[boxNumber][item] = 1
			}
		}
	}
	//backtracking starting from the top-left corner of the board, printing if we find a solution.
	if backtrack(board, 0, 0, row, column, box) {
		printBoard(board)
	}
	return board
}	

func backtrack(board [][]int, x int, y int, row [][]int, column [][]int, box[][]int) bool {
	// iterate until we are done filling the board (y==9) or we have found an empty spot.
	for y < BOARD_SIZE && (x == BOARD_SIZE || board[y][x] != 0) {
		if x == BOARD_SIZE { // we reached the end of a row, go to the next row
			x = 0
			y++
		} else { // otherwise, keep going down the row.
			x++
		}
	}
	if y == 9 { // we finished the board
		return true
	}

	// try all the different digits, if they are feasible. If we find a solution, 
	// then we are done, and we can return. Otherwise, backtrack (undo the digit that we tried).
	for i := 1; i <= 9; i++ {
		if feasible(board, x, y, i, row, column, box) {
			//try adding number i to board[y][x]. We add it to the row, column and box that it is in as well.
			boxNumber := (y/3) * 3 + x/3
			row[y][i] = 1
			column[x][i] = 1
			box[boxNumber][i] = 1
			board[y][x] = i
			if backtrack(board, x+1, y, row, column, box){ // if success, don't undo the move, just return.
				return true
			} else { // if this try failed, we backtrack and undo the changes that we tried.
				row[y][i] = 0
				column[x][i] = 0
				box[boxNumber][i] = 0
				board[y][x] = 0
			}
		}
	}
	return false
}

func feasible(board [][]int, x int, y int, n int, row [][]int, column [][]int, box[][]int) bool {
	//n is the number that we are trying to place at board[y][x].
	boxNumber := (y/3) * 3 + x/3
	return !(row[y][n] == 1 || column[x][n] == 1 || box[boxNumber][n] == 1)
}