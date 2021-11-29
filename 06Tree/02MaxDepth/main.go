package main

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max = 0

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max++
	leftDepth:=maxDepth(root.Left)
	rightDepth:=maxDepth(root.Right)

	return int(math.Max(float64(leftDepth),float64(rightDepth)))
}
func main() {

}
