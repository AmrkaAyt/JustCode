package Task1

// easiest sol
// func twoSum(nums []int, target int) []int {
//	for i, r := range nums {
//		for j := i + 1; j < len(nums); j++ {
//			if r+nums[j] == target && i < j {
//				return []int{i, j}
//			}
//		}
//	}
//	return []int{}
//}

// best sol

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		j, r := m[target-n]
		if r {
			return []int{j, i}
		}
		m[n] = i
	}
	return []int{}
}
