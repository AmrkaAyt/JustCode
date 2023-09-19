package Task1

import "sort"

func CompareTwoSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	sort.Ints(s1)
	sort.Ints(s2)

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
