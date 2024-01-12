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

// tryValuesConcurrently tries values 1 to 9 concurrently in the given cell
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

// makeCopyOfPuzzle creates a deep copy of the puzzle
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
	input := [][]int {
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

	print(SolveSudoku(input))
	// Uncomment the below line to measure the average run time
	// measureAvgRunTime()

}
*/

/**
 * Measures the average run time of the SolveSudoku function by running it 500 times
 **/
func measureAvgRunTime() {
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

	startTime := time.Now()
	for i := 0; i < 500; i++ {
		SolveSudoku(input)
	}
	endTime := time.Now()
	fmt.Println("Average run time: ", endTime.Sub(startTime)/500)
}
