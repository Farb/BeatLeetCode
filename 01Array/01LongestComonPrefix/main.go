package main

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]              //设置第一个字符串位最长公共字符串
	for _, str := range strs[1:] { //遍历每个字符串
		for strings.Index(str, prefix) !=0 { // 如果当前字符串找不到该最长公共字符串，则不断缩减该最长公共字符串
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1] //不断刷新最长公共字符串，不断缩减
		}
	}
	return prefix
}
func main() {
	var strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"c", "acc", "ccc"}
	fmt.Println(longestCommonPrefix(strs))

}
