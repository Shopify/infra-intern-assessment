package main

import (
	"math/bits"
	"strings"
)

var DefaultSolver = &SudokuSolver{}

// SolveSudoku solves the given sudoku board with DefaultSolver.
func SolveSudoku(board Sudoku) Sudoku {
	return DefaultSolver.Solve(board)
}

type SudokuSolver struct {
	rows, cols [9]uint
	block      [3][3]uint
	empty      [][2]int
}

func (s *SudokuSolver) Solve(board Sudoku) Sudoku {
	s.reset()
	s.init(board)
	s.dfs(board, 0)
	return board
}

func (s *SudokuSolver) dfs(board Sudoku, n int) bool {
	if n == len(s.empty) {
		return true
	}

	i, j := s.empty[n][0], s.empty[n][1]
	mask := ^(s.rows[i] | s.cols[j] | s.block[i/3][j/3]) & 0b111111111

	for ; mask > 0; mask &= mask - 1 {
		k := bits.TrailingZeros(mask)
		s.flip(i, j, k)
		board[i][j] = k + 1
		if s.dfs(board, n+1) {
			return true
		}
		s.flip(i, j, k)
	}
	return false
}

func (s *SudokuSolver) reset() {
	s.rows = [9]uint{}
	s.cols = [9]uint{}
	s.block = [3][3]uint{}
	s.empty = nil
}

func (s *SudokuSolver) init(board Sudoku) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				s.empty = append(s.empty, [2]int{i, j})
			} else {
				s.flip(i, j, board[i][j]-1)
			}
		}
	}
}

func (s *SudokuSolver) flip(i, j, n int) {
	s.rows[i] ^= 1 << n
	s.cols[j] ^= 1 << n
	s.block[i/3][j/3] ^= 1 << n
}

// Sudoku type represents a 2D array to hold the Sudoku board.
type Sudoku [][]int

func (s Sudoku) String() string {
	sb := strings.Builder{}
	sb.WriteString("[\n")
	for i := 0; i < 9; i++ {
		sb.WriteString("  [")
		for j := 0; j < 9; j++ {
			sb.WriteByte(byte(s[i][j] - 1 + '1'))
			if j != 8 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("]\n")
	}
	sb.WriteString("]\n")
	return sb.String()
}
