package main

import "testing"

func TestRemoveNthFromEnd1(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = removeNthFromEnd(head, 2)
	_printListNodes(head)
}

func TestRemoveNthFromEnd2(t *testing.T) {
	head := &ListNode{1, nil}
	head = removeNthFromEnd(head, 1)
	_printListNodes(head)
}

func TestRemoveNthFromEnd3(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}
	head = removeNthFromEnd(head, 1)
	_printListNodes(head)
}

func TestRemoveNthFromEnd4(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}
	head = removeNthFromEnd(head, 2)
	_printListNodes(head)
}
