package main

import "fmt"

//Hey! I'm Athiru Pathiraja & this is my submission for Shopify's infrastructure assessment (PS: I applied for data engineering, but found this challenge more fun)
//TLDR: I implemented the solution using a recursive backtracking algorithm with caching

//Time complexity: O(9 ^ (N * N)), where N is the length/width of board and 9 denotes the possible numeric assignments to each cell. Since the input size is fixed, this can be reduced to O(1)
//PS: I used a caching solution that improved the time complexity from O( 9 ^ (N*N*N)), by improving the check for a valid number assignment to a cell from O(N) to O(1)
//Space complexity: O(3 * (N ^ 2)) for the caches, that can be reduced to O(1) since the input board size is fixed.

//Trade-offs: the caching solution increased the space complexity, but reduced the time complexity by using hashsets with O(1) removing, retrieval & insertion.

const BoardSize = 9

func SolveSudoku(board [][]int) [][]int {

	//declaring & instantiating maps to store if a number exists in a particular row, column, or 3x3 square.
	//all caches have this structure: cache[key]: List[bool] where the list is of length 9 (used to represent numbers 1-9)
	rowsCache := make([]map[int]bool, BoardSize)
	colsCache := make([]map[int]bool, BoardSize)
	squaresCache := make(map[string]map[int]bool, BoardSize)

	initCache(board, &rowsCache, &colsCache, &squaresCache)
	backtrack(board, rowsCache, colsCache, squaresCache)
	return board
}

func backtrack(board [][]int, rowsCache, colsCache []map[int]bool, squaresCache map[string]map[int]bool) bool {

	//Base Case: helper function to check if no cell in board has a '0' (has been completely filled), and we can return from recursive call.
	if isComplete(board) {
		return true
	}

	//nested for loops to identify empty cells
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if board[row][col] == 0 {
				//if an empty cell is found, try all possible number assignments
				for num := 1; num <= BoardSize; num++ {
					if isValid(row, col, num, rowsCache, colsCache, squaresCache) {

						//if the current number doesn't exist is in current cell's row, column or 3x3 square, then assign it to the board in place & add to cache.
						insertDeleteCache(true, row, col, num, &rowsCache, &colsCache, &squaresCache)
						board[row][col] = num

						//recursively call to fill other cells
						if backtrack(board, rowsCache, colsCache, squaresCache) {
							return true
						}

						//if previous recursive call did not complete board, backtrack by removing from cache and resetting board to try different assignment.
						insertDeleteCache(false, row, col, num, &rowsCache, &colsCache, &squaresCache)
						board[row][col] = 0
					}
				}
				//if no valid number assignment for current cell, backtrack to reset board
				return false
			}
		}
	}
	//return false if not possible to solve board
	return false
}

func isComplete(board [][]int) bool {
	//Base case helper function to check if board has been completed.
	for r := 0; r < BoardSize; r++ {
		for c := 0; c < BoardSize; c++ {
			if board[r][c] == 0 {
				return false
			}
		}
	}
	return true
}

func initCache(board [][]int, rowsCache, colsCache *[]map[int]bool, squaresCache *map[string]map[int]bool) {
	for row := 0; row < BoardSize; row++ {
		//complete constructing cache with where value for each key-value pair is a list of booleans
		(*rowsCache)[row] = make(map[int]bool)
		(*colsCache)[row] = make(map[int]bool)

		for col := 0; col < BoardSize; col++ {
			//initializing all cells to false, representing that the cell is empty.
			(*rowsCache)[row][col] = false
			(*colsCache)[row][col] = false
		}
	}

	//constructing & initializing squares cache
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//key is found by dividing to the rows & col by 3, used to represent each 3x3 board
			squareKey := fmt.Sprintf("%d%d", i, j)
			(*squaresCache)[squareKey] = make(map[int]bool)
			for k := 0; k < BoardSize; k++ {
				(*squaresCache)[squareKey][k] = false
			}
		}
	}

	//initializing caches with filled in cells from input board
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			num := board[row][col]
			if num == 0 {
				continue
			} else {
				//setting element corresponding to num in all cache's values to true
				squareKey := fmt.Sprintf("%d%d", row/3, col/3)
				(*rowsCache)[row][num-1] = true
				(*colsCache)[col][num-1] = true
				(*squaresCache)[squareKey][num-1] = true
			}
		}
	}
}

// helper function to check if num exists in any cache, if it does then return false for invalid assignment in O(1) time.
func isValid(row, col, num int, rowsCache, colsCache []map[int]bool, squaresCache map[string]map[int]bool) bool {
	squareKey := fmt.Sprintf("%d%d", row/3, col/3)

	return !(rowsCache[row][num-1] || colsCache[col][num-1] || squaresCache[squareKey][num-1])

}

// helper function to insert or delete from caches in O(1) time.
func insertDeleteCache(insert bool, row, col, num int, rowsCache, colsCache *[]map[int]bool, squaresCache *map[string]map[int]bool) {
	squareKey := fmt.Sprintf("%d%d", row/3, col/3)

	(*rowsCache)[row][num-1] = insert
	(*colsCache)[col][num-1] = insert
	(*squaresCache)[squareKey][num-1] = insert

}

// helper functions to print state of caches (used for debugging)
func printCacheState(rowsCache, colsCache []map[int]bool, squaresCache map[string]map[int]bool) {
	fmt.Println("Rows Cache:")
	printCacheSlice(rowsCache)
	fmt.Println("\nColumns Cache:")
	printCacheSlice(colsCache)
	fmt.Println("\nSquares Cache:")
	printSquareCache(squaresCache)
}

func printSquareCache(squaresCache map[string]map[int]bool) {
	for key, m := range squaresCache {
		fmt.Printf("Key %s: [", key)
		for num := 0; num < BoardSize; num++ {
			fmt.Printf(" %t", m[num])
		}
		fmt.Println(" ]")
	}
}

func printCacheSlice(cache []map[int]bool) {
	for i, m := range cache {
		fmt.Printf("Index %d: [", i)
		for num := 0; num < BoardSize; num++ {
			fmt.Printf(" %t", m[num])
		}
		fmt.Println(" ]")
	}
}
