package slices

import (
	"testing"

	"github.com/gravitton/assert"
)

func TestMap(t *testing.T) {
	si := []int{1, 2, 3}
	s1 := Map(si, func(i int) int {
		return i
	})

	assert.Equal(t, s1, si)
	assert.NotSame(t, s1, si)

	s2 := Map([]int{1, 2, 3}, func(i int) int {
		return 10 + i*2
	})

	assert.Equal(t, s2, []int{12, 14, 16})

	var nilSlice []int
	assert.True(t, Map(nilSlice, func(i int) int {
		return i
	}) == nil)
	assert.False(t, Map([]int{}, func(i int) int {
		return i
	}) == nil)
}
