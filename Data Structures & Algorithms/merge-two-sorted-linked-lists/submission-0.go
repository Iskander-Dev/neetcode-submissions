/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
	prev := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			prev = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			prev = list2
			list2 = list2.Next
		}
	}

	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}

	return dummy.Next
}
