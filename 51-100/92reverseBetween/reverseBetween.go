package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	res := reverseBetween(head, 2, 4)
	PrintListNode(res)
}

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
说明:
1 ≤ m ≤ n ≤ 链表长度。
示例:
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

var successor *ListNode // 一定要提出来做一个全局变量，不然在函数内部定义，递归时会改变值

func reverseN(head *ListNode, m int) *ListNode {

	if m == 1 {
		successor = head.Next
		return head
	}

	last := reverseN(head.Next, m-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
