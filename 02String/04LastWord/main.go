package main

import "strings"

func lengthOfLastWord(s string) int {
	arr:=strings.Split(strings.Trim(s," ")," ")
	return len(arr[len(arr)-1])
}

func main() {
	s:= "Hello World"
	println(lengthOfLastWord(s))

	s = "   fly me   to   the moon  "
	println(lengthOfLastWord(s))

	s = "luffy is still joyboy"
	println(lengthOfLastWord(s))
}