package main

// https://leetcode-cn.com/problems/flood-fill/
var flag = -1 //将所有需要染色的格子标记为-1，下次遍历时再染色为新色即可
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	process(image, sr, sc, image[sr][sc])
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			if image[i][j] == flag {
				image[i][j] = newColor
			}
		}
	}
	return image
}

func process(image [][]int, row, col, targetColor int) {
	if !inMetrix(image, row, col) || image[row][col] != targetColor {
		return
	}
	image[row][col] = flag //染色标记为-1

	process(image, row-1, col, targetColor) //上
	process(image, row+1, col, targetColor) //下
	process(image, row, col-1, targetColor) //左
	process(image, row, col+1, targetColor) //右
}

func inMetrix(image [][]int, row, col int) bool {
	return 0 <= row && row < len(image) && 0 <= col && col < len(image[0])
}
