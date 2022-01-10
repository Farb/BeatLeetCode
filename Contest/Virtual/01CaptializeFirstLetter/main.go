package main

import "strings"

//https://leetcode-cn.com/contest/biweekly-contest-69/problems/capitalize-the-title/
func capitalizeTitle(title string) string {
	arr:=strings.Split(title," ")
	for i := 0; i < len(arr); i++ {
		word:=arr[i]
		if len(word)<=2 {
			arr[i]=strings.ToLower(word)
			continue
		}
		arr[i]=strings.ToUpper(string(rune(word[0])))+strings.ToLower(word[1:])
	}
	return strings.Join(arr," ")
}
