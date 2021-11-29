package main

// Definition for a binary tree node
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

type ReturnType struct{
	
}

func isBalanced(root *TreeNode) bool {
	return process(root,0)
}
func process(root *TreeNode,height int) bool  {
	if root==nil {
		return true
	}
	leftIsBST:=process(root.Left,height)
	height++
	if !leftIsBST {
		return false
	}
	return process(root.Right,height)
}
// https://leetcode-cn.com/problems/balanced-binary-tree/
func main() {
	
}