package Tasks

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		hasAns bool
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 3, hasAns: true},
		{nums: []int{2, 7, 11, 15}, target: 9, hasAns: true},
		{nums: []int{3, 2, 4}, target: 6, hasAns: true},
		{nums: []int{3, 3}, target: 6, hasAns: true},
		{nums: []int{9, 9, 9}, target: 2, hasAns: false},
		{nums: []int{}, target: 5, hasAns: false},
		{nums: []int{3, 5, 99, 4132, 432, 3245, 43, 23, 2}, target: 5, hasAns: true},
	}

	for _, tc := range tests {
		got := TwoSum(tc.nums, tc.target)

		switch {
		case tc.hasAns && len(got) == 2:
			if tc.target != (tc.nums[got[0]] + tc.nums[got[1]]) {
				t.Fatalf("%v answer is incorrect", got)
			}
		case !tc.hasAns && len(got) == 0:
		default:
			t.Fatalf("%v answer is incorrect", got)
		}

	}
}
