package main

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{Val: 2}}})
}
