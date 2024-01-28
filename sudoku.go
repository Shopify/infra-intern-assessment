package main

const gridSize = 9 // the grid size is 9x9 and does not change

func SolveSudoku(grid [][]int) [][]int {
	return grid
}

// findEmptyCell is a helper function that returns the position of the cell if it contains contains 0
// meaning that the cell is empty
func findEmptyCell(grid [][]int) []int {
    for row := 0; row < gridSize; row++ {
        for col := 0; col < gridSize; col++ {
            if grid[row][col] == 0 {
                return []int{row, col}
            }
        }
    }
    return nil
}