package main


// Returns solves the puzzle starting from the given location (for efficiency) and
// returns true. Otherwise, returns false and the puzzle is unmodified.
func RunSolve(input *[][]int, loc_x int, loc_y int) bool {
	//move x, y to the first blank

	//find all possible numbers in current position

	//runsolve on each number until one solves

	//if runsolve cant solve, move onto the next possible number.
	//if all numbers are used, reset position to zero and return false

	//if a number works, return true
	return true;
}

func SolveSudoku(input [][]int) [][]int {
	RunSolve(&input, 0, 0)
	return input
}
