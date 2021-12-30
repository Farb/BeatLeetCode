package main

//https://leetcode-cn.com/problems/combinations/

func combine(n int, k int) [][]int {
	res := [][]int{}
	if n < k { //特殊情况:当要选取的个数大于数组的长度时，返回空数组
		return res
	}
	var dfs func(n, start, k int, path []int)
	dfs = func(n, start, k int, path []int) {
		if len(path) == k { //dfs结束条件:当路径的长度等于k时，满足组合条件，结束
			// finishedPath := make([]int, len(path))
			// copy(finishedPath, path) //因为切片时引用类型，将路径切片拷贝到新切片，并存入答案数组
			res = append(res, append([]int{},path...))
			return
		}
		for i := start; i <= n-(k-len(path)+1); i++ { //对n个数进行遍历，
			path = append(path, i) //将遍历的每个数存入路径数组
			dfs(n, i+1, k, path)
			path = path[:len(path)-1] //重置状态：将路径的最后一个元素弹出，回到树的当前层，继续遍历下个元素
		}
	}
	path := make([]int, 0, k) //定义一个路径切片，长度为要组合的个数
	dfs(n, 1, k, path) //对n个数进行dfs,起始数字为1
	return res
}
