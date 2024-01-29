package main

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Hello!! Thank you for taking a look at my solution. In scouting strategies, I came across two contenders, 
// Backtracking and Dancing Links. Backtracking was the more common solution pairing efficiency with simplicity, 
// while Dancing Links was faster in many cases but involved combinatorics, posing sudoku as an exact cover problem.
// I had a read through the Donald Knuth Dancing Links paper and it really set combinatorics on my radar for future study! 
// However, to better keep with the project's scope I chose Backtracking using maps for a simplified solution.
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Gives an integer list of candidate values at a certain location
func candidates(grid [][]int, row, col int) []int {
	candidateList := make([]int, 0)
    candidateMap := make(map[int]bool)
	
	// Use integer truncation to derive the starting coordinates of the sub-grid which the target number is in
	subGridRow, subGridCol := (row / 3) * 3, (col / 3) * 3

	// Check if the target num is already present in its column or row
    for i := 0; i < 9; i++ {
        candidateMap[grid[row][i]] = true
        candidateMap[grid[i][col]] = true
    }

	// Check if the target num is already present in its sub-grid
    for i := subGridRow; i < subGridRow + 3; i++ {
        for j := subGridCol; j < subGridCol + 3; j++ {
            candidateMap[grid[i][j]] = true
        }
    }

	// Append all options not in the map into a possible candidates list
    for i := 1; i < 10; i++ {
		// If a key isn't present in the map, candidateMap[i] returns false 
        if !candidateMap[i] {
            candidateList = append(candidateList, i)
        }
    }

    return candidateList
}

func SolveTheSudoku(grid [][]int, row, col int) bool {
	// Recursive base case reached if all rows successfully traversed
    if row == 9 {
        return true
	// If end of row reached, reset col to 0 and move down a row
    } else if col == 9 {
        return SolveTheSudoku(grid, row + 1, 0)
	// Skip square if a number already present
    } else if grid[row][col] != 0 {
        return SolveTheSudoku(grid, row, col + 1)
    } else {
		for _, num := range candidates(grid, row, col) {
			// Try a candidate and start down a recursive path
            grid[row][col] = num
            if SolveTheSudoku(grid, row, col + 1) {
                return true
            }

			// If recursive calls lead down wrong path, reset to 0 
            grid[row][col] = 0
        }

        return false
    }
}

func SolveSudoku(grid [][]int) [][]int {
    SolveTheSudoku(grid, 0, 0)
    return grid
}
