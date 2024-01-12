/*
	Explantion to the solution:

	1. I defined a function `tryValuesConcurrently` that takes a Sudoku puzzle and a cell position (row and column) as input.

	2. I created a `WaitGroup` to keep track of all the goroutines that I'm going to start.

	3. I created a buffered channel `solution` that can hold one solution. This is where my goroutines will send the solved puzzle.

	4. I started a loop from 1 to 9. For each number:

	- I incremented the `WaitGroup` counter by 1.

	- I started a goroutine. Inside the goroutine:

		- I made a copy of the puzzle.

		- I checked if the current number is a valid entry for the current cell using the `isPossible` function.

		- If the number is valid, I placed it in the cell and called `SolveSudoku` recursively on the modified puzzle.

		- If `SolveSudoku` found a solution, I sent the solved puzzle to the `solution` channel. I used a `select` statement to ensure that if the channel is already full (a solution has already been found), the goroutine doesn't block and can finish.

		- I decremented the `WaitGroup` counter by 1.

	5. After starting all goroutines, I waited for all of them to finish using `wg.Wait()`. This ensures that the function doesn't return before all goroutines have had a chance to find a solution.

	This way, I used concurrency to speed up the Sudoku solving process by trying all possible numbers for a cell at the same time.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * Solves a sudoku puzzle using backtracking
 * @param puzzle The sudoku puzzle to solve
 * @return The solved sudoku puzzle
 **/
func SolveSudoku(puzzle [][]int) [][]int {
	// Iterate over each row
	for row := 0; row < 9; row++ {
		// Iterate over each column
		for col := 0; col < 9; col++ {
			if puzzle[row][col] == 0 {
				return tryValuesConcurrently(puzzle, row, col)
			}
		}
	}
	return puzzle
}

/**
  * tryValuesConcurrently tries values 1 to 9 concurrently in the given cell
  * @param puzzle The sudoku puzzle
  * @param row The row of the cell
  * @param col The column of the cell
  * @return The solved sudoku puzzle
**/
func tryValuesConcurrently(puzzle [][]int, row int, col int) [][]int {
	var wg sync.WaitGroup
	solution := make(chan [][]int, 1)

	for val := 1; val <= 9; val++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			puzzleCopy := makeCopyOfPuzzle(puzzle)
			if isPossible(row, col, val, puzzleCopy) {
				puzzleCopy[row][col] = val
				if solved := SolveSudoku(puzzleCopy); solved != nil {
					select {
					case solution <- solved:
					default:
					}
				}
			}
		}(val)
	}

	wg.Wait()
	close(solution)

	if len(solution) > 0 {
		return <-solution
	}

	return nil
}

/**
  * makeCopyOfPuzzle creates a deep copy of the puzzle
  * @param src The sudoku puzzle to copy
  * @return The copy of the sudoku puzzle
**/
func makeCopyOfPuzzle(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

/**
 * Checks if a value is possible in a given position.
 * We check if there exists a value in the same row, column or box with the same value.
 * @param y The y coordinate / row
 * @param x The x coordinate / column
 * @param val The value to check
 * @param puzzle The sudoku puzzle
 * @return True if the value is possible, false otherwise
 **/
func isPossible(row int, col int, val int, puzzle [][]int) bool {

	// Check for row
	for i := 0; i < 9; i++ {
		if puzzle[row][i] == val {
			return false
		}
	}

	// Check for column
	for i := 0; i < 9; i++ {
		if puzzle[i][col] == val {
			return false
		}
	}

	// Check for box
	boxX := (col / 3) * 3
	boxY := (row / 3) * 3
	for i := boxY; i < boxY+3; i++ {
		for j := boxX; j < boxX+3; j++ {
			if puzzle[i][j] == val {
				return false
			}
		}
	}
	return true
}

/**
 * Checks if a sudoku puzzle is valid
 * @param puzzle The sudoku puzzle to check
 * @return True if the puzzle is valid, false otherwise
 **/
func isValid(puzzle [][]int) bool {
	if len(puzzle) != 9 {
		return false
	}
	if len(puzzle[0]) != 9 {
		return false
	}
	for _, row := range puzzle {
		for _, val := range row {
			if val < 0 || val > 9 {
				return false
			}
		}
	}
	return true
}

/**
 * Prints a sudoku puzzle
 * @param puzzle The sudoku puzzle to print
 **/
func print(puzzle [][]int) {
	for _, row := range puzzle {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// The below code can be uncommented to print the solution. Replace the input with the puzzle you want to solve.
/**
func main() {
	input := [][]int{
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

	if isValid(input) {
		print(SolveSudoku(input))
		// Uncomment the below line to measure the average run time
		measureAvgRunTime(input)
	}

}
*/

/**
 * Measures the average run time of the SolveSudoku function by running it 500 times
 **/
func measureAvgRunTime(input [][]int) {
	startTime := time.Now()
	for i := 0; i < 500; i++ {
		SolveSudoku(input)
	}
	endTime := time.Now()
	fmt.Println("Average run time: ", endTime.Sub(startTime)/500)
}
