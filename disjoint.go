// Package disjoint contains a UnionFind (Disjoint-set) data structure.
// It is proven that the  running amortized time for `m` calls to to `FindPartition`
// and `Merge` on a set of `n` elements will require `O(m * Ak(n))` array access.
// Ak(n) is an approximation of the inverse Ackermann function, growing
// so slowly that it can be considered a constant.
package disjoint

import (
	"fmt"
	"reflect"
)

//Item in element map
type item struct {
	root interface{}
	rank uint
}

// DSet is a Disjoint Set data structure that keeps track of a set of elements
// partitioned into a number of disjoint (nonoverlapping) subsets.
//
// DSet supports two operations.
//
// - Find: Returns which subset a particular element is in. Find returns an item that can be
// compared with other return values to determine it the two elements are in the same subset.
//
// - Union: Join two subsets into a single subset.
type DSet struct {
	partitions map[interface{}]*item
}

// NewDSet creates a new DSet with the elements
// provide each part of their own set.  A error is returned if the input
// elements contain duplicate elements.
func NewDSet(initialSet []interface{}) (*DSet, error) {
	dset := DSet{
		make(map[interface{}]*item),
	}
	for _, element := range initialSet {
		if _, ok := dset.partitions[element]; ok {
			return nil, fmt.Errorf("Element exists in partitions")
		}
		dset.partitions[element] = &item{element, 1}
	}
	return &dset, nil
}

// Size returns the number of elements in the set.
func (ds *DSet) Size() int {
	return len(ds.partitions)
}

// Add adds a new element to its own partition
func (ds *DSet) Add(element interface{}) bool {
	if _, ok := ds.partitions[element]; ok {
		return false
	}
	ds.partitions[element] = &item{element, 1}
	return true
}

// FindPartition finds the partition the given element belongs to returning
// false if the element is not in the partitions.
func (ds *DSet) FindPartition(element interface{}) (interface{}, bool) {
	val, ok := ds.partitions[element]
	if !ok {
		return nil, false
	}
	if val.root == element {
		return element, true
	}
	root, _ := ds.FindPartition(val.root)
	val.root = root
	return root, true
}

//Merge merges the subsets containing the two elements passed.
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

	info1 := ds.partitions[p1]
	info2 := ds.partitions[p2]

	if info1.rank >= info2.rank {
		info2.root = info1.root
		info1.rank += info2.rank
	} else {
		info1.root = info2.root
		info2.rank += info1.rank
	}
	return true
}

//AreDisjoint returns false if element1 and element2 belong to the same partition
//otherwise it returns true.  If either element does not exist
//it returns (false, false)
func (ds *DSet) AreDisjoint(element1 interface{}, element2 interface{}) (bool, bool) {
	p1, ok1 := ds.FindPartition(element1)
	p2, ok2 := ds.FindPartition(element2)
	if !ok1 || !ok2 {
		return false, false
	}
	return p1 != p2, true
}
