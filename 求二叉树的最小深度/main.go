package main

import "math"

func main(){

}
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root==nil {
		return 0

	}//空树返回0
	if root.Right==nil&&root.Left==nil {
		return 1

	}//左右节点不存在，返回1
	if root.Left==nil{
		return minDepth(root.Right)+1

	}//有右节点，没有左边的节点，继续递归
	if root.Right==nil {
		return minDepth(root.Left)+1

	}//有左节点，没有右节点，继续递归
	return int(math.Min(float64(minDepth(root.Left)),float64(minDepth(root.Right))))+1

}