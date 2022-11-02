package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testTable := []struct {
		firstList  *ListNode
		secondList *ListNode
		output     *ListNode
	}{
		{
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
	}

	for _, item := range testTable {
		fmt.Println(1)
		output := MergeTwoLists(item.firstList, item.secondList)
		fmt.Println(2)
		if !isListEqual(listToSlice(output), listToSlice(item.output)) {
			t.Errorf("Got: %v, want: %v.", output, item.output)
		}
	}
}

func listToSlice(list *ListNode) []int {
	result := []int{}

	for list.Next != nil {
		result = append(result, list.Val)
		list = list.Next
	}

	return result
}

func isListEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func (l *ListNode) String() string {
	x := make([]int, 0)
	curr := l

	for curr != nil {
		x = append(x, curr.Val)
		curr = curr.Next
	}
	return fmt.Sprint(x)
}

// func convertToList(arr []int) *ListNode {
// 	var head *ListNode
// 	for _, val := range arr {
// 		head = insert(head, val)
// 	}
// 	return head
// }

// func insert(root *ListNode, elem int) *ListNode {
// 	temp := ListNode{Val: elem, Next: nil}
// 	if root == nil {
// 		root = &temp
// 		return root
// 	}
// 	curr := root
// 	for curr.Next != nil {
// 		curr = curr.Next
// 	}
// 	curr.Next = &temp
// 	return root
// }
