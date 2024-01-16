package main

import (
	"reflect"
	"testing"
)

// TestSolveSudoku is a test function for the SolveSudoku function.
// Each test case in the function is defined with a name, input board and the expected solved board.
func TestSolveSudoku(t *testing.T) {
	// testCases is an array of anonymous structs for each test case.
	var testCases = []struct {
		name     string  // name of the test case
		input    [][]int //Sudoku board that needs to be solved
		expected [][]int // The board that we expect after solving the Sudoku
	}{
		{
			name: "Normal Puzzle",
			input: [][]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			expected: [][]int{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9},
			},
		},
		{
			name: "Hard Puzzle",
			input: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 3, 0, 8, 5},
				{0, 0, 1, 0, 2, 0, 0, 0, 0},
				{0, 0, 0, 5, 0, 7, 0, 0, 0},
				{0, 0, 4, 0, 0, 0, 1, 0, 0},
				{0, 9, 0, 0, 0, 0, 0, 0, 0},
				{5, 0, 0, 0, 0, 0, 0, 7, 3},
				{0, 0, 2, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 4, 0, 0, 0, 9},
			},
			expected: [][]int{
				{9, 8, 7, 6, 5, 4, 3, 2, 1},
				{2, 4, 6, 1, 7, 3, 9, 8, 5},
				{3, 5, 1, 9, 2, 8, 7, 4, 6},
				{1, 2, 8, 5, 3, 7, 6, 9, 4},
				{6, 3, 4, 8, 9, 2, 1, 5, 7},
				{7, 9, 5, 4, 6, 1, 8, 3, 2},
				{5, 1, 9, 2, 8, 6, 4, 7, 3},
				{4, 7, 2, 3, 1, 9, 5, 6, 8},
				{8, 6, 3, 7, 4, 5, 2, 1, 9},
			},
		},
		{
			name: "Another Puzzle",
			input: [][]int{
				{0, 2, 0, 6, 0, 8, 0, 0, 0},
				{5, 8, 0, 0, 0, 9, 7, 0, 0},
				{0, 0, 0, 0, 4, 0, 0, 0, 0},
				{3, 7, 0, 0, 0, 0, 5, 0, 0},
				{6, 0, 0, 0, 0, 0, 0, 0, 4},
				{0, 0, 8, 0, 0, 0, 0, 1, 3},
				{0, 0, 0, 0, 2, 0, 0, 0, 0},
				{0, 0, 9, 8, 0, 0, 0, 3, 6},
				{0, 0, 0, 3, 0, 6, 0, 9, 0},
			},
			expected: [][]int{
				{1, 2, 3, 6, 7, 8, 9, 4, 5},
				{5, 8, 4, 2, 3, 9, 7, 6, 1},
				{9, 6, 7, 1, 4, 5, 3, 2, 8},
				{3, 7, 2, 4, 6, 1, 5, 8, 9},
				{6, 9, 1, 5, 8, 3, 2, 7, 4},
				{4, 5, 8, 7, 9, 2, 6, 1, 3},
				{8, 3, 6, 9, 2, 4, 1, 5, 7},
				{2, 1, 9, 8, 5, 7, 4, 3, 6},
				{7, 4, 5, 3, 1, 6, 8, 9, 2},
			},
		},
		{
			name: "Invalid Puzzle",
			input: [][]int{
				{5, 5, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			//In case an invalid puzzle is given, SolveSudoku returns a board full of zeroes.
			expected: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}

	//Iterate over each testcase
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solved := SolveSudoku(tc.input)
			//If solved board does not match the expected board, the test fails, then print the expected and actual boards
			if !reflect.DeepEqual(solved, tc.expected) {
				t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", tc.expected, solved)
			}
		})
	}
}
