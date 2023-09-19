package Task1

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
		hasAns   bool
	}{
		{strs: []string{"flower", "flow", "flight"}, expected: "fl", hasAns: true},
		{strs: []string{"dog", "racecar", "car"}, expected: "", hasAns: false},
		{strs: []string{"abc", "abcdef", "abcde"}, expected: "abc", hasAns: true},
		{strs: []string{"apple", "app", "ap"}, expected: "ap", hasAns: true},
		{strs: []string{"abcd", "abcd", "abcd"}, expected: "abcd", hasAns: true},
		{strs: []string{"", "xyz", "123"}, expected: "", hasAns: false},
		{strs: []string{"", "", ""}, expected: "", hasAns: false},
		{strs: []string{"abc"}, expected: "abc", hasAns: true},
		{strs: []string{}, expected: "", hasAns: false},
		{strs: []string{"abc", "def", "ghi"}, expected: "", hasAns: false},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			result := longestCommonPrefix(tc.strs)

			switch {
			case tc.hasAns && result == tc.expected:
			case !tc.hasAns && result == "":
			default:
				t.Fatalf("Expected '%s' but got '%s'", tc.expected, result)
			}
		})
	}
}
