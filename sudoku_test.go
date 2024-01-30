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

	// I added a few other testcases here, generated from an online Sudoku generator.
	// credits: https://sudokutodo.com/generator
	medium_test := [][]int {
		{0, 7, 8, 3, 5, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 9, 3, 5, 0},
		{6, 5, 3, 2, 0, 0, 0, 7, 9},
		{0, 0, 0, 7, 9, 8, 0, 0, 5},
		{0, 9, 0, 0, 0, 5, 1, 3, 2},
		{0, 6, 4, 0, 0, 2, 0, 0, 0},
		{0, 0, 9, 6, 2, 4, 0, 0, 3},
		{0, 2, 6, 0, 1, 0, 0, 8, 7},
		{3, 1, 0, 9, 8, 0, 6, 2, 0},
	}

	medium_expected := [][]int {
		{9, 7, 8, 3, 5, 6, 2, 4, 1},
		{1, 4, 2, 8, 7, 9, 3, 5, 6},
		{6, 5, 3, 2, 4, 1, 8, 7, 9},
		{2, 3, 1, 7, 9, 8, 4, 6, 5},
		{8, 9, 7, 4, 6, 5, 1, 3, 2},
		{5, 6, 4, 1, 3, 2, 7, 9, 8},
		{7, 8, 9, 6, 2, 4, 5, 1, 3},
		{4, 2, 6, 5, 1, 3, 9, 8, 7},
		{3, 1, 5, 9, 8, 7, 6, 2, 4},
	}

	medium_solution := SolveSudoku(medium_test)

	if !reflect.DeepEqual(medium_solution, medium_expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", medium_expected, medium_solution)
	}

	hard_test := [][]int {
		{0, 5, 6, 0, 4, 1, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 5, 0, 6},
		{8, 0, 0, 0, 0, 6, 0, 2, 1},
		{0, 0, 3, 0, 8, 7, 2, 0, 0},
		{0, 0, 4, 5, 0, 0, 8, 9, 0},
		{9, 8, 0, 0, 0, 0, 1, 0, 3},
		{1, 3, 2, 0, 0, 8, 0, 4, 0},
		{4, 0, 0, 1, 0, 0, 9, 0, 8},
		{7, 0, 0, 4, 6, 0, 0, 0, 2},
	}

	hard_expected := [][]int {
		{3, 5, 6, 2, 4, 1, 7, 8, 9},
		{2, 4, 1, 8, 7, 9, 5, 3, 6},
		{8, 7, 9, 3, 5, 6, 4, 2, 1},
		{5, 1, 3, 9, 8, 7, 2, 6, 4},
		{6, 2, 4, 5, 1, 3, 8, 9, 7},
		{9, 8, 7, 6, 2, 4, 1, 5, 3},
		{1, 3, 2, 7, 9, 8, 6, 4, 5},
		{4, 6, 5, 1, 3, 2, 9, 7, 8},
		{7, 9, 8, 4, 6, 5, 3, 1, 2},
	}

	hard_solution := SolveSudoku(hard_test)

	if !reflect.DeepEqual(hard_solution, hard_expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", hard_expected, hard_solution)
	}
}
