package main

import (
	"fmt"
	"unicode"
)

func detectCapitalUse(word string) bool {
	if len(word)==1 {
		return true
	}
	if unicode.IsLower(rune(word[0])) {
		return isAllLower(word[1:])
	}
	if unicode.IsUpper(rune(word[len(word)-1])) {
		return isAllUpper(word[1 : len(word)-1])
	}
	return isAllLower(word[1 : len(word)-1])
}

func isAllLower(s string) bool {
	for _, v := range s {
		if !unicode.IsLower(v) {
			return false
		}
	}
	return true
}

func isAllUpper(s string) bool {
	for _, v := range s {
		if unicode.IsLower(v) {
			return false
		}
	}
	return true
}

// https://leetcode-cn.com/problems/detect-capital/
func main() {
	word := "USA"
	fmt.Println(detectCapitalUse(word)) // true

	word = "FlaG"
	fmt.Println(detectCapitalUse(word)) // false

	word = "abc"
	fmt.Println(detectCapitalUse(word)) // true

	word = "a"
	fmt.Println(detectCapitalUse(word)) // true

	word = "AB"
	fmt.Println(detectCapitalUse(word)) // true

	word = "A"
	fmt.Println(detectCapitalUse(word)) // true
}
