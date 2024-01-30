package main
/*
Hello! My name is Ayush Satyavarpu, and I am very interested in the Shopify position.
Here is my implementation of the sudoku solver. It uses a backtracking approach with additional
optimizations. 

Please contact me at ayushsatyavarpu@gmail.com or visit my website at ayushsat.github.io

Infrastructure wise, I have experience working with numerous tooling,
like AWS, GCP, Jenkins, Docker, and IIS Express.
*/

//checks if inserting num at row, col into arr creates a valid state
func canInsert(arr [][]int, row int, col int, num int) bool{
	//loop through the column and the row to see if the number is already there
	for x := 0; x<9;x++{
		if arr[row][x] == num{
			return false
		}
		if arr[x][col] == num{
			return false
		}
	}

	//calculate the bounds for the small 3x3 grid, and loop through to ensure the number isn't present
	leftBound := int(col/3) * 3
	topBound := int(row/3) * 3
	for r := topBound;r<topBound + 3;r++ {
		for c := leftBound;c<leftBound + 3;c++ {
			if arr[r][c] == num{
				return false
			}
		}	
	}
	return true
}

//starts performing the solve starting at (row, col)
func SolveSudokuInternal(arr [][]int, row int, col int) bool{
	//if you hit the bottom without returning false, it's solved!
	if row == 8 && col == 8{
		return true;
	}
	//shorthand notation to quickly wrap indices around the grid
	if col == 9{
		col = 0
		row++
	}
	//if the value already exists, just skip
	if arr[row][col] > 0{
		return SolveSudokuInternal(arr, row, col + 1)
	}
	//otherwise, loop through every possible value, set it to arr[row][col]
	for val := 1;val <10;val++{
		//first check if its safe to insert
		if canInsert(arr, row, col, val){
			arr[row][col] = val
			if SolveSudokuInternal(arr, row, col + 1){
				return true
			}
			//if it wasn't, then remove it
			arr[row][col] = 0
		}
	}
	//no possible value is valid, just return false
	return false
}

func SolveSudoku(arr [][]int) [][]int {
	SolveSudokuInternal(arr, 0, 0)
	return arr
}

func main() {
    
}