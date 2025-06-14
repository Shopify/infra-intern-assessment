package main

import (
	"reflect"
	"strconv"
	"testing"
)

// constraints state that input grid will have exactly one solution
func TestSolveSudoku(t *testing.T) {
	runSudokuTest(t, [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}, [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	})
	runSudokuTest(t,
		convertStringToGrid("3.542.81.4879.15.6.29.5637485.793.416132.8957.74.6528.2413.9.655.867.192.965124.8"),
		convertStringToGrid("365427819487931526129856374852793641613248957974165283241389765538674192796512438"),
	)
	runSudokuTest(t,
		convertStringToGrid("..2.3...8.....8....31.2.....6..5.27..1.....5.2.4.6..31....8.6.5.......13..531.4.."),
		convertStringToGrid("672435198549178362831629547368951274917243856254867931193784625486592713725316489"),
	)
	runSudokuTest(t,
		convertStringToGrid(".......1.4.........2...........5.4.7..8...3....1.9....3..4..2...5.1........8.6..."),
		convertStringToGrid("693784512487512936125963874932651487568247391741398625319475268856129743274836159"),
	)
}

// generic testing function
func runSudokuTest(t *testing.T, input, expected [][]int) {
	// t.Logf("Input:\n%v\n", input)
	solved := SolveSudoku(input)
	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Expected:\n%v\n\nGot:\n%v\n", expected, solved)
	}
}

// helper function to convert sudoku string input to grid
func convertStringToGrid(s string) [][]int {
	grid := make([][]int, 9)
	for i := range grid {
		grid[i] = make([]int, 9)
		for j := range grid[i] {
			if s[i*9+j] >= '1' && s[i*9+j] <= '9' {
				grid[i][j], _ = strconv.Atoi(string(s[i*9+j]))
			}
		}
	}
	return grid
}

// helper function to convert grid to string
func convertGridToString(grid [][]int) string {
	s := ""
	for i := range grid {
		for j := range grid[i] {
			s += strconv.Itoa(grid[i][j])
		}
	}
	return s
}
