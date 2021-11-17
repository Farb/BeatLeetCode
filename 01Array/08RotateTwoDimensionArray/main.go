package main

// need extra array
func rotate1(matrix [][]int) {
	res := make([][]int, 0,len(matrix))
	for i := 0; i < len(matrix); i++ {
		inner := make([]int,0, len(matrix))
		for j := len(matrix) - 1; j >= 0; j-- {
			inner = append(inner, matrix[j][i])
		}
		res = append(res, inner)
	}
	copy(matrix, res)
}

// need no extra array
func rotate(matrix [][]int) {
	// 1.水平翻转
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i],matrix[len(matrix)-1-i]=matrix[len(matrix)-1-i],matrix[i]
	}
	// 2.对角线交换
	for i := 0; i < len(matrix); i++ {
		for j := i+1; j < len(matrix); j++ { // 内循环从i+1开始（否则又交换2次等于没有交换），i==j的情况忽略
			matrix[i][j],matrix[j][i]=matrix[j][i],matrix[i][j]
		}
	}
}
// https://leetcode-cn.com/problems/rotate-image/
func main() {

}
