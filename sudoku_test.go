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

	// From sudoku book
	/*
		   func TestSolveSudoku(t *testing.T) {
		   	input := [][]int{
		   		{0, 0, 6, 0, 0, 0, 5, 0, 8},
		   		{1, 0, 2, 3, 8, 0, 0, 0, 4},
		   		{0, 0, 0, 2, 0, 0, 1, 9, 0},
		   		{0, 0, 0, 0, 6, 3, 0, 4, 5},
		   		{0, 6, 3, 4, 0, 5, 8, 7, 0},
		   		{5, 8, 0, 9, 2, 0, 0, 0, 0},
		   		{0, 4, 7, 0, 0, 4, 0, 0, 0},
		   		{2, 0, 0, 0, 9, 8, 4, 0, 7},
		   		{4, 0, 9, 0, 0, 0, 3, 0, 0},
		   	}

		   	expected := [][]int{
		   		{3, 9, 6, 7, 4, 1, 5, 2, 8},
		   		{1, 5, 2, 3, 8, 9, 7, 6, 4},
		   		{8, 7, 4, 2, 5, 6, 1, 9, 3},
		   		{7, 2, 1, 8, 6, 3, 9, 4, 5},
		   		{9, 6, 3, 4, 1, 5, 8, 7, 2},
		   		{5, 4, 8, 9, 2, 7, 6, 3, 1},
		   		{6, 8, 7, 1, 3, 4, 2, 5, 9},
		   		{2, 3, 5, 6, 9, 8, 4, 1, 7},
		   		{4, 1, 9, 5, 7, 2, 3, 8, 6},
		   	}

		   	solved := SolveSudoku(input)

		   	if !reflect.DeepEqual(solved, expected) {
		   		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
		   	}
		}

	*/
}
