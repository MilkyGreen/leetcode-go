package main

// bfs遍历，取每一层的最后一个元素放入结果    击败100% go提交
func rightSideView(root *TreeNode) []int {

	answer := make([]int,0)
	if root == nil {
		return answer
	}
	l := 1 // 一下层要遍历的元素数量
	list := make([]*TreeNode,1) // 待遍历元素列表
	list = append(list, root)
	answer = append(answer,root.Val)
	for l > 0 {
		nl:=0
		newList := make([]*TreeNode,0)
		for _,node := range list{
			if node != nil{  // 按理说不用加这一行，但是会报空指针
				if node.Left != nil {
					newList = append(newList,node.Left)
					nl++
				}
				if node.Right != nil {
					newList = append(newList,node.Right)
					nl++
				}
			}
		}
		if nl > 0{
			answer = append(answer,newList[nl-1].Val)
		}
		
		list = newList
		l = nl
	}
	return answer
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
