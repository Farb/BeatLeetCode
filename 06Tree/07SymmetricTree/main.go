package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isMirror func(left, right *TreeNode) bool
	isMirror = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}
	return isMirror(root.Left, root.Right)
}

// https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xoxzgv/
func main() {

}
