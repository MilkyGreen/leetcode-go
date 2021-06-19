package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归调用翻转函数：先把head拿出来，翻转后面的链表，然后把head放在最后
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil{
		return head
	}
	newHead := reverseList(head.Next) // 翻转head.Next的部分
	head.Next.Next = head // 原来的head.Next变成了新的tail，把head放在tail后面变成新tail
	head.Next = nil // head现在是tail，后面没有元素了
	return newHead
}
