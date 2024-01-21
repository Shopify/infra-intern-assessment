package main

import (
	"fmt"
)


// Best solution for solving sudoku problems is 'Dancing Links'
// Further info can be found here: https://en.wikipedia.org/wiki/Dancing_Links
// Research paper: https://arxiv.org/pdf/cs/0011047.pdf
// Unfortunately I cannot solve it like so in the reccomended 1-2 hour time
// Rather I will be implementing a simpler and slower approach: backtracking
// We can further optimize this backtracking via Memoization

// These variables will store all values that are not empty/0
var rows_visit [9][9][10]bool;
var cols_visit [9][9][10]bool;
var box_visit [9][9][10]bool;

  


// SolveSudoku takes a 2d array of 9x9 size
// Treats array as sudoku board that needs to be solved
// SolveSudoku [][]int -> [][]int
func SolveSudoku(sudoku_grid [][]int) [][]int {


	// checks all visited rows, cols and 3x3 boxes visited
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			curr_num := sudoku_grid[i][j]
		  if curr_num != 0 {
			var box_index int = (i/3)*3 + j/3;
			rows_visit[i][j][curr_num] = true;
			cols_visit[i][j][curr_num] = true;
			box_visit[box_index][j][curr_num] = true;
		  }
		}
	  }

	backTrack(sudoku_grid);
	showBoard(sudoku_grid);



	return sudoku_grid;
}


// this is the main backtracking function
// Using DFS principals
// runs as so:
/*
	Loop through the board

	if the current "cell" is empty(0)

		fire another forloop iterating from 1-9(all possible values)
		check if it is a valid answer via the isValid function below
			if so then modify the grid
			recall backTrack() and repeat while return true or false if 
			the sudoku is valid
			if false try another number else return true

*/
// backTrack [][]int -> bool
func backTrack(sudoku_grid [][]int) bool{
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if sudoku_grid[i][j]==0{
				//means we need to add something

				for curr_num := 1; curr_num <= 9; curr_num++ {
					var box_index int = (i / 3) * 3 + j / 3;

					//check if curr_num works on the sudoku
					if isValid( sudoku_grid, i, j, curr_num){
						sudoku_grid[i][j] = curr_num;
						// add value to grid and then check if it works

						// add value to rows, cols and boxes visited
						rows_visit[i][j][curr_num] = true;
						cols_visit[i][j][curr_num] = true;
						box_visit[box_index][j][curr_num] = true;
						

						if backTrack(sudoku_grid){
							return true;
						}else{
							// if it doesn't work as a number revert back to 0
							// then try again with another number between 1-10;
							sudoku_grid[i][j] = 0;

							//remove the values from the memoization map
							rows_visit[i][j][curr_num] = false;
							cols_visit[i][j][curr_num] = false;
							box_visit[box_index][j][curr_num] = false;
						}
						

					}
				}
				// if no value fits then we return false :(
				return false;
				
			}

		}
	}
	return true;
}


// Checks to see if test_num has already existed within 
// The rows, cols, and boxes visited
// then checks if exists in sudoku grid
// isValid int int int -> bool
func isValid(grid [][]int, x_coords int, y_coords  int, test_num int) bool {
	var box_index int = (x_coords / 3) * 3 + y_coords / 3;

	if rows_visit[x_coords][y_coords][test_num]{
		return false;
	}
	if cols_visit[x_coords][y_coords][test_num]{
		return false;
	}
	if box_visit[box_index][y_coords][test_num]{
		return false;
	}
	
	
	//check if exists in the entire grid's x_coords row
	for i := 0; i < 9; i++ {
		if (grid[x_coords][i] == test_num ){
		  return false;
		}
	  }
	
	  //check if exists in the entire grid's y_coords row
	  for i := 0; i < 9; i++ {
		if (grid[i][y_coords] == test_num) {
		  return false;
		}
	  }
	
	  // Check the 3x3 grid in sudoku
	  startRow  :=  x_coords - x_coords%3;
	  startCol  :=  y_coords - y_coords%3;
	  for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
		  if grid[i+startRow][j+startCol] == test_num {
			return false;
		  }
		}
	  }
	  
	  return true
  }



// display existing board
// Note: not needed in project 
// made for debugging purposes
// showBoard [][]int -> void
func showBoard(grid [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}




