package main

// 调整大顶堆，n为数组长度，i为要调整的元素下标（父节点下标）
func heapify(tree []int, n, i int) {
	if i >= n { //basecase:索引越界退出
		return
	}
	max := i                            //将要调整的元素下标看作是最大元素的下标
	c1, c2 := 2*i+1, 2*i+2              //左右孩子的下标分别是2i+1和2i+2
	if c1 < n && tree[max] < tree[c1] { //左孩子未越界，且大于父节点值，则左孩子作为最大值
		max = c1
	}
	if c2 < n && tree[max] < tree[c2] { //右孩子未越界，且大于父节点值，则右孩子作为最大值
		max = c2
	}
	if max != i { //如果父节点下标不等于最大值的下标（最大值肯定是左右孩子），那么交换父节点的值和最大值
		tree[max], tree[i] = tree[i], tree[max]
		heapify(tree, n, max) //递归，从次最大值下标开始调整堆
	}
	//否则，说明已经是大顶堆了，不需要调整
}

// 构建大顶堆树
func buildTree(tree []int, n int) {
	lastNode := (len(tree) - 1) / 2  //先找到最后一个要调整的节点，下标为(n-1)/2
	for i := lastNode; i >= 0; i-- { //然后依次递减下标至0，调整大顶堆，每次从i开始调整
		heapify(tree, n, i)
	}
}

// 堆排序
func heapSort(tree []int) {
	n := len(tree)
	buildTree(tree, n)            //1.先构建大顶堆树
	for i := n - 1; i >= 0; i-- { //i为每次遍历的节点个数
		tree[i], tree[0] = tree[0], tree[i]
		heapify(tree, i, 0) //2.每次根节点和末尾元素交换，之后再砍掉末尾元素最大值，从第一个元素继续调整为大顶堆
	}
}
