<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# orderedmap

```go
import "github.com/fufuok/utils/orderedmap"
```

## Index

- [type ByPair](<#type-bypair>)
  - [func (a ByPair) Len() int](<#func-bypair-len>)
  - [func (a ByPair) Less(i, j int) bool](<#func-bypair-less>)
  - [func (a ByPair) Swap(i, j int)](<#func-bypair-swap>)
- [type OrderedMap](<#type-orderedmap>)
  - [func New() *OrderedMap](<#func-new>)
  - [func (o *OrderedMap) Delete(key string)](<#func-orderedmap-delete>)
  - [func (o *OrderedMap) Get(key string) (interface{}, bool)](<#func-orderedmap-get>)
  - [func (o *OrderedMap) Keys() []string](<#func-orderedmap-keys>)
  - [func (o OrderedMap) MarshalJSON() ([]byte, error)](<#func-orderedmap-marshaljson>)
  - [func (o *OrderedMap) Set(key string, value interface{})](<#func-orderedmap-set>)
  - [func (o *OrderedMap) SetEscapeHTML(on bool)](<#func-orderedmap-setescapehtml>)
  - [func (o *OrderedMap) Sort(lessFunc func(a *Pair, b *Pair) bool)](<#func-orderedmap-sort>)
  - [func (o *OrderedMap) SortKeys(sortFunc ...func(keys []string))](<#func-orderedmap-sortkeys>)
  - [func (o *OrderedMap) UnmarshalJSON(b []byte) error](<#func-orderedmap-unmarshaljson>)
  - [func (o *OrderedMap) Values() map[string]interface{}](<#func-orderedmap-values>)
- [type Pair](<#type-pair>)
  - [func (kv *Pair) Key() string](<#func-pair-key>)
  - [func (kv *Pair) Value() interface{}](<#func-pair-value>)


## type ByPair

```go
type ByPair struct {
    Pairs    []*Pair
    LessFunc func(a *Pair, j *Pair) bool
}
```

### func \(ByPair\) Len

```go
func (a ByPair) Len() int
```

### func \(ByPair\) Less

```go
func (a ByPair) Less(i, j int) bool
```

### func \(ByPair\) Swap

```go
func (a ByPair) Swap(i, j int)
```

## type OrderedMap

```go
type OrderedMap struct {
    // contains filtered or unexported fields
}
```

### func New

```go
func New() *OrderedMap
```

### func \(\*OrderedMap\) Delete

```go
func (o *OrderedMap) Delete(key string)
```

### func \(\*OrderedMap\) Get

```go
func (o *OrderedMap) Get(key string) (interface{}, bool)
```

### func \(\*OrderedMap\) Keys

```go
func (o *OrderedMap) Keys() []string
```

### func \(OrderedMap\) MarshalJSON

```go
func (o OrderedMap) MarshalJSON() ([]byte, error)
```

### func \(\*OrderedMap\) Set

```go
func (o *OrderedMap) Set(key string, value interface{})
```

### func \(\*OrderedMap\) SetEscapeHTML

```go
func (o *OrderedMap) SetEscapeHTML(on bool)
```

### func \(\*OrderedMap\) Sort

```go
func (o *OrderedMap) Sort(lessFunc func(a *Pair, b *Pair) bool)
```

Sort the map using your sort func

### func \(\*OrderedMap\) SortKeys

```go
func (o *OrderedMap) SortKeys(sortFunc ...func(keys []string))
```

SortKeys Sort the map keys using your sort func

### func \(\*OrderedMap\) UnmarshalJSON

```go
func (o *OrderedMap) UnmarshalJSON(b []byte) error
```

### func \(\*OrderedMap\) Values

```go
func (o *OrderedMap) Values() map[string]interface{}
```

## type Pair

```go
type Pair struct {
    // contains filtered or unexported fields
}
```

### func \(\*Pair\) Key

```go
func (kv *Pair) Key() string
```

### func \(\*Pair\) Value

```go
func (kv *Pair) Value() interface{}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)