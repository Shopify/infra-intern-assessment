package main

/// Solves a sudoku grid, returning a copy of the answer.
///
/// The grid must have exactly one solution.
func SolveSudoku(g [][]int) [][]int {
	// Copy to a fixed size grid.
	var c grid
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c[i][j] = g[i][j]
		}
	}
	if !solve(&c) {
		panic("The sudoku is unsolvable!")
	}
	// Copy back to a slice of slice.
	ans := make([][]int, 9)
	for i := 0; i < 9; i++ {
		ans[i] = c[i][0:9]
	}
	return ans
}

/// A 9x9 sudoku grid.
type grid [9][9]int

/// Solves the given grid.
///
/// If the grid is solvable, returns true and mutates the grid. Otherwise, returns false.
func solve(g *grid) bool {
	// Check if already solved.
	if checkSolved(g) {
		return true
	}
	// Otherwise, we can just guess. We use backtracking.
	return guess(g)
}

/// Returns true if the given grid is solved.
func checkSolved(g *grid) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if g[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

/// Returns an array representing the possible values a cell could take according
/// to the existing clues. Note that the array is off by 1.
func generatePossibilities(g *grid, i int, j int) [9]bool {
	// Keep track of the possible values that the cell can take.
	var pos [9]bool
	for k := 0; k < 9; k++ {
		pos[k] = true
	}
	// Go over the row and column.
	for k := 0; k < 9; k++ {
		if g[i][k] != 0 {
			pos[g[i][k] - 1] = false
		}
		if g[k][j] != 0 {
			pos[g[k][j] - 1] = false
		}
	}
	// Go over the current square by rounding down to the nearest multiple of 3.
	y := i / 3 * 3
	x := j / 3 * 3
	for k := y; k < y + 3; k++ {
		for l := x; l < x + 3; l++ {
			if g[k][l] != 0 {
				pos[g[k][l] - 1] = false
			}
		}
	}
	return pos
}

/// Attempts to solve the grid by guessing. Returns true and mutates the grid if successful.
///
/// The grid must not already be solved.
func guess(g *grid) bool {
	// Find the first empty cell.
	i := 0
	j := 0
	out:
	for ; i < 9; i++ {
		for j = 0; j < 9; j++ {
			// Try to fill in the empty cells by guessing.
			if g[i][j] == 0 {
				break out
			}
		}
	}
	// Generate possibilities that the cell could take.
	pos := generatePossibilities(g, i, j)
	for k := 0; k < 9; k++ {
		if pos[k] {
			// Attempt to set this value and solve.
			g[i][j] = k + 1
			if solve(g) {
				return true
			}
			g[i][j] = 0
		}
	}
	// If nothing worked, then it is unsolvable.
	return false
}
