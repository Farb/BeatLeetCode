package main

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q { //基本情况：root为空或p，q其中一个为root,那么最低公共祖先就是root，很好想的
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)   //在左子树上找到了p或q的祖先
	right := lowestCommonAncestor(root.Right, p, q) //在右子树上找到了p或q的祖先
	if left == nil {
		return right //说明在右子树上找到了祖先
	}
	if right == nil {
		return left //说明在左子树上找到了祖先
	}
	return root //说明在左右子树上都找到了祖先，那么root必然是最低公共祖先
}

// 迭代法
func lowestCommonAncestor_iteration(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	parentMap := map[*TreeNode]*TreeNode{root: nil}
	queue := []*TreeNode{root}
	for {
		_, existP := parentMap[p]
		_, existQ := parentMap[q]
		if existP && existQ { // 如果已经找到了p,q结点则跳出循环
			break
		}
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if node.Left != nil {
			queue = append(queue, node.Left)
			parentMap[node.Left] = node //将当前结点的左节点作为key,当前节点作为Value,
		}
		if node.Right != nil { //同理
			queue = append(queue, node.Right)
			parentMap[node.Right] = node
		}
	}

	ancestors := map[*TreeNode]*TreeNode{} //从p开始，找p节点的祖先,一直找到root结束，root的父节点为nil
	for p != nil {
		ancestors[p] = p
		p = parentMap[p]
	}

	//从q开始，找q节点的祖先，直到p的祖先中第一个包含的祖先节点为止
	for i := 0; i < len(ancestors); i++ {
		if commonAncestor, ok := ancestors[q]; ok {
			return commonAncestor
		}
		q = parentMap[q]
	}
	return nil
}
