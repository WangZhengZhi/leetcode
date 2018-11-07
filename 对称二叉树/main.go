package miain

//对称二叉树
type Element interface {
}
type TreeNode struct {
	val   Element
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return Ismirror(root, root)
}
func Ismirror(left, right *TreeNode) bool {

	if left == nil { //left为nil
		return right == nil
	} else if right == nil { //right为nil
		return left == nil
	} else if left.val == right.val { // left.val==right.val
		return Ismirror(left.Left, right.Right) && Ismirror(left.Right, right.Left) //递归

	} else {
		return false
	}

}
