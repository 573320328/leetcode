package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
	third := head.Next.Next
	cur.Next = head
	head.Next = swapPairs(third)
	return cur
}

func swapPairs2(head *ListNode) *ListNode {
	dump := new(ListNode)
	dump.Next = head
	pre := dump
	nod := dump.Next
	for nod != nil {
		pre.Next = nod
		nod.Next = nod.Next.Next
		nod.Next.Next = nod
		pre = nod
		nod = nod.Next
	}
	return dump.Next
}
