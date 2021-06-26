package main

/**
LC#19 https://leetcode.com/problems/remove-nth-node-from-end-of-list
移除链表中倒数第N个结点
思路：双指针，其中一个先移N步，然后两个一起移，直到第一个移到尾部，将第二个指针所指结点的的next指向后一个结点
**/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	tail := head
	for ; n > 0; n-- {
		tail = tail.Next
	}

	if tail == nil {
		return head.Next
	}

	prev := &ListNode{0, head}
	for tail != nil {
		prev = prev.Next
		tail = tail.Next
	}
	prev.Next = prev.Next.Next
	return head
}
