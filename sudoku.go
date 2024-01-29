package main

func GetSolns(input [][]int, loc_x int, loc_y int) []int {
	var possible_solns []int
	possible_solns = append(possible_solns, 1)
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

	//find all possible numbers in current position

	var possible_solns = GetSolns(input, loc_x, loc_y)
	if len(possible_solns) == 0 {
		input[loc_y][loc_x] = 0
		return false // Base case: the given input has no solution
	}

	//runsolve on each number until one solves
	for _, soln := range possible_solns {
		input[loc_y][loc_x] = soln
		if RunSolve(input, loc_x, loc_y) {
			return true
		}
	}

	//if runsolve cant solve, move onto the next possible number.
	//if all numbers are used, reset value at position to zero and return false
	input[loc_y][loc_x] = 0
	return false

	//if a number works, return true
}

func SolveSudoku(input [][]int) [][]int {
	RunSolve(input, 0, 0)
	return input
}
