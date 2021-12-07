package main

func numIslands(grid [][]byte) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '1' {
				continue
			}
			searchLand(grid, i, j)
			num++
		}
	}
	return num
}

func searchLand(grid [][]byte, i, j int)  {

	if !isInArea(grid,i,j) { //过滤界外点
		return 
	}
	if grid[i][j] != '1' { //过滤海洋方格和已访问点
		return 
	}

	grid[i][j]='2' // 标记为感染点

	searchLand(grid,i-1,j) // 上
	searchLand(grid,i+1,j) // 下
	searchLand(grid,i,j-1) // 左
	searchLand(grid,i,j+1) // 右
}

func isInArea(grid [][]byte, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
}
