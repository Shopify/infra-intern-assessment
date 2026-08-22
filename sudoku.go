// Package main implements a Sudoku solver.
//
// The digits already used in each row, column and 3x3 box are tracked as 9-bit
// masks. That reduces "is this placement legal?" to a single bitwise AND, and
// makes counting a cell's candidates cheap enough to drive a minimum-remaining-
// values heuristic: the most constrained empty cell is always filled first,
// which prunes far more of the search tree than scanning in reading order.
package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

const (
	size    = 9 // side length of the grid
	boxSize = 3 // side length of a sub-grid
	// full has bits 1..9 set; bit d means "digit d is taken".
	full = 0b1111111110
)

// state is the grid being solved plus the usage masks that keep the constraint
// checks O(1).
type state struct {
	grid  [size][size]int
	rows  [size]uint16
	cols  [size]uint16
	boxes [size]uint16
	empty int // cells still to fill
}

// boxIndex maps a cell coordinate to the index of its 3x3 sub-grid.
func boxIndex(r, c int) int {
	return (r/boxSize)*boxSize + c/boxSize
}

// place writes digit d into (r, c) and marks it used in all three masks.
func (s *state) place(r, c, d int) {
	bit := uint16(1) << uint(d)
	s.grid[r][c] = d
	s.rows[r] |= bit
	s.cols[c] |= bit
	s.boxes[boxIndex(r, c)] |= bit
	s.empty--
}

// remove undoes place, emptying the cell again.
func (s *state) remove(r, c, d int) {
	bit := uint16(1) << uint(d)
	s.grid[r][c] = 0
	s.rows[r] &^= bit
	s.cols[c] &^= bit
	s.boxes[boxIndex(r, c)] &^= bit
	s.empty++
}

// candidates returns the mask of digits that may legally go in (r, c).
func (s *state) candidates(r, c int) uint16 {
	return full &^ (s.rows[r] | s.cols[c] | s.boxes[boxIndex(r, c)])
}

// solve fills every remaining cell by backtracking, reporting whether the
// puzzle could be completed.
func (s *state) solve() bool {
	if s.empty == 0 {
		return true
	}

	// Minimum remaining values: branch on the empty cell with the fewest legal
	// digits. One candidate means the cell is forced, so this doubles as
	// constraint propagation; zero means this branch is already dead.
	bestR, bestC, bestCount := -1, -1, size+1
	var bestMask uint16
	for r := 0; r < size && bestCount > 1; r++ {
		for c := 0; c < size; c++ {
			if s.grid[r][c] != 0 {
				continue
			}
			mask := s.candidates(r, c)
			if count := bits.OnesCount16(mask); count < bestCount {
				bestR, bestC, bestMask, bestCount = r, c, mask, count
				if count <= 1 {
					break // no other cell can be a better branch point
				}
			}
		}
	}

	if bestCount == 0 {
		return false // some empty cell can take no digit at all
	}

	// Iterate the set bits of bestMask, lowest first; mask &= mask-1 clears it.
	for mask := bestMask; mask != 0; mask &= mask - 1 {
		d := bits.TrailingZeros16(mask)
		s.place(bestR, bestC, d)
		if s.solve() {
			return true
		}
		s.remove(bestR, bestC, d)
	}
	return false
}

// newState loads grid into a state, reporting false if the grid is not 9x9,
// holds a value outside 0..9, or already breaks a Sudoku rule.
func newState(grid [][]int) (*state, bool) {
	if len(grid) != size {
		return nil, false
	}
	s := &state{empty: size * size}
	for r, row := range grid {
		if len(row) != size {
			return nil, false
		}
		for c, d := range row {
			if d == 0 {
				continue
			}
			if d < 1 || d > size {
				return nil, false
			}
			if s.candidates(r, c)&(1<<uint(d)) == 0 {
				return nil, false // digit repeats in its row, column or box
			}
			s.place(r, c, d)
		}
	}
	return s, true
}

// SolveSudoku solves the given 9x9 puzzle, in which 0 marks an empty cell, and
// returns the completed grid. The input grid is left untouched. It returns nil
// if the input is not a well-formed 9x9 puzzle or has no solution.
func SolveSudoku(grid [][]int) [][]int {
	s, ok := newState(grid)
	if !ok || !s.solve() {
		return nil
	}

	solved := make([][]int, size)
	for r := range solved {
		solved[r] = make([]int, size)
		copy(solved[r], s.grid[r][:])
	}
	return solved
}

// formatGrid renders a grid in the bracketed form used by the README.
func formatGrid(grid [][]int) string {
	var b strings.Builder
	b.WriteString("[\n")
	for i, row := range grid {
		cells := make([]string, len(row))
		for j, d := range row {
			cells[j] = strconv.Itoa(d)
		}
		b.WriteString("  [" + strings.Join(cells, ", ") + "]")
		if i < len(grid)-1 {
			b.WriteByte(',')
		}
		b.WriteByte('\n')
	}
	b.WriteString("]")
	return b.String()
}

func main() {
	puzzle := [][]int{
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

	solved := SolveSudoku(puzzle)
	if solved == nil {
		fmt.Println("puzzle has no solution")
		return
	}
	fmt.Println(formatGrid(solved))
}
