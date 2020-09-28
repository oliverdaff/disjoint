// Package disjoint contains a UnionFind (Disjoint-set) data structure.
package disjoint

import (
	"fmt"
	"reflect"
)

type void struct{}

var member void

type set map[interface{}]void

// DSet is a Disjoint Set data structure that keeps track of a set of elements
// partitioned into a number of disjoint (nonoverlapping) subsets.
//
// DSet supports two operations
//
// - Find: Returs which subset a particular element is in. Find returns an item that can be
// compared with other return values to determine it the two elements are in the same subset.
// - Union: Join two subsets into a single subset.
type DSet struct {
	partitions map[interface{}]set
}

func NewDSet(initialSet []interface{}) (*DSet, error) {
	dset := DSet{
		make(map[interface{}]set),
	}
	for _, element := range initialSet {
		if _, ok := dset.partitions[element]; ok {
			return nil, fmt.Errorf("Element exists in partitions")
		}
		elements := make(set)
		elements[element] = member
		dset.partitions[element] = elements
	}
	return &dset, nil
}

func (ds *DSet) Add(element interface{}) bool {
	if _, ok := ds.partitions[element]; ok {
		return false
	}
	elements := make(set)
	elements[element] = member
	ds.partitions[element] = elements
	return true
}

func (ds *DSet) FindPartition(element interface{}) (set, bool) {
	val, ok := ds.partitions[element]
	return val, ok
}

func (ds *DSet) AreDisjoint(element1 interface{}, element2 interface{}) (bool, bool) {
	p1, ok1 := ds.FindPartition(element1)
	p2, ok2 := ds.FindPartition(element2)
	if !ok1 || !ok2 {
		return false, false
	}
	return reflect.DeepEqual(p1, p2), true
}
