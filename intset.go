package main

// IntSet the set of Ints

type void struct{}

var member void

type IntSet map[int]void

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *IntSet) Add(item int) *IntSet {
	(*s)[item] = member
	return s
}

// Delete removes the int from the Set and returns Has(int)
func (s *IntSet) Delete(item int) bool {
	_, ok := (*s)[item]
	if ok {
		delete(*s, item)
	}
	return ok
}

// Has returns true if the Set contains the int
func (s *IntSet) Has(item int) bool {
	_, ok := (*s)[item]
	return ok
}

// Ints returns the int(s) stored
func (s *IntSet) Ints() []int {
	items := []int{}
	for i := range *(s) {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *IntSet) Size() int {
	return len(*s)
}
