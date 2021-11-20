package main

import "fmt"

func reverse(x int) int {
	res := 0
	for x != 0 {
		n := x % 10
		if res > 214748364 || (res == 214748364 && n > 7) {
			return 0
		}
		if res < (-214748364) || (res == -214748364 && n < (-8)) {
			return 0
		}
		x = x / 10
		res = res*10 + n
	}
	return res
}
// https://leetcode-cn.com/problems/reverse-integer/
func main() {
	fmt.Println(reverse(0)) // 0
	fmt.Println(reverse(12)) // 21
	fmt.Println(reverse(-12)) // -21
	fmt.Println(reverse(120)) // 21
	fmt.Println(reverse(-120)) // -21
	fmt.Println(reverse(-1234567899)) // 0
	fmt.Println(reverse(1234567899)) // 0
}
