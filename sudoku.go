package main

func SolveSudoku(puzzle [][] int)  [][] int{
    // Define arrays to memorize if a value is seen in a row, a column, or a 3x3 grid.
    // Using bit mask to save some space.
    rowSeen := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
    colSeen := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
    gridSeen := [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	// Scan the original puzzle. Note that we are using zero-based indices to represent numbers 1-9.
    for y := 0; y < 9; y ++ {
        for x := 0; x < 9; x ++ {
            if puzzle[y][x] != 0 {
                k := int(puzzle[y][x]) - 1 // Convert numbers 1-9 into zero-based indices.
                rowSeen[y] |= (1 << k)
                colSeen[x] |= (1 << k)
                gridSeen[y/3][x/3] |= (1 << k)
            }
        }
    }

    /**
	 * Use backtracking to try all combinations. 
	 *
	 * @param idx represents the index of the puzzle after flattening. 80 indicates (8, 8).
	 * @return a boolean flag to indicate if the puzzle is solved.
	 */
    var backtrack func(idx int) bool

    backtrack = func(idx int) bool{
        if idx >= 81 {
            return true
        }

        y := idx / 9
        x := idx % 9        

        if puzzle[y][x] != 0 {
            return backtrack(idx + 1)
        }

        finish := false

		// Try every candidate (1-9) and check if the candidate has been seen before (in a row, a column, or a grid).
		// Note that here we are use indices to represent numbers 1-9.
        for k := 0; k < 9; k ++ {
            if (rowSeen[y] >> (k)) & 1 == 0 && (colSeen[x] >> (k)) & 1  == 0 && (gridSeen[y/3][x/3] >> (k)) & 1 == 0 {
                puzzle[y][x] = k + 1 // Convert zero-based indices into numbers 1-9.
				rowSeen[y] |= 1 << k
                colSeen[x] |= 1 << k
                gridSeen[y/3][x/3] |= 1 << k
				
                finish = backtrack(idx + 1)
				// If the puzzle is solved, the break will skip the later steps to keep the new values in the puzzle.
                if finish {
                    break
                }

				// Restore the puzzle if the current candidate does not fit.
                rowSeen[y] &= ^(1 << k)
                colSeen[x] &= ^(1 << k)
                gridSeen[y/3][x/3] &= ^(1 << k)
				puzzle[y][x] = 0
            }
        }
        return finish        
    }

    backtrack(0)

	return puzzle
}