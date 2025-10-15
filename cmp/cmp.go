package cmp

// Comparable is an interface for types that can be compared.
type Comparable[T any] interface {
	Compare(*T) int
	*T
}
