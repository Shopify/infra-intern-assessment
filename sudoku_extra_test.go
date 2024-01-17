package main

import (
	"reflect"
	"testing"
)

// I made this extra file because the instructions did not specify that we could modify sudoku_test.go,
// but I wanted to add some extra tests for certainty.

func TestSolveSudokuExtra(t *testing.T) {

	input := [][][]int{
		{
			{7, 9, 0, 0, 0, 0, 1, 0, 2},
			{1, 3, 4, 2, 8, 0, 0, 9, 0},
			{0, 5, 0, 1, 9, 4, 8, 0, 3},
			{3, 0, 0, 0, 4, 9, 0, 1, 6},
			{0, 6, 1, 8, 0, 3, 0, 0, 9},
			{0, 7, 9, 0, 1, 2, 5, 0, 0},
			{8, 2, 3, 7, 5, 1, 0, 0, 0},
			{0, 1, 0, 4, 0, 0, 2, 8, 7},
			{6, 0, 7, 9, 2, 8, 3, 0, 0},
		},
		{
			{5, 0, 6, 0, 0, 0, 8, 0, 1},
			{9, 8, 1, 5, 0, 2, 0, 0, 7},
			{7, 3, 0, 0, 6, 1, 2, 0, 0},
			{6, 5, 0, 9, 0, 3, 1, 2, 4},
			{0, 4, 0, 6, 5, 0, 9, 0, 3},
			{8, 9, 0, 1, 2, 4, 0, 5, 6},
			{3, 7, 9, 0, 4, 0, 5, 0, 0},
			{2, 6, 5, 7, 0, 8, 3, 0, 9},
			{4, 0, 8, 0, 0, 5, 6, 7, 0},
		},
		{
			{8, 0, 1, 3, 4, 5, 0, 6, 0},
			{6, 3, 9, 0, 1, 0, 5, 4, 2},
			{7, 4, 5, 9, 2, 0, 8, 1, 3},
			{1, 7, 6, 4, 3, 9, 2, 5, 8},
			{0, 5, 4, 0, 8, 7, 1, 0, 6},
			{2, 0, 3, 5, 6, 1, 0, 9, 0},
			{3, 1, 7, 0, 9, 4, 6, 2, 5},
			{4, 9, 8, 6, 0, 2, 3, 7, 1},
			{5, 0, 2, 0, 7, 3, 4, 8, 0},
		},
		{
			{0, 2, 7, 0, 1, 8, 4, 9, 3},
			{1, 4, 3, 5, 7, 0, 6, 2, 8},
			{9, 6, 8, 4, 2, 3, 1, 7, 5},
			{0, 7, 2, 1, 9, 0, 8, 0, 0},
			{0, 1, 0, 8, 6, 2, 9, 5, 7},
			{0, 5, 9, 3, 4, 7, 0, 0, 0},
			{4, 3, 0, 9, 5, 0, 7, 8, 2},
			{7, 8, 1, 0, 3, 4, 5, 6, 9},
			{2, 9, 0, 7, 8, 6, 3, 4, 0},
		},
		{
			{9, 8, 4, 0, 1, 3, 0, 6, 2},
			{7, 5, 0, 6, 8, 0, 1, 3, 0},
			{3, 6, 1, 5, 4, 0, 8, 0, 7},
			{2, 4, 0, 3, 7, 8, 6, 1, 5},
			{0, 1, 0, 4, 0, 6, 0, 7, 3},
			{6, 3, 7, 2, 0, 1, 0, 8, 0},
			{0, 0, 8, 9, 0, 0, 3, 2, 1},
			{1, 0, 0, 8, 2, 5, 7, 4, 0},
			{4, 0, 6, 1, 3, 7, 0, 5, 0},
		},
	}

	expected := [][][]int{
		{
			{7, 9, 8, 3, 6, 5, 1, 4, 2},
			{1, 3, 4, 2, 8, 7, 6, 9, 5},
			{2, 5, 6, 1, 9, 4, 8, 7, 3},
			{3, 8, 2, 5, 4, 9, 7, 1, 6},
			{5, 6, 1, 8, 7, 3, 4, 2, 9},
			{4, 7, 9, 6, 1, 2, 5, 3, 8},
			{8, 2, 3, 7, 5, 1, 9, 6, 4},
			{9, 1, 5, 4, 3, 6, 2, 8, 7},
			{6, 4, 7, 9, 2, 8, 3, 5, 1},
		},
		{
			{5, 2, 6, 4, 7, 9, 8, 3, 1},
			{9, 8, 1, 5, 3, 2, 4, 6, 7},
			{7, 3, 4, 8, 6, 1, 2, 9, 5},
			{6, 5, 7, 9, 8, 3, 1, 2, 4},
			{1, 4, 2, 6, 5, 7, 9, 8, 3},
			{8, 9, 3, 1, 2, 4, 7, 5, 6},
			{3, 7, 9, 2, 4, 6, 5, 1, 8},
			{2, 6, 5, 7, 1, 8, 3, 4, 9},
			{4, 1, 8, 3, 9, 5, 6, 7, 2},
		},
		{
			{8, 2, 1, 3, 4, 5, 9, 6, 7},
			{6, 3, 9, 7, 1, 8, 5, 4, 2},
			{7, 4, 5, 9, 2, 6, 8, 1, 3},
			{1, 7, 6, 4, 3, 9, 2, 5, 8},
			{9, 5, 4, 2, 8, 7, 1, 3, 6},
			{2, 8, 3, 5, 6, 1, 7, 9, 4},
			{3, 1, 7, 8, 9, 4, 6, 2, 5},
			{4, 9, 8, 6, 5, 2, 3, 7, 1},
			{5, 6, 2, 1, 7, 3, 4, 8, 9},
		},
		{
			{5, 2, 7, 6, 1, 8, 4, 9, 3},
			{1, 4, 3, 5, 7, 9, 6, 2, 8},
			{9, 6, 8, 4, 2, 3, 1, 7, 5},
			{6, 7, 2, 1, 9, 5, 8, 3, 4},
			{3, 1, 4, 8, 6, 2, 9, 5, 7},
			{8, 5, 9, 3, 4, 7, 2, 1, 6},
			{4, 3, 6, 9, 5, 1, 7, 8, 2},
			{7, 8, 1, 2, 3, 4, 5, 6, 9},
			{2, 9, 5, 7, 8, 6, 3, 4, 1},
		},
		{
			{9, 8, 4, 7, 1, 3, 5, 6, 2},
			{7, 5, 2, 6, 8, 9, 1, 3, 4},
			{3, 6, 1, 5, 4, 2, 8, 9, 7},
			{2, 4, 9, 3, 7, 8, 6, 1, 5},
			{8, 1, 5, 4, 9, 6, 2, 7, 3},
			{6, 3, 7, 2, 5, 1, 4, 8, 9},
			{5, 7, 8, 9, 6, 4, 3, 2, 1},
			{1, 9, 3, 8, 2, 5, 7, 4, 6},
			{4, 2, 6, 1, 3, 7, 9, 5, 8},
		},
	}

	for i, _ := range input {
		solved := SolveSudoku(input[i])

		if !reflect.DeepEqual(solved, expected[i]) {
			t.Errorf("Sudoku puzzle %v was not solved correctly. Expected:\n%v\n\nGot:\n%v", i, expected, solved)
		}

	}

}
