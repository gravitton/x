package slices

import (
	"testing"

	"github.com/gravitton/assert"
)

func TestInsert(t *testing.T) {
	ss := []int{1, 3, 5}

	s1 := Insert(ss, 6, false)
	assert.NotSame(t, s1, ss)
	assert.Equal(t, s1, []int{1, 3, 5, 6})

	s2 := Insert(s1, 6, true)
	assert.Same(t, s2, s1)
	assert.Equal(t, s2, []int{1, 3, 5, 6})

	s3 := Insert(ss, 2, false)
	assert.Equal(t, s3, []int{1, 2, 3, 5})
}

func TestDelete(t *testing.T) {
	ss := []int{1, 3, 5}

	s1 := Delete(ss, 2)
	assert.Same(t, s1, ss)
	assert.Equal(t, s1, []int{1, 3, 5})

	s2 := Delete(ss, 3)
	assert.Equal(t, s2, []int{1, 5})
}

func TestContains(t *testing.T) {
	ss := []int{1, 3, 5}

	assert.True(t, Contains(ss, 3))
	assert.False(t, Contains(ss, 2))
}
