package main

/*
LC#83 https://leetcode.com/problems/remove-duplicates-from-sorted-list
移除有序链表中重复地元素
思路：在循环中，比较当前指针所指结点的值与其后结点的值，相同则后指针改指向更后的结点，不同则当前指针后移，
直至后指针为nil，结束循环
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head
	for {
		if current.Next == nil {
			break
		}
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
