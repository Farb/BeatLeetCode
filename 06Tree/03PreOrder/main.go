package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal_recursion(root *TreeNode) []int {
	var values = make([]int, 0)
	return preOrder(root, values)
}

func preOrder(root *TreeNode, values []int) []int {
	if root == nil {
		return values
	}
	values = append(values, root.Val)
	values = preOrder(root.Left, values)
	values = preOrder(root.Right, values)
	return values
}

// 递归法+内置函数
func preorderTraversal_recursion2(root *TreeNode) (values []int) {
	var pre func(root *TreeNode)
	pre=func (root *TreeNode)  {
		if root!=nil {
			values = append(values, root.Val)
			pre(root.Left)
			pre(root.Right)
		}
	}
	pre(root)
	return 
}

// 迭代法
func preorderTraversal(root *TreeNode) (values []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			values = append(values, node.Val)
			stack = append(stack, node)
			node = node.Left // 遍历左子树
		}
		node, stack = stack[len(stack)-1], stack[:len(stack)-1] //  取栈顶元素
		node = node.Right                                       // 遍历右子树
	}
	return values
}

//迭代法2 ，这个版本更好理解
func preOrder_Iteration2(root *TreeNode) (values []int) {
	if root==nil {
		return
	}
	stack:=[]*TreeNode{}
	stack = append(stack, root) // 将头节点入栈
	node:=&TreeNode{}
	for len(stack)>0{
		node,stack=stack[len(stack)-1],stack[:len(stack)-1] // 出栈
		values = append(values, node.Val)

		if node.Right!=nil { // 右孩子入栈 ，因为是先序遍历，所以打印顺序是中左右，故入栈是中右左
			stack = append(stack, node.Right)
		}

		if node.Left!=nil { // 左孩子入栈
			stack = append(stack, node.Left)
		}
	}
	return
}

// https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xeywh5/
func main() {

}
