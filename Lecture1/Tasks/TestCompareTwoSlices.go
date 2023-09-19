package Tasks

import "testing"

func TestCompareTwoSlices(t *testing.T) {
	tests := []struct {
		s1       []int
		s2       []int
		expected bool
	}{
		{s1: []int{1, 2, 3}, s2: []int{1, 2, 3}, expected: true},
		{s1: []int{1, 2, 3}, s2: []int{3, 2, 1}, expected: true},
		{s1: []int{1, 2, 3}, s2: []int{1, 2, 4}, expected: false},
		{s1: []int{1, 2, 3, 4}, s2: []int{1, 2, 3}, expected: false},
		{s1: []int{}, s2: []int{}, expected: true},
		{s1: []int{1, 2, 3}, s2: []int{4, 5, 6}, expected: false},
		{s1: []int{1, 2, 3}, s2: []int{1, 2, 2}, expected: false},
		{s1: []int{1, 2, 3}, s2: []int{1, 2, 3, 4}, expected: false},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			result := CompareTwoSlices(tc.s1, tc.s2)
			if result != tc.expected {
				t.Fatalf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}
