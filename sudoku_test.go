package main

import (
	"reflect"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	input := [][]int{
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

	expected := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}

func TestSolveSudoku2(t *testing.T) {
	input := [][]int{
		{9, 3, 4, 0, 6, 0, 0, 5, 0},
		{0, 0, 6, 0, 0, 4, 9, 2, 3},
		{0, 0, 8, 9, 0, 0, 0, 4, 6},
		{8, 0, 0, 5, 4, 6, 0, 0, 7},
		{6, 0, 0, 0, 1, 0, 0, 0, 5},
		{5, 0, 0, 3, 9, 0, 0, 6, 2},
		{3, 6, 0, 4, 0, 1, 2, 7, 0},
		{4, 7, 0, 6, 0, 0, 5, 0, 0},
		{0, 8, 0, 0, 0, 0, 6, 3, 4},
	}

	expected := [][]int{
		{9, 3, 4, 1, 6, 2, 7, 5, 8},
		{1, 5, 6, 8, 7, 4, 9, 2, 3},
		{7, 2, 8, 9, 3, 5, 1, 4, 6},
		{8, 9, 2, 5, 4, 6, 3, 1, 7},
		{6, 4, 3, 2, 1, 7, 8, 9, 5},
		{5, 1, 7, 3, 9, 8, 4, 6, 2},
		{3, 6, 5, 4, 8, 1, 2, 7, 9},
		{4, 7, 9, 6, 2, 3, 5, 8, 1},
		{2, 8, 1, 7, 5, 9, 6, 3, 4},
	}

	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}
