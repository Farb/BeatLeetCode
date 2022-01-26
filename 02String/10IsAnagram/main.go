package main

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
func isAnagram(s string, t string) bool {
	arr := make([]int, 26)
	for _, r := range s {
		arr[r-'a']++
	}
	for _, r := range []byte(t) {
		arr[r-'a']--
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}
