package main

import (
	"fmt"
)

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

//https://leetcode-cn.com/problems/reverse-string-ii/
func reverseStr(s string, k int) string {
	sb := []byte(s)
	length := len(s)
	for i := 0; i < length; i+=2*k {
		sub:=sb[i:min(k+i,length)]
		for j := 0; j < len(sub)/2; j++ {
			sub[j], sub[len(sub)-1-j] = sub[len(sub)-1-j], sub[j]
		}
	}
	return string(sb)
}

func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))

	s = []byte("a")
	reverseString(s)
	fmt.Println(string(s))

	s2:= "abcdefg"
	 k := 2
	 fmt.Println(reverseStr(s2,k))

	 s2= "abcd"
	 k = 5
	 fmt.Println(reverseStr(s2,k))
}
