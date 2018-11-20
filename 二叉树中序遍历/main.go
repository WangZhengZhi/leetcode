package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tmp []int //保存数据

func inorderTraversal(root *TreeNode) []int {
	tmp = []int{}
	Help(root)
	return tmp
}
func Help(root *TreeNode) { //中序递归遍历
	if root == nil {
		return
	}
	Help(root.Left)
	tmp = append(tmp, root.Val)
	Help(root.Right)
}
