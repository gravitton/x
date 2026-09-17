<div align="center" width="100%">

<a href="https://github.com/gravitton">
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/gravitton/x/refs/heads/main/docs/images/logo-dark.svg">
  <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/gravitton/x/refs/heads/main/docs/images/logo-light.svg">
  <img alt="Gravitton x" src="https://raw.githubusercontent.com/gravitton/x/refs/heads/main/docs/images/logo-light.svg" width="300">
</picture>
</a>

[![Latest Stable Version][ico-release]][link-release]
[![Build Status][ico-workflow]][link-workflow]
[![Coverage Status][ico-coverage]][link-coverage]
[![Go Dev Reference][ico-go-dev-reference]][link-go-dev-reference]
[![Software License][ico-license]][link-licence]

Collection of extra packages that implement some useful features that are missing in the standard library

<hr>

</div>


## Features

- **Heap** – generic min-heap over `container/heap`, with ordered and comparable constructors.
- **Queue** – generic FIFO queue backed by a slice.
- **Slices** – `Map`, plus `Insert`, `Delete`, and `Contains` on sorted slices.
- **Duration** – `time.Duration` that marshals as `"200ms"` in JSON, TOML, and YAML.
- **Zero dependencies** – only the standard library at runtime.

## Installation

```shell
go get github.com/gravitton/x
```

## Usage

Each package mirrors its standard-library namesake and is meant to be imported next to it.

Heap over any ordered type:

```go
import "github.com/gravitton/x/container/heap"

h := heap.NewOrdered[int]()
h.Push(3)
h.Push(1)
h.Push(2)

h.Peek() // 1
h.Pop()  // 1
h.Len()  // 2
```

Heap over a custom type with a comparison function:

```go
type task struct {
	name     string
	priority int
}

h := heap.New(func(a, b task) int {
	return cmp.Compare(a.priority, b.priority)
})
```

Priority queue with in-place updates. `SetIndex` records each element's position so `Fix` and `Remove` can find it:

```go
type item struct {
	name     string
	priority int
	index    int
}

func (i *item) Compare(other *item) int {
	return cmp.Compare(i.priority, other.priority)
}

pq := heap.NewComparable[item]()
pq.SetIndex(func(i *item, index int) {
	i.index = index
})

orange := &item{name: "orange", priority: 1}
pq.Push(orange)
pq.Push(&item{name: "apple", priority: 2})

orange.priority = 5
pq.Fix(orange.index)

pq.Pop().name // "apple"
```

Queue:

```go
import "github.com/gravitton/x/container/queue"

q := queue.New[string]()
q.Push("a")
q.Push("b")

q.Peek()  // "a"
q.Pop()   // "a"
q.Len()   // 1
q.Clear()
```

Slices:

```go
import "github.com/gravitton/x/slices"

slices.Map([]float64{1.1, 2.6}, math.Round) // []float64{1, 3}
slices.Map([]int{1, 2}, strconv.Itoa)       // []string{"1", "2"}

s := []int{1, 3, 5}
s = slices.Insert(s, 4, true) // []int{1, 3, 4, 5}, sorted; true skips duplicates
s = slices.Delete(s, 3)       // []int{1, 4, 5}
slices.Contains(s, 4)         // true, by binary search
```

Duration in config files:

```go
import xtime "github.com/gravitton/x/time"

type Config struct {
	Timeout xtime.Duration `json:"timeout"`
}

var cfg Config
json.Unmarshal([]byte(`{"timeout":"200ms"}`), &cfg)

cfg.Timeout.String()                      // "200ms"
cfg.Timeout.Equal(200 * time.Millisecond) // true
```

Full reference: [pkg.go.dev][link-go-dev-reference].

## Conventions

**Heap:** `New` takes a `cmp.Compare`-style function returning negative, zero, or positive. `NewOrdered` uses `cmp.Compare` for `cmp.Ordered` types. `NewComparable` works with pointer elements whose type implements `Compare(*T) int`, the `cmp.Comparable` constraint from this module. `Pop` and `Peek` panic on an empty heap; check `Empty` first.

**Queue:** `Pop` and `Peek` panic on an empty queue. `Slice` returns the backing slice, not a copy, so modifying it invalidates the queue. The same holds for `Heap.Slice`.

**Sorted slices:** `Insert`, `Delete`, and `Contains` assume the slice is already sorted in ascending order and use binary search. Results are unspecified for unsorted input.

**Map:** A nil input maps to a nil slice, so nil and empty stay distinguishable through a mapping.

**Duration:** Marshalling goes through `encoding.TextMarshaler` and `encoding.TextUnmarshaler`, so it works with `encoding/json` and any TOML or YAML library that honours those interfaces. Parsing accepts anything `time.ParseDuration` does.

## Credits

- [Tomáš Novotný](https://github.com/tomas-novotny)
- [All Contributors][link-contributors]

## License

The MIT License (MIT). Please see [License File][link-licence] for more information.


[ico-license]:              https://img.shields.io/github/license/gravitton/x.svg?style=flat-square&colorB=blue
[ico-workflow]:             https://img.shields.io/github/actions/workflow/status/gravitton/x/main.yml?branch=main&style=flat-square
[ico-release]:              https://img.shields.io/github/v/release/gravitton/x?style=flat-square&colorB=blue
[ico-go-dev-reference]:     https://img.shields.io/badge/go.dev-reference-blue?style=flat-square
[ico-coverage]:             https://img.shields.io/coverallsCoverage/github/gravitton/x?style=flat-square

[link-author]:              https://github.com/gravitton
[link-release]:             https://github.com/gravitton/x/releases
[link-contributors]:        https://github.com/gravitton/x/contributors
[link-licence]:             ./LICENSE.md
[link-changelog]:           ./CHANGELOG.md
[link-workflow]:            https://github.com/gravitton/x/actions
[link-go-dev-reference]:    https://pkg.go.dev/github.com/gravitton/x
[link-coverage]:            https://coveralls.io/github/gravitton/x
