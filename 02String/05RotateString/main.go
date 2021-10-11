package main

//https://leetcode-cn.com/problems/rotate-string/
func rotateString(s string, goal string) bool {
	a := []byte(s)
	for i := 0; i < len(a); i++ {
		a = append(a[1:], a[:1]...)
		if string(a) == goal {
			return true
		}
	}
	return false
}

func main() {
	a,b:= "abcde","cdeab"
	println(rotateString(a,b)) //true

	a,b= "abcde","abced"
	println(rotateString(a,b)) // false

	a,b= " ","abced"
	println(rotateString(a,b)) // false
}
