package main

import (
	"reflect"
	"testing"
)

// The original provided test case
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

// A solved sudoku
func TestFilledSudoku(t *testing.T) {
	input := [][]int{
		{9, 4, 2, 7, 6, 5, 3, 8, 1},
		{3, 8, 1, 9, 4, 2, 7, 5, 6},
		{6, 5, 7, 3, 1, 8, 2, 4, 9},
		{5, 2, 3, 8, 7, 1, 9, 6, 4},
		{7, 9, 4, 5, 2, 6, 8, 1, 3},
		{8, 1, 6, 4, 3, 9, 5, 7, 2},
		{4, 3, 9, 6, 5, 7, 1, 2, 8},
		{2, 7, 8, 1, 9, 4, 6, 3, 5},
		{1, 6, 5, 2, 8, 3, 4, 9, 7},
	}

	expected := [][]int{
		{9, 4, 2, 7, 6, 5, 3, 8, 1},
		{3, 8, 1, 9, 4, 2, 7, 5, 6},
		{6, 5, 7, 3, 1, 8, 2, 4, 9},
		{5, 2, 3, 8, 7, 1, 9, 6, 4},
		{7, 9, 4, 5, 2, 6, 8, 1, 3},
		{8, 1, 6, 4, 3, 9, 5, 7, 2},
		{4, 3, 9, 6, 5, 7, 1, 2, 8},
		{2, 7, 8, 1, 9, 4, 6, 3, 5},
		{1, 6, 5, 2, 8, 3, 4, 9, 7},
	}

	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}

// Expert difficulty - 6 backtracks
func TestExpertSudoku(t *testing.T) {
	input := [][]int{
		{0, 0, 0, 0, 0, 2, 0, 7, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 8},
		{4, 8, 0, 7, 0, 6, 0, 2, 0},
		{1, 0, 0, 0, 4, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 4, 0, 3},
		{0, 9, 0, 0, 0, 0, 7, 0, 0},
		{3, 0, 0, 2, 0, 7, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 3, 1},
		{6, 0, 0, 0, 0, 9, 2, 0, 0},
	}

	expected := [][]int{
		{5, 6, 9, 3, 8, 2, 1, 7, 4},
		{7, 3, 2, 4, 1, 5, 9, 6, 8},
		{4, 8, 1, 7, 9, 6, 3, 2, 5},
		{1, 7, 3, 6, 4, 8, 5, 9, 2},
		{2, 5, 6, 9, 7, 1, 4, 8, 3},
		{8, 9, 4, 5, 2, 3, 7, 1, 6},
		{3, 1, 5, 2, 6, 7, 8, 4, 9},
		{9, 2, 7, 8, 5, 4, 6, 3, 1},
		{6, 4, 8, 1, 3, 9, 2, 5, 7},
	}

	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}

// Expert difficulty - 2 backtracks
func TestExpert2Sudoku(t *testing.T) {
	input := [][]int{
		{0, 5, 7, 0, 2, 0, 8, 0, 0},
		{0, 2, 0, 0, 6, 0, 0, 0, 3},
		{8, 0, 0, 0, 7, 0, 5, 1, 0},
		{0, 7, 0, 1, 0, 5, 0, 6, 0},
		{0, 8, 0, 0, 0, 0, 0, 0, 0},
		{4, 0, 5, 0, 0, 0, 0, 0, 0},
		{1, 4, 2, 0, 0, 0, 0, 0, 0},
		{0, 3, 0, 2, 0, 0, 0, 5, 0},
		{0, 0, 0, 0, 0, 0, 6, 0, 0},
	}

	expected := [][]int{
		{3, 5, 7, 4, 2, 1, 8, 9, 6},
		{9, 2, 1, 5, 6, 8, 7, 4, 3},
		{8, 6, 4, 9, 7, 3, 5, 1, 2},
		{2, 7, 3, 1, 9, 5, 4, 6, 8},
		{6, 8, 9, 7, 4, 2, 1, 3, 5},
		{4, 1, 5, 8, 3, 6, 2, 7, 9},
		{1, 4, 2, 6, 5, 9, 3, 8, 7},
		{7, 3, 6, 2, 8, 4, 9, 5, 1},
		{5, 9, 8, 3, 1, 7, 6, 2, 4},
	}

	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}

// Medium difficulty - 0 backtracks
func TestMediumSudoku(t *testing.T) {
	input := [][]int{
		{0, 0, 0, 7, 0, 0, 2, 0, 0},
		{7, 0, 0, 0, 0, 5, 0, 3, 1},
		{0, 0, 2, 1, 3, 0, 0, 4, 5},
		{4, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 9, 0, 0, 8, 0, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 5, 0},
		{0, 0, 5, 0, 9, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 4, 5, 0, 0},
		{3, 8, 1, 2, 0, 0, 0, 0, 0},
	}

	expected := [][]int{
		{5, 1, 3, 7, 4, 6, 2, 8, 9},
		{7, 9, 4, 8, 2, 5, 6, 3, 1},
		{8, 6, 2, 1, 3, 9, 7, 4, 5},
		{4, 5, 8, 9, 6, 3, 1, 2, 7},
		{2, 7, 9, 5, 1, 8, 3, 6, 4},
		{1, 3, 6, 4, 7, 2, 9, 5, 8},
		{6, 4, 5, 3, 9, 1, 8, 7, 2},
		{9, 2, 7, 6, 8, 4, 5, 1, 3},
		{3, 8, 1, 2, 5, 7, 4, 9, 6},
	}
	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}

// Easy difficulty - 0 backtracks
func TestEasySudoku(t *testing.T) {
	input := [][]int{
		{0, 3, 0, 5, 7, 0, 0, 4, 6},
		{0, 0, 0, 1, 3, 0, 0, 0, 0},
		{0, 0, 6, 4, 0, 0, 0, 0, 9},
		{4, 7, 0, 9, 6, 0, 0, 0, 3},
		{0, 5, 0, 7, 0, 3, 0, 0, 0},
		{0, 0, 8, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 7, 9, 0, 5},
		{9, 0, 0, 6, 5, 4, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	expected := [][]int{
		{8, 3, 2, 5, 7, 9, 1, 4, 6},
		{7, 9, 4, 1, 3, 6, 8, 5, 2},
		{5, 1, 6, 4, 2, 8, 7, 3, 9},
		{4, 7, 1, 9, 6, 5, 2, 8, 3},
		{2, 5, 9, 7, 8, 3, 4, 6, 1},
		{3, 6, 8, 2, 4, 1, 5, 9, 7},
		{6, 4, 3, 8, 1, 7, 9, 2, 5},
		{9, 2, 7, 6, 5, 4, 3, 1, 8},
		{1, 8, 5, 3, 9, 2, 6, 7, 4},
	}
	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}
