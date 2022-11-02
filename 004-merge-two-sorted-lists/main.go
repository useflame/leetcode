package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current := &ListNode{}
	result := current

	for list1 != nil && list2 != nil {
		fmt.Println(list1, list2)
		if list1.Val < list2.Val {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
		} else {
			current.Next = list2
			current = current.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return result.Next
}
