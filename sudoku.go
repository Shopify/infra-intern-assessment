package main

// This file uses a shortened version of Knuth's Algorithm X to solve a version of
// the exact cover problem specialized for Sudoku.

// constraint types
const (
	occupied = iota
	rowNum
	colNum
	boxNum
)

// sudoku constants
const (
	boxSize = 3
	n       = 9
)

// a constraint in the exact cover problem
// kind is one of occupied, rowNum, colNum, boxNum
// a, b are the "arguments" for the constraint
type constraint struct {
	kind, a, b int
}

// represents a cell in the grid along with its value
type point struct {
	x, y, value int
}

// a constraint c -> a point p -> whether p satisfies c
type coverConstraints map[constraint]map[point]bool

// a point p -> the set of constraints that are satisfied by p
type coverOptions map[point][]constraint

// initialize constraint/option set for the exact cover problem
func ExactCover() (coverConstraints, coverOptions) {
	constraints := make(coverConstraints)
	choices := make(coverOptions)

	// populate constraints/choices
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			// try every value for the cell
			for val := 1; val <= n; val++ {
				// get box index (0-8)
				b := (y/boxSize)*boxSize + (x / boxSize)
				p := point{x, y, val}
				choices[p] = []constraint{
					{occupied, x, y}, // number cannot exist if cell is already occupied
					{rowNum, y, val}, // number cannot exist if it's already in the row
					{colNum, x, val}, // number cannot exist if it's already in the col
					{boxNum, b, val}, // number cannot exist if it's already in the box
				}
				for _, con := range choices[p] {
					if _, ok := constraints[con]; !ok {
						constraints[con] = make(map[point]bool)
					}
					constraints[con][p] = true
				}
			}
		}
	}

	return constraints, choices
}

// Updates the constraints/choice sets in-place based on the given point
func selectPoint(constraints coverConstraints, choices coverOptions, p point) []coverConstraints {
	// cols is used for backtracking in deselectPoint
	cols := []coverConstraints{}

	// for each constraint con that is satisfied by p...
	for _, con := range choices[p] {
		// remove all points p2 that conflict with p and con
		for p2 := range constraints[con] {
			for _, con2 := range choices[p2] {
				if con != con2 {
					delete(constraints[con2], p2)
				}
			}
		}

		// save the state of the constraints for backtracking
		newCol := make(coverConstraints)
		for con2, points := range constraints {
			newCol[con2] = make(map[point]bool)
			for p, isSatisfied := range points {
				newCol[con2][p] = isSatisfied
			}
		}
		cols = append(cols, newCol)
		delete(constraints, con)
	}

	return cols
}

// Backtracks from a selectPoint operation (i.e. deselects p)
func deselectPoint(
	constraints coverConstraints,
	choices coverOptions,
	p point,
	cols []coverConstraints,
) {
	// for all constraints con satisfied by p, in reverse order of selection...
	for i := len(choices[p]) - 1; i >= 0; i-- {
		// restore state of constraint from cols
		con := choices[p][i]
		constraints[con] = cols[len(cols)-1][con]
		cols = cols[:len(cols)-1]
		for p2 := range constraints[con] {
			for _, con2 := range choices[p2] {
				if con != con2 {
					constraints[con2][p2] = true
				}
			}
		}
	}
}

// Recursively applies constraints until the puzzle is solved.
func solve(constraints coverConstraints, choices coverOptions, solution []point) [][]point {
	// no more constraints to satisfy
	if len(constraints) == 0 {
		return [][]point{solution}
	}

	// choose constraint with fewest remaining options
	smallestColumn := constraint{}
	minSize := -1
	for k, v := range constraints {
		if len(v) < minSize || minSize == -1 {
			smallestColumn = k
			minSize = len(v)
		}
	}

	// try all choices for the chosen constraint
	solutions := [][]point{}
	for p := range constraints[smallestColumn] {
		newSolution := append([]point{}, solution...)
		newSolution = append(newSolution, p)

		// apply constraints, recurse, then backtrack
		// add points from solve to solutions if it's a valid solution
		cols := selectPoint(constraints, choices, p)
		solutions = append(solutions, solve(constraints, choices, newSolution)...)
		deselectPoint(constraints, choices, p, cols)
	}

	return solutions
}

// Returns a solved version of the input grid. Assumes grid is a valid sudoku board
// -- that is, grid is a 9 by 9 matrix containing only elements from 0 to 9, where
// 0 represents an empty cell.
func SolveSudoku(grid [][]int) [][]int {
	constraints, choices := ExactCover()

	// pre-select existing numbers in the grid
	for y, row := range grid {
		for x, cell := range row {
			if cell != 0 {
				selectPoint(constraints, choices, point{x, y, cell})
			}
		}
	}

	solutions := solve(constraints, choices, []point{})
	if len(solutions) > 0 {
		firstSolution := solutions[0]
		for _, p := range firstSolution {
			grid[p.y][p.x] = p.value
		}
	}

	return grid
}
