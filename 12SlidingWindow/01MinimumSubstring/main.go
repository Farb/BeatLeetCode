package main

import (
	"fmt"
	"math"
	"strings"
)

func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	if ls < lt {
		return ""
	}
	tMap := make(map[byte]int, lt)
	for i := 0; i < lt; i++ {
		tMap[t[i]]++
	}

	min, minStr := math.MaxInt16, ""
	for i, j := 0, lt-1; i < ls && j < ls; {
		if !isValid(s[i:j+1], tMap) {
			j++
			continue
		}
		if j+1-i < min {
			min = j + 1 - i
			minStr = s[i : j+1]
		}
		i++
	}
	return minStr
}

func isValid(s string, tMap map[byte]int) bool {
	sMap := make(map[string]int, len(s))
	for i := 0; i < len(s); i++ {
		sMap[s[i:i+1]] = sMap[s[i:i+1]] + 1
	}
	for key := range tMap {
		if sMap[key] < tMap[key] {
			return false
		}
	}
	return true
}

// isValid1 是一种简单判别s是否包含t中的每个字符，但是当数据量大时会超时
func isValid1(s, t string) bool {
	for _, v := range t {
		index := strings.IndexRune(s, v)
		if index == -1 {
			return false
		}
		s = s[:index] + s[index+1:] // 之所以修改字符串是因为如果t中有重复字符，必须将匹配到的字符先去除，防止影响后续判断
	}
	return true
}

// https://leetcode-cn.com/problems/minimum-window-substring/solution/
func main() {
	s, t := "ADOBECODEBANC", "ABC"
	fmt.Println(minWindow(s, t)) // BANC

	s, t = "a", "aa"
	fmt.Println(minWindow(s, t)) // ""

	s, t = "a", "a"
	fmt.Println(minWindow(s, t)) // a

	s, t = "aa", "aa"
	fmt.Println(minWindow(s, t)) // aa
}
