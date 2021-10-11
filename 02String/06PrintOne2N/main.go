package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/

func printNumbers(n int) []int {
	if n < 0 {
		return []int{}
	}
	max := int(math.Pow10(n))
	res := make([]int, 0, max-1)
	for i := 1; i < max; i++ {
		res = append(res, i)
	}
	fmt.Println(cap(res))
	return res
}

func main() {
	n := 0
	fmt.Println(printNumbers(n))

	n = 1
	fmt.Println(printNumbers(n))

	n = 2
	fmt.Println(printNumbers(n))
}
