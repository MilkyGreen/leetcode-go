package main

// 时间击败88% 空间击败95%
func removeElements(head *ListNode, val int) *ListNode {

	var pre *ListNode = nil
	node := head
	var ans *ListNode = nil
	for node != nil{
		if node.Val == val{
			if pre != nil{ // pre有可能还未空
				pre.Next = node.Next // 遇到val，跳过去
			}
		}else {
			if ans == nil{ // 第一个不为val的node是头
				ans = node
			}
			pre = node
		}
		node = node.Next
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}
