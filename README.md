# Disjoint

[![PkgGoDev](https://pkg.go.dev/badge/github.com/oliverdaff/disjoint)](https://pkg.go.dev/github.com/oliverdaff/disjoint) [![Go Report Card](https://goreportcard.com/badge/github.com/oliverdaff/disjoint)](https://goreportcard.com/report/github.com/oliverdaff/disjoint) [![CircleCI](https://circleci.com/gh/oliverdaff/disjoint.svg?style=shield)](https://circleci.com/gh/oliverdaff/disjoint)

Disjoint package is a implementation of a UnionFind ([Disjoint-set](https://en.wikipedia.org/wiki/Disjoint-set_data_structure)).  Each element in the Disjoint-set is part of exactly one (non-overlapping) set.

The DSet supports two methods

*   Find: Which returns which subset the element is a part of.
* Union: Which joins two subset sets together into a single subset.

## API

__Create new DSet__

A new `DSet` is created using the `NewDSet` function, passing the elements to populate the set.

`NewDSet` returns a error if the elements set contains duplicates.

```go
import "disjoint"

dset, err := disjoint.NewDSet([]interface{}{1, 2, 3, 4})
```

__Add a new element__
Adds a new element to the the set.

```go
import "disjoint"

dset, err := disjoint.NewDSet([]interface{}{1, 2, 3, 4})
dset.Add(5)
```

__Find partition for element__
Returns the identifier the passed in element is part of.

```go
import (
    "disjoint"
)

dset, err := disjoint.NewDSet([]interface{}{1, 2, 3, 4})
p1, ok  := dset.FindPartition(1)
p2, ok := dset.FindPartition(2)
```

__Merge two sets__
Merges the sets the two elements are part of.

```go
import (
    "disjoint"
)

dset, err := disjoint.NewDSet([]interface{}{1, 2, 3, 4})
p1, ok  := dset.Merge(1,2)
```

__Check if disjoint__
Checks if two elements are part of the same set or not.

```go
import (
    "disjoint"
)

dset, err := disjoint.NewDSet([]interface{}{1, 2, 3, 4})
p1, ok  := dset.AreDisjoint(1,2)
```

## Tests
The tests can be invoked with `go test`

## License
MIT © Oliver Daff