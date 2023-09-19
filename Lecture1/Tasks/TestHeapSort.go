package Tasks

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 9, 3, 6, 2, 8, 1, 7}, []int{1, 2, 3, 5, 6, 7, 8, 9}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{5, 5, 5, 5}, []int{5, 5, 5, 5}},
		{[]int{9, 7, 5, 3, 1}, []int{1, 3, 5, 7, 9}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 1, 3, 5, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 3, 5, 7, 9, 2, 4, 6}, []int{1, 2, 3, 4, 5, 6, 7, 9}},
		{[]int{7, 3, 2, 8, 1, 6, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, test := range tests {
		arr := make([]int, len(test.input))
		copy(arr, test.input)
		HeapSort(arr)
		if !reflect.DeepEqual(arr, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, arr)
		}
	}
}
