package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 笨办法：层级遍历
func maxDepth_levelOrder(root *TreeNode) (max int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		max++
		nextQueue := []*TreeNode{}
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)
			}
		}
		queue = nextQueue
	}
	return
}

// 自顶向下递归
func maxDepth_topBottom(root *TreeNode) int {
	max := 0
	var getDepth func(root *TreeNode, depth int)
	getDepth = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			max = getMax(max, depth)
		}
		getDepth(root.Left, depth+1)
		getDepth(root.Right, depth+1)
	}
	getDepth(root, 1)
	return max
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 自底向上
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func main() {

}
