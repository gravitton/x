package queue

import (
	"testing"

	"github.com/gravitton/assert"
)

func TestQueue(t *testing.T) {
	q := New[int]()

	assert.Equal(t, q.Len(), 0)
	assert.True(t, q.Empty())

	q.Push(2)
	q.Push(3)
	q.Push(1)

	assert.Equal(t, q.Len(), 3)
	assert.False(t, q.Empty())

	assert.Equal(t, q.Peek(), 2)
	assert.Equal(t, q.Len(), 3)

	assert.Equal(t, q.Pop(), 2)
	assert.Equal(t, q.Len(), 2)

	s := q.Slice()
	assert.Equal(t, s, []int{3, 1})
	q.Clear()

	assert.Equal(t, s, []int{0, 0})
	assert.Equal(t, q.Len(), 0)
	assert.True(t, q.Empty())
}
