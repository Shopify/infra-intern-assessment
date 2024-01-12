package main

// representation of a sudoku board
type SudokuBoard struct {
	board    [][]int
	num_rows int
	num_cols int
}

// creat new sudoku board from a 2D slice of integers
func NewSudokuBoard(board [][]int) *SudokuBoard {
	return &SudokuBoard{
		board:    board,
		num_rows: len(board),
		num_cols: len(board[0]),
	}
}

func (s *SudokuBoard) IsValidMove(r, c, num int) bool {
	// first check rows
	for row_idx := 0; row_idx < 9; row_idx++ {
		if s.board[row_idx][c] == num {
			return false
		}
	}

	// check cols
	for col_idx := 0; col_idx < 9; col_idx++ {
		if s.board[r][col_idx] == num {
			return false
		}
	}

	// check the corresponding 3x3 box
	box_row := r - r%3
	box_col := c - c%3

	for row_idx := 0; row_idx < 3; row_idx++ {
		for col_idx := 0; col_idx < 3; col_idx++ {
			if s.board[box_row+row_idx][box_col+col_idx] == num {
				return false
			}
		}
	}

	return true
}

func (s *SudokuBoard) Solve(cur_r, cur_c int) bool {
	if cur_r == s.num_rows-1 && cur_c == s.num_cols {
		return true
	}

	// end of column, move to next row and first column
	if cur_c > s.num_cols-1 {
		cur_r++
		if cur_r >= s.num_rows {
			return true // reached the end of the board
		}
		cur_c = 0
	}

	// current spot is already populated, go to next box
	if s.board[cur_r][cur_c] != 0 {
		return s.Solve(cur_r, cur_c+1)
	}

	for possible_num := 1; possible_num <= 9; possible_num++ {
		if s.IsValidMove(cur_r, cur_c, possible_num) {
			// recursively check if this move leads to a solution
			// if it does, return true, otherwise revert the change
			s.board[cur_r][cur_c] = possible_num

			if s.Solve(cur_r, cur_c+1) {
				return true
			} else {
				s.board[cur_r][cur_c] = 0
			}

		}
	}

	return false
}

func SolveSudoku(board [][]int) [][]int {
	sudoku := NewSudokuBoard(board)

	// currently we don't use the return value of Solve()
	// but we could use it to check if the board is solvable
	// and perform some error handling
	_ = sudoku.Solve(0, 0)

	return board
}
