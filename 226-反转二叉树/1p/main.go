package main

//二叉树结构体
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}

/**
* 官方解题
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

**/
