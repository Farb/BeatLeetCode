package main
//https://leetcode-cn.com/problems/max-area-of-island/

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, getIslandArea(grid, i, j))
			}
		}
	}
	return maxArea
}

func getIslandArea(grid [][]int, row, col int) int {
	if !inMetrix(grid, row, col) || grid[row][col] != 1 {
		return 0
	}

	grid[row][col] = 2 //已访问过

	return 1 +
		getIslandArea(grid, row-1, col) + //上
		getIslandArea(grid, row+1, col) + //下
		getIslandArea(grid, row, col-1) + //左
		getIslandArea(grid, row, col+1) //右
}

func inMetrix(grid [][]int, row, col int) bool {
	return 0 <= row && row < len(grid) && 0 <= col && col < len(grid[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
