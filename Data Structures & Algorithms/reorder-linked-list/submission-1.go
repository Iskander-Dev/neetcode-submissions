/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
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

	first, second := head, head1

	for second != nil {
		next1, next2 := first.Next, second.Next
		first.Next = second
		second.Next = next1
		first, second = next1, next2
	}

	// dummy := &ListNode{}
	// prev := dummy

	// for head1 != nil {
	// 	prev.Next = head
    //     prev = prev.Next
	// 	next := head.Next
    //     prev.Next = head1
    //     prev = prev.Next
    //     head = next
    //     head1 = head1.Next
	// }
	// prev.Next = head
	// return dummy.Next
}
