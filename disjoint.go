package disjoint

import "fmt"

type void struct{}

var member void

type Set map[interface{}]void

type DSet struct {
	partitions map[interface{}]Set
}

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

func (ds *DSet) Add(element interface{}) bool {
	if _, ok := ds.partitions[element]; ok {
		return false
	}
	elements := make(Set)
	elements[element] = member
	ds.partitions[element] = elements
	return true
}
