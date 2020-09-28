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
			{1, true},
			{1, false},
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
		testname := fmt.Sprintf("%#v", tt.value)
		t.Run(testname, func(t *testing.T) {
			_, ok := ds.FindPartition(tt.value)
			if ok != tt.expected {
				t.Errorf("Expected response %t but got %t for element %d", tt.expected, ok, tt.value)
			}

		})
	}
}

func TestAreDisjoint(t *testing.T) {
	initial := []interface{}{1, 2, 3}
	ds, _ := NewDSet(initial)
	var tests = []struct {
		value1, value2          interface{}
		expected, expectedError bool
	}{
		{1, 1, false, true},
		{2, 2, false, true},
		{3, 3, false, true},
		{1, 3, true, true},
		{4, 1, false, false},
		{5, 2, false, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%#v", tt)
		t.Run(testname, func(t *testing.T) {
			result, ok := ds.AreDisjoint(tt.value1, tt.value2)
			if result != tt.expected {
				t.Errorf("Expected response %t but got %t for element (%d,%d)",
					tt.expected, result, tt.value1, tt.value2)
			}
			if ok != tt.expectedError {
				t.Errorf("Expected error %t but got %t for element (%d,%d)",
					tt.expectedError, ok, tt.value1, tt.value2)

			}

		})
	}
}

func TestMerge(t *testing.T) {
	initial := []interface{}{1, 2, 3}
	var tests = []struct {
		value1, value2 interface{}
		expected       bool
	}{
		{1, 1, false},
		{5, 2, false},
		{1, 2, true},
		{1, 3, true},
	}
	for _, tt := range tests {
		ds, _ := NewDSet(initial)
		testname := fmt.Sprintf("%#v", ds)
		t.Run(testname, func(t *testing.T) {
			val, _ := ds.AreDisjoint(tt.value1, tt.value2)
			if tt.expected && !val {
				t.Errorf("Values should be disjoint initially")
			}
			result := ds.Merge(tt.value1, tt.value2)
			if result != tt.expected {
				t.Errorf("Expected response %t but got %t for element (%d,%d)",
					tt.expected, result, tt.value1, tt.value2)
			}
			val, _ = ds.AreDisjoint(tt.value1, tt.value2)
			if tt.expected && val {
				t.Errorf("Values should be not be disjoint after merge: %#v", ds)
			}

		})
	}
}
