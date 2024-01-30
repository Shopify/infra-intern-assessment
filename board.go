package main

import "fmt"

const END_OF_GRID cell = 81

type cell uint8

type grid [][]int

func (c cell) row() int {
	return int(c) / 9
}

func (c cell) column() int {
	return int(c) % 9
}

func (c cell) String() string {
	return fmt.Sprintf("(%d, %d)", c.row(), c.column())
}

func (c cell) next() cell {
	if c >= END_OF_GRID {
		panic("next() called on cell greater than END_OF_GRID")
	}

	return c + 1
}

func (g grid) At(c cell) int {
	return g[c.row()][c.column()]
}

func (g grid) Set(c cell, v int) {
	g[c.row()][c.column()] = v
}

func (g grid) Reset(c cell) {
	g[c.row()][c.column()] = 0
}