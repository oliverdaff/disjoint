package disjoint

import "fmt"

type DSet struct {
	partitions map[interface{}]interface{}
}

func NewDSet(initialSet []interface{}) (*DSet, error) {
	dset := DSet{make(map[interface{}]interface{})}
	for _, element := range initialSet {
		if _, ok := dset.partitions[element]; ok {
			return nil, fmt.Errorf("Element exists in partitions")
		} else {
			elements := map[interface{}]struct{}{}
			elements[element] = struct{}{}
			dset.partitions[element] = elements
		}
	}
	return &dset, nil
}
