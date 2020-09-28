package disjoint

import (
	"fmt"
	"testing"
)

func TestNewDSet(t *testing.T) {
	var tests = []struct {
		elements    []interface{}
		errExpected bool
	}{
		{[]interface{}{1, 2, 3, 4}, false},
		{[]interface{}{1, 1}, true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.elements)
		t.Run(testname, func(t *testing.T) {
			_, err := NewDSet(tt.elements)
			if err == nil && tt.errExpected {
				t.Errorf("Error expected")
			} else if err != nil && !tt.errExpected {
				t.Errorf("Error expected")
			}
		})
	}
}

type AddExpected struct {
	value    interface{}
	expected bool
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		elements []AddExpected
	}{
		{[]AddExpected{
			AddExpected{1, true},
			AddExpected{1, false},
		},
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%#v", tt.elements)
		t.Run(testname, func(t *testing.T) {
			val, _ := NewDSet(make([]interface{}, 0))
			for _, element := range tt.elements {
				result := val.Add(element.value)
				if result != element.expected {
					t.Errorf("Expected response %t but got %t for element %d", element.expected, result, element.value)
				}

			}
		})
	}
}

func TestFindPartition(t *testing.T) {
	initial := []interface{}{1, 2, 3}
	ds, _ := NewDSet(initial)
	var tests = []struct {
		value    interface{}
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%#v", tt.expected)
		t.Run(testname, func(t *testing.T) {
			_, ok := ds.FindPartition(tt.value)
			if ok != tt.expected {
				t.Errorf("Expected response %t but got %t for element %d", tt.expected, ok, tt.value)
			}

		})
	}
}
