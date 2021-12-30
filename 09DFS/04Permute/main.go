package main

//https://leetcode-cn.com/problems/permutations/
// 思路：DFS + 回溯算法

func permute(nums []int) [][]int {
	var res = [][]int{}
	numsLen := len(nums)
	if numsLen == 0 {
		return res
	}
	var dfs func(nums []int, numsLen, depth int, path []int, used []bool)
	dfs = func(nums []int, numsLen, depth int, path []int, used []bool) {
		if depth == numsLen {
			finishedPath := make([]int, numsLen)
			copy(finishedPath, path)
			res = append(res, finishedPath)
			return
		}
		for i := 0; i < numsLen; i++ {
			if !used[i] {
				used[i] = true                          //修改当前的节点索引为已使用状态
				path = append(path, nums[i])            //将当前节点加入路径
				dfs(nums, numsLen, depth+1, path, used) //继续下一层搜索
				used[i] = false                         //深一层搜索完毕后，恢复先前的状态
				path = path[:len(path)-1]               //路径也要撤销当前节点
			}
		}
	}
	path, used := make([]int, 0, numsLen), make([]bool, numsLen)
	dfs(nums, numsLen, 0, path, used)
	return res
}
