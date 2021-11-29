package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归+内置函数法
func inorderTraversal(root *TreeNode) (values []int) {
	if root == nil {
		return
	}
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root != nil {
			inOrder(root.Left)
			values = append(values, root.Val)
			inOrder(root.Right)
		}
	}
	inOrder(root)
	return
}

// 迭代法
func inorderTraversal_interation(root *TreeNode) (values []int) {
	stack := []*TreeNode{}
	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}
		curNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
		values = append(values, curNode.Val)
		curNode = curNode.Right
	}
	return
}

type Pair struct{
	Node *TreeNode
	Color int
}
// 颜色标记法
func inorderTraversal_color(root *TreeNode) (values []int){
	White,Gray:=0,1
	stack:=[]*Pair{{root,White}}
	for len(stack)>0{
		pair:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		if pair.Node==nil {
			continue
		}
		if pair.Color==White {
			stack = append(stack, &Pair{pair.Node.Right,White})
			stack = append(stack, &Pair{pair.Node,Gray})
			stack = append(stack, &Pair{pair.Node.Left,White})
		}else{
			values = append(values, pair.Node.Val)
		}
	}
	return
}
// https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xecaj6/
func main() {

}
