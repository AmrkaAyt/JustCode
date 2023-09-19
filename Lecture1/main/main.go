package main

import (
	"Lecture1/Tasks"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	target := 8
	fmt.Println(Tasks.TwoSum(nums, target))
}
