package main

// ListNode ==> Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		# Python Implementation
		def add_two_numbers(self, l1: ListNode, l2: ListNode) -> ListNode:
		   dummy = cur = ListNode()
		   carry = 0
		   while l1 or l2 or carry:
			   if l1:
				   carry += l1.val
				   l1 = l1.next
			   if l2:
				   carry += l2.val
				   l2 = l2.next
			   cur.next = ListNode(carry % 10)
			   cur = cur.next
			   carry //= 10
		   return dummy.next
	*/
	carry := 0
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: carry % 10}
		cur = cur.Next
		carry /= 10
	}
	return dummy.Next
}
