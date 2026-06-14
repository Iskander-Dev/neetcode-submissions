/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) *ListNode {
    slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	cur, head1 := slow.Next, (*ListNode)(nil)
	slow.Next = (*ListNode)(nil)

	for cur != nil {
		next := cur.Next
		cur.Next = head1
		head1 = cur
		cur = next
	}

	dummy := &ListNode{}
	prev := dummy
	// 2, 4
	// 6, 8

	for head1 != nil {
		prev.Next = head
        prev = prev.Next
		next := head.Next
        prev.Next = head1
        prev = prev.Next
        head = next
        head1 = head1.Next
	}
	prev.Next = head
	return dummy.Next
}
