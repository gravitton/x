package queue

// Queue is generic FIFO queue.
type Queue[E any] struct {
	s []E
}

// New creates a new queue backed by a slice.
func New[E any]() *Queue[E] {
	return &Queue[E]{}
}

// Push adds an element to last position of the queue.
func (q *Queue[E]) Push(e E) {
	q.s = append(q.s, e)
}

// Pop removes and returns the first element from the queue.
// Pop panics if the queue is empty.
func (q *Queue[E]) Pop() E {
	current := q.s[0]

	var zero E
	q.s[0] = zero
	q.s = q.s[1:]

	return current
}

// Peek returns the first element in the queue.
// Peek panics if the queue is empty.
func (q *Queue[E]) Peek() E {
	return q.s[0]
}

// Len returns the number of elements in the queue.
func (q *Queue[E]) Len() int {
	return len(q.s)
}

// Empty returns true if the queue is empty.
func (q *Queue[E]) Empty() bool {
	return len(q.s) == 0
}

// Clear removes all elements from the queue.
func (q *Queue[E]) Clear() {
	clear(q.s)
	q.s = q.s[:0]
}

// Slice returns the underlying slice.
// The queue retains the returned slice, so altering the slice may break the invariants and invalidate the queue.
func (q *Queue[E]) Slice() []E {
	return q.s
}
