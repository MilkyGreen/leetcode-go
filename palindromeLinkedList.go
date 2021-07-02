package main


func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head

	for fast.Next != nil &&  fast.Next.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	half := slow
	if slow.Next != nil {
		half = reverse(slow.Next) // 下半个链表的开始,翻转一下顺序
	}
	for half != nil{
		if head.Val != half.Val{
			return false
		}
		head = head.Next
		half = half.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head.Next == nil{
		return head
	}
	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
