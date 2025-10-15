package slices

// Map applies a function to each element of a slice and returns a new slice.
func Map[S ~[]E, E any, T any](input S, fn func(E) T) []T {
	ts := make([]T, len(input))
	for i, v := range input {
		ts[i] = fn(v)
	}

	return ts
}
