package main

// IntSet the set of Ints
type IntSet struct {
	items map[int]bool
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *IntSet) Add(t int) *IntSet {
	if s.items == nil {
		s.items = make(map[int]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Clear removes all elements from the Set
func (s *IntSet) Clear() {
	s.items = make(map[int]bool)
}

// Delete removes the int from the Set and returns Has(int)
func (s *IntSet) Delete(item int) bool {
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

// Has returns true if the Set contains the int
func (s *IntSet) Has(item int) bool {
	_, ok := s.items[item]
	return ok
}

// Ints returns the int(s) stored
func (s *IntSet) Ints() []int {
	items := []int{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *IntSet) Size() int {
	return len(s.items)
}
