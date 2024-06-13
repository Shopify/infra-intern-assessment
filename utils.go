package main

func Copy[T any](s [][]T) [][]T {
	copy := make([][]T, len(s))

	for i := 0; i < len(s); i++ {
		copy[i] = make([]T, len(s[i]))

		for j, v := range s[i] {
			copy[i][j] = v
		}
	}

	return copy
}