package main

import (
	"fmt"
	"unicode"
)

//https://leetcode-cn.com/problems/XltzEq/
func isPalindrome(s string) bool {
	slice := make([]rune, 0, len(s))
	for _, v := range s {
		if unicode.IsDigit(v) {
			slice = append(slice, v)
		}else if unicode.IsLetter(v) {
			slice = append(slice, unicode.ToLower(v))
		}
	}

	for i := 0; i < len(slice)/2; i++ {
		if slice[i]!=slice[len(slice)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))

	s = "race a car"
	fmt.Println(isPalindrome(s))

	s = "r"
	fmt.Println(isPalindrome(s))

	s = " "
	fmt.Println(isPalindrome(s))
}
