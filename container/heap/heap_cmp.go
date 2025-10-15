package heap

import (
	"cmp"

	xcmp "github.com/gravitton/x/cmp"
)

// NewOrdered constructs a new Heap with a ordered elements
func NewOrdered[E cmp.Ordered]() *Heap[E] {
	return &Heap[E]{sliceHeap[E]{cmp: cmp.Compare[E]}}
}

// NewComparable constructs a new Heap with a comparable elements
func NewComparable[T any, E xcmp.Comparable[T]]() *Heap[E] {
	return &Heap[E]{sliceHeap[E]{cmp: func(x E, y E) int {
		return x.Compare(y)
	}}}
}
