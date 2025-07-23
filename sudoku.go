package main

// import "fmt"
func isValid(grid [][]int, r int, c int, k int) bool {
	check_row := true
	check_col := true
    for i:=0; i < 9; i++ {
		if(k == grid[r][i]){
			not_in_row = false
			break
		}
	}
	for i:=0; i < 9; i++ {
		if(k == grid[i][c]){
			not_in_row = false
			break
		}
	}
	check_box := true
    for i := r / 3 * 3; i < r/3*3+3; i++ {
        for j := c / 3 * 3; j < c/3*3+3; j++ {
            if grid[i][j] == k {
                not_in_box = false
                break
            }
        }
    }

	return check_row && check_col && check_box;

}

func SolveSudoku(grid[][]int) [][]int{
	if(solve(grid,0,0)){ // call solve
		return grid
	}
	return nil
	
}
// Backtracking algorithm
func solve(grid [][]int, r, c int) bool {
	if r == 9 { // done
		return true
	} else if c ==9 { // move to next row
		return solve(grid,r+1,0)
	} else if grid[r][c] != 0 { // already filled
		return solve(grid,r,c+1)
	} else { //0
		for k:=1 ; k <= 9;k++{
			if(isValid(grid,r,c,k)){ // check the valid number in the position
				grid[r][c] = k
				if (solve(grid,r,c+1)){ // call solve on the rest with c+1
					return true
				}
				grid[r][c] = 0
			}
		}
	}
	return false
}

