<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# hashset

```go
import "github.com/fufuok/utils/generic/hashset"
```

Package hashset provides an implementation of a hashset\.

<details><summary>Example</summary>
<p>

```go
package main

import (
	"fmt"
	g "github.com/fufuok/utils/generic"
	"github.com/fufuok/utils/generic/hashset"
)

func main() {
	set := hashset.New[string](3, g.Equals[string], g.HashString)
	set.Put("foo")
	set.Put("bar")
	set.Put("baz")

	fmt.Println(set.Has("foo"))
	fmt.Println(set.Has("quux"))
}
```

#### Output

```
true
false
```

</p>
</details>

## Index

- [type Set](<#type-set>)
  - [func New[K any](capacity uint64, equals g.EqualsFn[K], hash g.HashFn[K]) *Set[K]](<#func-new>)
  - [func (s *Set[K]) Copy() *Set[K]](<#func-setk-copy>)
  - [func (s *Set[K]) Each(fn func(key K))](<#func-setk-each>)
  - [func (s *Set[K]) Has(val K) bool](<#func-setk-has>)
  - [func (s *Set[K]) Put(val K)](<#func-setk-put>)
  - [func (s *Set[K]) Remove(val K)](<#func-setk-remove>)
  - [func (s *Set[K]) Size() int](<#func-setk-size>)


## type [Set](<https://gitee.com/fufuok/utils/blob/master/generic/hashset/set.go#L10-L12>)

Set implements a hashset\, using the hashmap as the underlying storage\.

```go
type Set[K any] struct {
    // contains filtered or unexported fields
}
```

### func [New](<https://gitee.com/fufuok/utils/blob/master/generic/hashset/set.go#L15>)

```go
func New[K any](capacity uint64, equals g.EqualsFn[K], hash g.HashFn[K]) *Set[K]
```

New returns an empty hashset\.

### func \(\*Set\[K\]\) [Copy](<https://gitee.com/fufuok/utils/blob/master/generic/hashset/set.go#L50>)

```go
func (s *Set[K]) Copy() *Set[K]
```

Copy returns a copy of this set\.

### func \(\*Set\[K\]\) [Each](<https://gitee.com/fufuok/utils/blob/master/generic/hashset/set.go#L43>)

```go
func (s *Set[K]) Each(fn func(key K))
```

Each calls 'fn' on every item in the set in no particular order\.

### func \(\*Set\[K\]\) [Has](<https://gitee.com/fufuok/utils/blob/master/generic/hashset/set.go#L27>)

```go
func (s *Set[K]) Has(val K) bool
```

Has returns true only if 'val' is in the set\.

### func \(\*Set\[K\]\) [Put](<https://gitee.com/fufuok/utils/blob/master/generic/hashset/set.go#L22>)

```go
func (s *Set[K]) Put(val K)
```

Put adds 'val' to the set\.

### func \(\*Set\[K\]\) [Remove](<https://gitee.com/fufuok/utils/blob/master/generic/hashset/set.go#L33>)

```go
func (s *Set[K]) Remove(val K)
```

Remove removes 'val' from the set\.

### func \(\*Set\[K\]\) [Size](<https://gitee.com/fufuok/utils/blob/master/generic/hashset/set.go#L38>)

```go
func (s *Set[K]) Size() int
```

Size returns the number of elements in the set\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)