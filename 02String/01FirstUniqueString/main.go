package main

import (
	"fmt"
	"strings"
)

// v1 使用strings中的api
func firstUniqCharV1(s string) int {
	for i := range s {
		str := string(s[i])
		if strings.Index(s, str) == strings.LastIndex(s, str) {
			return strings.Index(s, str)
		}
	}
	return -1
}

// v2使用数组存储字母频率
func firstUniqChar(s string) int {
	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

//https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func main() {
	s := "leetcode" // 返回 0
	fmt.Println(firstUniqChar(s))

	s = "loveleetcode" // 返回 2
	fmt.Println(firstUniqChar(s))

}
