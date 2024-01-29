// Adapted from https://pkg.go.dev/github.com/golang-collections/collections/stack
package main

type Coords struct {
	row int
	col int
}

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value Coords
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() Coords {
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() Coords {
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value Coords) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}
