package heap

import (
	"cmp"

	xcmp "github.com/gravitton/x/cmp"
)

// NewOrdered constructs a new Heap of ordered elements, using cmp.Compare.
func NewOrdered[E cmp.Ordered]() *Heap[E] {
	return New(cmp.Compare[E])
}

// NewComparable constructs a new Heap of comparable elements, using their Compare method.
func NewComparable[T any, E xcmp.Comparable[T]]() *Heap[E] {
	return New(func(x E, y E) int {
		return x.Compare(y)
	})
}
