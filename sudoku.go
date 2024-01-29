package main

func SolveSudoku(sudoku [][]int) [][]int {

	// 2D arrays, in ith, jth location, store 1 if number j+1 is present in ith row, otherwise store 0
	PresentInRow := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	PresentInColumn := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	PresentInBox := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	zeroCounter := 0

	for i := 0; i < 9; i++ {
    
		for j := 0; j < 9; j++ {
			number := sudoku[i][j]

			// either document that a number is present, or increment zero counter
			if number != 0 {
				PresentInRow[i][number-1] = 1 // number-1 to change to 0 indexing
				PresentInColumn[j][number-1] = 1
				boxId := (i/3)*3 + j/3 // upper left box is boxID 0, upper middle boxID 1, upper right boxID 2,...
				PresentInBox[boxId][number-1] = 1
			} else {
				zeroCounter += 1
			}

		}
	}

	sudoku, zeroCounter = sudokuSolverCore(
		sudoku, PresentInRow, PresentInColumn, PresentInBox, zeroCounter, 0, 0)

	return sudoku
}

func sudokuSolverCore(sudoku [][]int, PresentInRow [][]int, PresentInColumn [][]int, PresentInBox [][]int,
	zeroCounter int, iStart int, jStart int) ([][]int, int) {

	for i := iStart; i < 9; i++ {
		for j := jStart; j < 9; j++ {
			jStart = 0 // reset starter variable so future rows are gone through entirely
			number := sudoku[i][j]

			// if location is unsolved
			if number == 0 {

				// cycle through all digits that can appear in a sudoku puzzle
				for numToTry := 1; numToTry < 10; numToTry++ {
					boxId := (i/3)*3 + j/3 // this takes advantage of integer division causing truncation

					// put the candidate for the location into the sudoku puzzle
					if PresentInRow[i][numToTry-1] == 0 && PresentInColumn[j][numToTry-1] == 0 && PresentInBox[boxId][numToTry-1] == 0 {
						sudoku[i][j] = numToTry
						PresentInRow[i][numToTry-1] = 1
						PresentInColumn[j][numToTry-1] = 1
						PresentInBox[boxId][numToTry-1] = 1
						zeroCounter -= 1

						// make recrusive call to go to next location and try out all valid digits there
						if j == 8 {
							sudoku, zeroCounter = sudokuSolverCore(sudoku, PresentInRow, PresentInColumn, PresentInBox, zeroCounter, i+1, 0)
						} else {
							sudoku, zeroCounter = sudokuSolverCore(sudoku, PresentInRow, PresentInColumn, PresentInBox, zeroCounter, i, j+1)
						}

						// revert to previous solution if recrusive calls did not find a solution for the current setup
						if zeroCounter != 0 {
							sudoku[i][j] = 0
							zeroCounter += 1
							PresentInRow[i][numToTry-1] = 0
							PresentInColumn[j][numToTry-1] = 0
							PresentInBox[boxId][numToTry-1] = 0
						}
            
					}
				}
				return sudoku, zeroCounter
			}
		}
	}
	return sudoku, zeroCounter
}
