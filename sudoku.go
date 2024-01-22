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

	for zeroCounter > 0 {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {

				number := sudoku[i][j]

				// if location is unsolved
				if number == 0 {

					boxId := (i/3)*3 + j/3
					eligibleNumber := 0
					eligibleNumbersCount := 0

					// see how many numbers can be used in that location
					for k := 0; k < 9; k++ {
						if PresentInRow[i][k] == 0 && PresentInColumn[j][k] == 0 && PresentInBox[boxId][k] == 0 {
							eligibleNumbersCount += 1
							eligibleNumber = k + 1
						}
					}

					// put the solution for the location into the sudoku puzzle (if only 1 solution)
					if eligibleNumbersCount == 1 {
						sudoku[i][j] = eligibleNumber
						PresentInRow[i][eligibleNumber-1] = 1
						PresentInColumn[j][eligibleNumber-1] = 1
						PresentInBox[boxId][eligibleNumber-1] = 1
						zeroCounter -= 1
					}

				}
			}
		}
	}

	return sudoku
}
