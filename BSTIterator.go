package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	elements []int
	size     int
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.toArray(root)
	return it
}

func (this *BSTIterator) toArray(curr *TreeNode) {
	if curr != nil {
		this.toArray(curr.Left)
		this.elements = append(this.elements, curr.Val)
		this.size = this.size + 1
		this.toArray(curr.Right)
	}
}

func (this *BSTIterator) Next() int {
	res := this.elements[0]
	this.elements = this.elements[1:]
	this.size = this.size - 1
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.size > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
