package main

type cell struct {
	i     int
	j     int
	value int

	//Basically a set, init using public newCell function
	possibleSol map[int]bool
}

func NewCell(i int, j int, val int) cell {
	var set map[int]bool
	if val == 0 {
		set = make(map[int]bool)
		for i := 1; i < 10; i++ {
			set[i] = true
		}
	}

	return cell{
		i:           i,
		j:           j,
		value:       val,
		possibleSol: set,
	}
}

func (c *cell) removeSol(v int) {
	c.possibleSol[v] = false
}

func (c *cell) removeSolSlice(v []int) {
	for _, v := range v {
		c.removeSol(v)
	}
}

func (c *cell) PossibleSol() []int {
	res := []int{}
	for k, v := range c.possibleSol {
		if v {
			res = append(res, k)
		}
	}
	return res
}


func transformToCells(in [][]int) [][]cell {
	board := make([][]cell, len(in))
	for i := range in {
		for j := range in[i] {
			board[i] = append(board[i], NewCell(i, j, in[i][j]))
		}
	}
	return board
}