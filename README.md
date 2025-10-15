# Extra

[![Latest Stable Version][ico-release]][link-release]
[![Build Status][ico-workflow]][link-workflow]
[![Coverage Status][ico-coverage]][link-coverage]
[![Go Report Card][ico-go-report-card]][link-go-report-card]
[![Go Dev Reference][ico-go-dev-reference]][link-go-dev-reference]
[![Software License][ico-license]][link-licence]

Collection of extra packages that implement some useful features that are missing in the standard library.


## Installation

```bash
go get github.com/gravitton/x
```


## Usage

### Heap

Generic Heap backed by slice implementation using `container/heap` package from existing [proposal](https://github.com/golang/go/issues/47632).

```go
package main

import (
	"cmp"
	"fmt"

	"github.com/gravitton/x/container/heap"
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

func main() {
	pq := heap.NewComparable[priorityItem]()
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

	// 02:apple 03:banana 04:pear 05:orange
}
```

### Queue

Generic queue implementation. It is just a wrapper around `[]T` with some extra methods.

```go
package main

import (
	"fmt"

	"github.com/gravitton/x/container/queue"
)

func main() {
	q := queue.New[int]()
	q.Push(2)
	q.Push(3)
	q.Push(1)

	for q.Len() > 0 {
		val := q.Pop()
		fmt.Println(val)
	}

	// 2 3 1
}
```

### Map

Generic Map method for slices.

```go
package main

import (
	"fmt"
	"math"

	"github.com/gravitton/x/slices"
)

func main() {
	sf := slices.Map([]float64{1.1, 2.6, 3.4}, math.Round)
	si := slices.Map(sf, func(x float64) int {
		return 10 + int(x)
	})

	fmt.Println(si)
	// [11, 13, 13]
}
```
## Credits

- [Tomáš Novotný](https://github.com/tomas-novotny)
- [All Contributors][link-contributors]


## License

The MIT License (MIT). Please see [License File][link-licence] for more information.


[ico-license]:              https://img.shields.io/github/license/gravitton/x.svg?style=flat-square&colorB=blue
[ico-workflow]:             https://img.shields.io/github/actions/workflow/status/gravitton/x/main.yml?branch=main&style=flat-square
[ico-release]:              https://img.shields.io/github/v/release/gravitton/x?style=flat-square&colorB=blue
[ico-go-dev-reference]:     https://img.shields.io/badge/go.dev-reference-blue?style=flat-square
[ico-go-report-card]:       https://goreportcard.com/badge/github.com/gravitton/x?style=flat-square
[ico-coverage]:             https://img.shields.io/coverallsCoverage/github/gravitton/x?style=flat-square

[link-author]:              https://github.com/gravitton
[link-release]:             https://github.com/gravitton/x/releases
[link-contributors]:        https://github.com/gravitton/x/contributors
[link-licence]:             ./LICENSE.md
[link-changelog]:           ./CHANGELOG.md
[link-workflow]:            https://github.com/gravitton/x/actions
[link-go-dev-reference]:    https://pkg.go.dev/github.com/gravitton/x
[link-go-report-card]:      https://goreportcard.com/report/github.com/gravitton/x
[link-coverage]:            https://coveralls.io/github/gravitton/x
