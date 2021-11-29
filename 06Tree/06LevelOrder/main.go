package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (values [][]int) {
	if root == nil {
		return
	}
	curNode := root
	queue := []*TreeNode{curNode}

	for len(queue) > 0 {
		val:=make([]int,0)
		nextLevelQueue:=[]*TreeNode{}
		for i := 0; i < len(queue); i++ {
			curNode = queue[i]
			val = append(val, curNode.Val)
			if curNode.Left != nil {
				nextLevelQueue = append(nextLevelQueue, curNode.Left)
			}
			if curNode.Right != nil {
				nextLevelQueue = append(nextLevelQueue, curNode.Right)
			}
		}
		queue=nextLevelQueue
		values = append(values, val)
	}
	return
}
