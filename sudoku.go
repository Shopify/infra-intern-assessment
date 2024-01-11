package main

import (
	"sort"
)

// Filters out the given val from the arr
func filter(arr []int, val int) []int {

	var vals []int
	for _, v := range arr {
		if v != val {
			vals = append(vals, v)
		}
	}
	return vals
}

// updates the domains of all the peers of the given cell, aka removes the given value from the domains of all peers
func update(peers map[[2]int][][2]int, domains map[[2]int][]int, cell [2]int, value int) {

	for _, neighbor := range peers[cell] {
		domains[neighbor] = filter(domains[neighbor], value)
	}
}

func backtrack(board [][]int, peers map[[2]int][][2]int, domains map[[2]int][]int, variables [][2]int) {

}

func SolveSudoku(board [][]int) [][]int {

	// Dimensions of the board
	const M = 9
	const N = 9

	domain := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Variables are defined as all cells that don't have values yet
	// Each variable is identified with it's coordinate (row, col)
	var variables [][2]int

	domains := make(map[[2]int][]int)
	peers := make(map[[2]int][][2]int)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			domains[[2]int{i, j}] = domain
			peers[[2]int{i, j}] = [][2]int{}
			// All other cells in same column and row
			for k := 0; k < M; k++ {
				if k != i {
					peers[[2]int{i, j}] = append(peers[[2]int{i, j}], [2]int{k, j})
				}
				if k != j {
					peers[[2]int{i, j}] = append(peers[[2]int{i, j}], [2]int{i, k})
				}
			}
			// We iterate through all cells in the same group as the current cell
			// We then append these cells to the peers
			for ro := i - i%3; ro < ((i/3 + 1) * 3); ro++ {
				for co := j - j%3; co < ((j/3 + 1) * 3); co++ {
					if ro != i || co != j {
						peers[[2]int{i, j}] = append(peers[[2]int{i, j}], [2]int{ro, co})
					}
				}
			}
		}
	}

	// Iterate through every cell and perform constraint propagation aka we reduce the domains of all peers to the cell

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] != 0 {
				domains[[2]int{i, j}] = []int{board[i][j]}
				update(peers, domains, [2]int{i, j}, board[i][j])
			} else {
				variables = append(variables, [2]int{i, j})
			}
		}
	}
	// Sorts the variables based off the size of domain for each variable

	sort.SliceStable(variables, func(i, j int) bool {
		return len(domains[variables[i]]) < len(domains[variables[j]])
	})

	return board
}
