package main

func GetSolns(input [][]int, loc_x int, loc_y int) map[int]bool {
	var possible_solns map[int]bool = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}

	for i := 0; i < 9; i++ {
		delete(possible_solns, input[loc_y][i])
		delete(possible_solns, input[i][loc_x])
	}

	box_x := (loc_x / 3) * 3
	box_y := (loc_y / 3) * 3 // Rounds to the upper left corner of 3x3 box

	for dx := 0; dx < 3; dx++ {
		for dy := 0; dy < 3; dy++ {
			delete(possible_solns, input[box_y+dy][box_x+dx])
		}
	}

	return possible_solns
}

// Returns solves the puzzle starting from the given location (for efficiency) and
// returns true. Otherwise, returns false and the puzzle is unmodified.
func RunSolve(input [][]int, loc_x int, loc_y int) bool {

	//move x, y to the first blank
	for input[loc_y][loc_x] != 0 {
		loc_x++
		if loc_x == 9 {
			loc_x = 0
			loc_y++
			if loc_y == 9 {
				return true // Base case: there are no more values to solve for
			}
		}
	}

	var possible_solns map[int]bool = GetSolns(input, loc_x, loc_y)
	if len(possible_solns) == 0 {
		input[loc_y][loc_x] = 0
		return false // Base case: the given input has no solution
	}

	for soln := range possible_solns {
		input[loc_y][loc_x] = soln
		if RunSolve(input, loc_x, loc_y) {
			return true
		}
	}

	input[loc_y][loc_x] = 0
	return false
}

func SolveSudoku(input [][]int) [][]int {
	RunSolve(input, 0, 0)
	return input
}
