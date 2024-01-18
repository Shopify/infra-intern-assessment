package main

// this is just the setup, will construct and debug from here
func SolveSudoku(puzzle [][]int) [][]int {
	solve(puzzle)
	return puzzle
}

func solve(SudokuBoard [][]int) bool {

	const BOARD_SIZE = 9
	const EMPTY_CELL = 0

	// we use these loops to iterate through each 
	// grid square [i][j] 
	for i := 0; i < BOARD_SIZE; i++ {
			for j := 0; j < BOARD_SIZE; j++ {

					// Preliminary check to see if square is empty
					if SudokuBoard[i][j] == EMPTY_CELL {

							// we try all values of n until 
							// one of them is valid 
							for n := 1; n <= 9; n++ {

									// we place the first valid value of n into (i, j)
									if canPlaceN(SudokuBoard, i, j, n) {
										SudokuBoard[i][j] = n


											// this is a recursive call to
											// puzzle, essentially, what is happening here is
											// solve(puzzle) branches into two cases 
											// (1) inserting n into (i, j) is valid in a solution
											//  	 of the puzzle, in which case solve(puzzle)
											//  	 continues where we left off, and returns true
											//     once completing the puzzle 
											// (2) inserting n into (i, j) is NOT a valid step 
											// 	   in a solution of the puzzle (given all previous
											//		 steps are set in stone), in which case we backtrack 
											// 	   and try the next value of n.	
											if solve(SudokuBoard) {
													return true
											}

											SudokuBoard[i][j] = EMPTY_CELL // backtrack
									}
							}
							return false
					}
			}
	}
	return true
}




func canPlaceN(p [][]int, i, j, n int) bool {
	const BOARD_SIZE = 9
	const SUBGRID_SIZE = 3

	// this function checks if we can place a number (n) at position (i, j)
	// in the sudoku grid (9x9 size assumed)
	for k := 0; k < BOARD_SIZE; k++ {
			// if we find another instance of 
			// n in the same col/row, we return early 
			if p[i][k] == n || p[k][j] == n {
					return false
			}

			// this part essentially 
			// patitions the 9x9 grid into 
			// 9 3x3 grids, and treats each i
			// as its own subspace, checking if
			// n already exists in the grid. 
			subGridRow := SUBGRID_SIZE * (i / SUBGRID_SIZE) 
			subGridColumn := SUBGRID_SIZE * (j / SUBGRID_SIZE)


			rowOffset, colOffset := divmod(k, SUBGRID_SIZE)

			// we use the index of the grid 
			if p[subGridRow + rowOffset][subGridColumn + colOffset] == n {
					return false
			}
	}
	return true 
}

func divmod(numerator, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator 
}


