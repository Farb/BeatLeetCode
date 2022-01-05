package main

//https://leetcode-cn.com/problems/2bCMpM/
//最近距离使用BFS，最远距离使用DFS
//BFS思路：先将矩阵中的0的坐标放入队列，然后将矩阵中的1改为-1，表示未访问过
func updateMatrix(mat [][]int) [][]int {
	var queue [][]int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}
	process(mat, queue)
	return mat
}
func process(mat, queue [][]int) {
	//弹出一个为0的格子，对其四周进行遍历
	var point []int
	dx, dy := []int{0, 0, -1, 1}, []int{1, -1, 0, 0} //上下左右四个格子的相对坐标
	for len(queue) > 0 {
		point, queue = queue[0], queue[1:]
		for i := 0; i < len(dx); i++ {
			x, y := point[0], point[1]
			newX, newY := x+dx[i], y+dy[i]
			if inMatrix(mat, newX, newY) && mat[newX][newY] == -1 { //如果四周的格子坐标在矩阵范围内且未访问过
				mat[newX][newY] = mat[x][y] + 1          //将格子为原本为1的值修改为内层格子的值（从格子为0开始遍历的距离）+1
				queue = append(queue, []int{newX, newY}) //将原本为1的格子加入队列，进行下次遍历计算距离
			}
		}
	}
}

func inMatrix(mat [][]int, row, col int) bool {
	return 0 <= row && row < len(mat) && 0 <= col && col < len(mat[0])
}
