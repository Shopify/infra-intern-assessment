package main

func SolveSudoku(grid [][]int) [][]int {
	SolveSudokuRecursive(grid)
	return grid
}

func GetPossibleValues(grid [][]int, coord [2]int) map[int] bool { 
	possible_values := make(map[int]bool)
	for i := 1; i <= 9; i++ { 
		possible_values[i] = true
	}
	for i := 0; i < 9; i++ {
		delete(possible_values, grid[coord[0]][i])
	}
	for i := 0; i < 9; i++ {
		delete(possible_values, grid[i][coord[1]])
	}
	for i := 0; i < 3; i++ { 
		for j := 0; j < 3; j++ { 
			delete(possible_values, grid[(coord[0] / 3) * 3 + i][(coord[1] / 3) * 3 + j])
		}
	}
	return possible_values
}

func SolveSudokuRecursive(grid [][]int) bool { 
	// check if the grid is solved
	zero_exists := false
	for i := 0; i < 9; i++ { 
		for j := 0; j < 9; j++ { 
			if grid[i][j] == 0 {
				zero_exists = true 
				break
			}
		}
	}
	if !zero_exists { 
		return true
	}

	set_squares_exist := true
	var possible_values [9][9]map[int]bool
	for set_squares_exist {
		set_squares_exist = false
		// find the possible values that each square could have 
		for i := 0; i < 9; i++ { 
			for j := 0; j < 9; j++ { 
				if grid[i][j] != 0 {
					continue 
				}
				possible_values[i][j] = GetPossibleValues(grid, [2]int{i, j})
			}
		}

		// if a square has 0 possible values, we can't solve this grid
		// if a square has only 1 possible value, that value is "set", and we can fill in that value
		for i := 0; i < 9; i++ { 
			for j := 0; j < 9; j++ {
				if grid[i][j] != 0  { 
					continue
				}
				if len(possible_values[i][j]) == 0 {
					return false
				} else if len(possible_values[i][j]) == 1 { 
					set_squares_exist = true
					for set_value, _ := range possible_values[i][j] { 
						grid[i][j] = set_value 
					}
				}
			}
		}
	}

	// iterate through the possible values of one particular square
	var coord_to_iterate_through [2]int
	zero_exists = false
	for i := 0; i < 9; i++ { 
		for j := 0; j < 9; j++ { 
			if grid[i][j] == 0 { 
				coord_to_iterate_through = [2]int{i, j}
				zero_exists = true
				break
			}
		}
	}
	if !zero_exists { 
		return true
	}
	for value, _ := range possible_values[coord_to_iterate_through[0]][coord_to_iterate_through[1]] { 
		grid[coord_to_iterate_through[0]][coord_to_iterate_through[1]] = value
		valid := SolveSudokuRecursive(grid)
		if valid { 
			return true
		}
	}
	return false
}