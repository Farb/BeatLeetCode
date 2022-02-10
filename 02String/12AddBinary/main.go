package main

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/add-binary/
func addBinary(a string, b string) string {
	na, nb := len(a), len(b)
	res := []string{}
	ca := 0
	for i, j := na-1, nb-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := ca
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		res = append(res, fmt.Sprintf("%d", sum%2))
		ca = sum / 2
	}
	if ca == 1 {
		res = append(res, "1")
	}
	last := len(res) - 1
	for i := 0; i < len(res)/2; i++ {
		res[i], res[last-i] = res[last-i], res[i]
	}
	return strings.Join(res, "")
}
