package main

type TreeNode struct{
	Val int 
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) (values []int) {
	var postOrder func(root *TreeNode)
	postOrder=func(root *TreeNode) {
		if root!=nil {
			postOrder(root.Left)
			postOrder(root.Right)
			values = append(values, root.Val)
		}
	}
	postOrder(root)
	return
}

func postOrder_interation(root *TreeNode)(values []int)  {
	

	return
}