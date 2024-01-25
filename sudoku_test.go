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

// source: https://valeur.org/sudoku/hard/a/free-printable-hard-sudoku-with-the-answer-13553.html
func TestSolveSudokuHard(t *testing.T) {
	input := [][]int{
		{0, 0, 6, 3, 0, 7, 0, 0, 0},
		{0, 0, 4, 0, 0, 0, 0, 0, 5},
		{1, 0, 0, 0, 0, 6, 0, 8, 2},
		{2, 0, 5, 0, 3, 0, 1, 0, 6},
		{0, 0, 0, 2, 0, 0, 3, 0, 0},
		{9, 0, 0, 0, 7, 0, 0, 0, 4},
		{0, 5, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 1, 0, 9, 0, 4, 0},
	}

	expected := [][]int{
		{5, 8, 6, 3, 2, 7, 4, 9, 1},
		{7, 2, 4, 8, 9, 1, 6, 3, 5},
		{1, 9, 3, 5, 4, 6, 7, 8, 2},
		{2, 4, 5, 9, 3, 8, 1, 7, 6},
		{8, 6, 7, 2, 1, 4, 3, 5, 9},
		{9, 3, 1, 6, 7, 5, 8, 2, 4},
		{4, 5, 2, 7, 6, 3, 9, 1, 8},
		{3, 1, 9, 4, 8, 2, 5, 6, 7},
		{6, 7, 8, 1, 5, 9, 2, 4, 3},
	}

	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}

func TestSolveUnsolvableSudoku(t *testing.T) {
	input := [][]int{
		{6, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solved := SolveSudoku(input)

	if solved != nil {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", nil, solved)
	}
}

func TestSolveInvalidSudoku(t *testing.T) {
	input := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2, 0},
		{6, 7, 2, 1, 9, 5, 3, 4, 8, 0},
		{1, 9, 8, 3, 4, 2, 5, 6, 7, 0},
		{8, 5, 9, 7, 6, 1, 4, 2, 3, 0},
		{4, 2, 6, 8, 5, 3, 7, 9, 1, 0},
		{7, 1, 3, 9, 2, 4, 8, 5, 6, 0},
		{9, 6, 1, 5, 3, 7, 2, 8, 4, 0},
		{2, 8, 7, 4, 1, 9, 6, 3, 5, 0},
	}

	solved := SolveSudoku(input)

	if solved != nil {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", nil, solved)
	}
}
