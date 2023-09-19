package main

import (
	"Lecture2/Tasks"
	"fmt"
)

func main() {
	list1 := &Tasks.ListNode{Val: 1, Next: &Tasks.ListNode{Val: 2, Next: &Tasks.ListNode{Val: 4}}}
	list2 := &Tasks.ListNode{Val: 1, Next: &Tasks.ListNode{Val: 3, Next: &Tasks.ListNode{Val: 5}}}

	mergedList := Tasks.MergeTwoLists(list1, list2)

	fmt.Println("Merged List:")
	Tasks.PrintList(mergedList)

}
