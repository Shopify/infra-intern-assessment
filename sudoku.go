/*
Shopify Infra Intern Sudoku Assignment!

This file, specifically SolveSudoku has the functionality to solve a sudoku grid within the problem specifications outlined in the README.

Author: Michael Ding
*/

package main

// Constants related to input board
const SUB_BOX_SIZE int = 3
const BOARD_SIZE int = 9
const EMPTY_CELL int = 0

// Possible nums to fill
var CONST_POSSIBLE_NUMS = [BOARD_SIZE]int{1, 2, 3, 4, 5, 6, 7, 8, 9}


// Solves the given sudoku grid.
// Preconditions: There must only exist 1 solution to the input_grid as constrained by the problem description.
func SolveSudoku (input_grid [][]int) [][]int {
	// Maintain state of the nums we have filled. This gives us O(1) lookups (though our board is technically 
	// constant 9x9 already)
	rowNums := make([]map[int]bool, BOARD_SIZE) // allocate the appropriate size array of maps
	colNums := make([]map[int]bool, BOARD_SIZE)
	subBoxNums := make([]map[int]bool, BOARD_SIZE)

	for i := 0; i < BOARD_SIZE; i++ {
        rowNums[i] = make(map[int]bool)
        colNums[i] = make(map[int]bool)
        subBoxNums[i] = make(map[int]bool)
    }

	// Populate the maps with existing numbers in the grid
    for row := 0; row < BOARD_SIZE; row++ {
        for col := 0; col < BOARD_SIZE; col++ {
            num := input_grid[row][col]
            if num != EMPTY_CELL {
                rowNums[row][num] = true
                colNums[col][num] = true
                subBoxNums[getSubBoxId(row, col)][num] = true
            }
        }
    }

	// in place modify input_grid
	backTrackSolve(input_grid, rowNums, colNums, subBoxNums, 0, 0)

	return input_grid
}

// Uses recursion/backtracking to solve this problem. Uses a DFS approach by filling in a number if possible and
// continuing until there the "path" (board) is a deadend (unsolved, no more options). Solves left to right, top to down
func backTrackSolve (input_grid [][]int, rowNums, colNums, subBoxNums []map[int]bool, curr_row, curr_col int) bool {
	// Check edge cases where we dont perform exhaustive backtracking
	if reachedBoardEnd(curr_row, curr_col) {
		return true
	} else if reachedColEnd(curr_col) {
		return backTrackSolve(input_grid, rowNums, colNums, subBoxNums, curr_row + 1, 0)
	} else if input_grid[curr_row][curr_col] != EMPTY_CELL {
		return backTrackSolve(input_grid, rowNums, colNums, subBoxNums, curr_row, curr_col + 1)
	}

	// DFS/recursive backtrack all the options
	for _, filledNum := range CONST_POSSIBLE_NUMS {
		// Explore down this path if it is a valid placement.
		if checkValidPlacement(rowNums, colNums, subBoxNums, curr_row, curr_col, filledNum) {
			input_grid[curr_row][curr_col] = filledNum
			rowNums[curr_row][filledNum] = true // update the maps with the filled number
			colNums[curr_col][filledNum] = true
			subBoxNums[getSubBoxId(curr_row, curr_col)][filledNum] = true

			// If the valid solution is down this path, then return
			if backTrackSolve(input_grid, rowNums, colNums, subBoxNums, curr_row, curr_col + 1) {
				return true
			}

			// remove to try the next option
			input_grid[curr_row][curr_col] = 0
			delete(rowNums[curr_row], filledNum)
			delete(colNums[curr_col], filledNum)
			delete(subBoxNums[getSubBoxId(curr_row, curr_col)], filledNum)
		}
	}

	// We have exhausted all choices, this is a dead path, backtrack and return false
	return false
}

// Check if we can actually place num in the location row, col.
func checkValidPlacement (rowNums, colNums, subBoxNums []map[int]bool, row, col, numOfInterest int) bool {
	// If this number exists in any row, col, or subBox, return false
	if rowNums[row][numOfInterest] || colNums[col][numOfInterest] || subBoxNums[getSubBoxId(row, col)][numOfInterest] {
		return false
	}

	return true
}

// Gets the subBox number of the row and col. The subBoxes are numbered starting from 0, from left
// to right, then top to bottom
func getSubBoxId (row, col int) int {
	// Floor division to get the row num in the SUB_BOX_SIZE x SUB_BOX_SIZE (3x3) grid
	subBoxRowNum := row / SUB_BOX_SIZE
	subBoxColNum := col / SUB_BOX_SIZE

	// Return the subBox id number, from left to right, top to bottom
	return subBoxRowNum * 3 + subBoxColNum
}

func reachedBoardEnd (curr_row, curr_col int) bool {
	return curr_row == BOARD_SIZE - 1 && curr_col == BOARD_SIZE
}

func reachedColEnd (curr_col int) bool {
	return curr_col == BOARD_SIZE
}

