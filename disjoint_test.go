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
