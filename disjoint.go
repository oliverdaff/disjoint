// Package disjoint contains a UnionFind (Disjoint-set) data structure.
package disjoint

import (
	"fmt"
	"reflect"
)

type void struct{}

var member void

//Set is used to store members of a set.
type Set map[interface{}]void

// DSet is a Disjoint Set data structure that keeps track of a set of elements
// partitioned into a number of disjoint (nonoverlapping) subsets.
//
// DSet supports two operations
//
// - Find: Returns which subset a particular element is in. Find returns an item that can be
// compared with other return values to determine it the two elements are in the same subset.
// - Union: Join two subsets into a single subset.
type DSet struct {
	partitions map[interface{}]Set
}

// NewDSet creates a new DSet with the elements
// provide each part of their own set.  A error is returned if the input
// elements contain duplicate elements.
func NewDSet(initialSet []interface{}) (*DSet, error) {
	dset := DSet{
		make(map[interface{}]Set),
	}
	for _, element := range initialSet {
		if _, ok := dset.partitions[element]; ok {
			return nil, fmt.Errorf("Element exists in partitions")
		}
		elements := make(Set)
		elements[element] = member
		dset.partitions[element] = elements
	}
	return &dset, nil
}

// Add adds a new element to its own partition
func (ds *DSet) Add(element interface{}) bool {
	if _, ok := ds.partitions[element]; ok {
		return false
	}
	elements := make(Set)
	elements[element] = member
	ds.partitions[element] = elements
	return true
}

// FindPartition finds the partition the given element belongs to returning
// false if the element is not in the partitions.
func (ds *DSet) FindPartition(element interface{}) (Set, bool) {
	val, ok := ds.partitions[element]
	return val, ok
}

//AreDisjoint returns true if element1 and element2 belong to the same partition
//otherwise it returns false.  If either element does not exist
//it returns (false, false)
func (ds *DSet) AreDisjoint(element1 interface{}, element2 interface{}) (bool, bool) {
	p1, ok1 := ds.FindPartition(element1)
	p2, ok2 := ds.FindPartition(element2)
	if !ok1 || !ok2 {
		return false, false
	}
	return reflect.DeepEqual(p1, p2), true
}

//Merge merges the subsets contianing the two elements passed.
//True is returned if the two subsets are merged, false
//is returned if either of the two elements do not exist
//or if element1 and element2 all ready belong to the same partition.
func (ds *DSet) Merge(element1 interface{}, element2 interface{}) bool {
	p1, ok1 := ds.FindPartition(element1)
	p2, ok2 := ds.FindPartition(element2)
	if !ok1 || !ok2 {
		return false
	}
	if reflect.DeepEqual(p1, p2) {
		return false
	}
	for k := range p1 {
		p2[k] = member
		ds.partitions[k] = p2
	}
	return true
}
