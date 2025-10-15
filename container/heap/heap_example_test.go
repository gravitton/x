package heap

import (
	"cmp"
	"fmt"
)

type priorityItem struct {
	value    string
	priority int
	index    int
}

func (i *priorityItem) Compare(item *priorityItem) int {
	return cmp.Compare(i.priority, item.priority)
}

func (i *priorityItem) setIndex(index int) {
	i.index = index
}

func ExampleHeap_priorityQueue() {
	pq := NewComparable[priorityItem]()
	pq.SetIndex((*priorityItem).setIndex)

	pq.Push(&priorityItem{value: "banana", priority: 3})
	pq.Push(&priorityItem{value: "apple", priority: 2})
	pq.Push(&priorityItem{value: "pear", priority: 4})

	orange := &priorityItem{value: "orange", priority: 1}
	pq.Push(orange)
	orange.priority = 5
	pq.Fix(orange.index)

	for pq.Len() > 0 {
		item := pq.Pop()
		fmt.Printf("%.2d:%s\n", item.priority, item.value)
	}

	// Output:
	// 02:apple
	// 03:banana
	// 04:pear
	// 05:orange
}
