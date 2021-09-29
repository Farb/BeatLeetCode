package main

import "fmt"

//https://leetcode-cn.com/problems/reverse-string/

//思路：双指针
func reverseString0(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))

	s = []byte("a")
	reverseString(s)
	fmt.Println(string(s))
}
