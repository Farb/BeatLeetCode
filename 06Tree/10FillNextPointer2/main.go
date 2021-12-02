package main

// https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xes3em/

 // Definition for a Node.
 type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
 }


 //层序遍历
 func connect(root *Node) *Node {
	if root==nil {
		return root
	}
	queue:=[]*Node{root}
	for len(queue)>0{
		size:=len(queue) //队列中的元素个数，也就是需要出队列的次数
		var pre *Node
		for i := 0; i < size; i++ {
			node:=queue[0]
			queue=queue[1:]
			if pre!=nil {
				pre.Next=node //如果pre不为nil，就让pre.Next指向当前节点;如果为nil，则将当前第一个节点给pre
			}
			pre=node //pre向后移动一个节点，即指向当前节点
			if node.Left!=nil {
				queue = append(queue, node.Left)
			}
			if node.Right!=nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}