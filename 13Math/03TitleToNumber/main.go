package main

import "math"

//https://leetcode-cn.com/problems/excel-sheet-column-number/
func titleToNumber(columnTitle string) (res int) {
	for i, j := len(columnTitle)-1, 0; i >= 0; i, j = i-1, j+1 {
		num := columnTitle[i] - 'A' + 1 //代表的10进制数
		res += int(num) * mypow(byte(26), byte(j))
	}
	return
}

func mypow(a, b byte) int {
	return int(math.Pow(float64(a), float64(b)))
}
