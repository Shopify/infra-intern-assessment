package main

import (
	"reflect"
	"testing"
)

// Table-driven tests implementation
func TestSolveSudoku(t *testing.T) {
	// Define struct for test cases
	type testCase struct {
		input    [][]int
		expected [][]int
	}

	// Create a slice of test cases
	tests := []testCase{
		{
			input : [][]int{
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
			expected : [][]int{
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
			input : [][]int{
				{0, 0, 0, 2, 6, 0, 7, 0, 1},
				{6, 8, 0, 0, 7, 0, 0, 9, 0},
				{1, 9, 0, 0, 0, 4, 5, 0, 0},
				{8, 2, 0, 1, 0, 0, 0, 4, 0},
				{0, 0, 4, 6, 0, 2, 9, 0, 0},
				{0, 5, 0, 0, 0, 3, 0, 2, 8},
				{0, 0, 9, 3, 0, 0, 0, 7, 4},
				{0, 4, 0, 0, 5, 0, 0, 3, 6},
				{7, 0, 3, 0, 1, 8, 0, 0, 0},
			},
			expected : [][]int{
				{4, 3, 5, 2, 6, 9, 7, 8, 1},
				{6, 8, 2, 5, 7, 1, 4, 9, 3},
				{1, 9, 7, 8, 3, 4, 5, 6, 2},
				{8, 2, 6, 1, 9, 5, 3, 4, 7},
				{3, 7, 4, 6, 8, 2, 9, 1, 5},
				{9, 5, 1, 7, 4, 3, 6, 2, 8},
				{5, 1, 9, 3, 2, 6, 8, 7, 4},
				{2, 4, 8, 9, 5, 7, 1, 3, 6},
				{7, 6, 3, 4, 1, 8, 2, 5, 9},
			},
		},
	}

	// Iterate over test cases individually
	for _, test := range tests {
		solved := SolveSudoku(test.input)
		if !reflect.DeepEqual(solved, test.expected) {
			t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", test.expected, solved)
		}
	}
	
}