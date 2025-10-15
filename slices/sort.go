package slices

import (
	"cmp"
	"slices"
)

// Insert inserts an element into a sorted slice.
func Insert[S ~[]E, E cmp.Ordered](s S, e E, unique bool) S {
	i, found := slices.BinarySearch(s, e)
	if found && unique {
		return s
	}
	return slices.Insert(s, i, e)
}

// Delete deletes an element from a sorted slice.
func Delete[S ~[]E, E cmp.Ordered](s S, e E) S {
	i, found := slices.BinarySearch(s, e)
	if found {
		return slices.Delete(s, i, i+1)
	}

	return s
}

// Contains checks if a slice contains an element.
func Contains[T cmp.Ordered](s []T, e T) bool {
	_, found := slices.BinarySearch(s, e)

	return found
}
