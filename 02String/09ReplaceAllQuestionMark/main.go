package main

//https://leetcode-cn.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/

func modifyString(s string) string {
	arr := []byte(s)
	questionMark, a := byte('?'), byte('a')
	var baseChar byte
	for i := 0; i < len(arr); i++ {
		if arr[i] != questionMark {
			continue
		}
		baseChar = a
		//i两边的字符不能和i相同，如果相同，就使用下个ASCII字符，直到和两边不同
		for i-1 >= 0 && baseChar == arr[i-1] || i+1 <= len(arr)-1 && baseChar == arr[i+1] {
			baseChar++
		}
		arr[i] = baseChar
	}
	return string(arr)
}

//map占用额外空间
func modifyString_v1(s string) string {
	chMap := make(map[byte]struct{})
	sArr := []byte(s)
	for i := 0; i < len(sArr); i++ {

		if sArr[i] != '?' {
			continue
		}
		if i-1 >= 0 { //?左边右字符
			chMap[sArr[i-1]] = struct{}{}
		}
		if i+1 < len(sArr) {
			chMap[sArr[i+1]] = struct{}{}
		}
		sArr[i] = getChar(chMap)
		chMap = make(map[byte]struct{})
	}
	return string(sArr)
}

func getChar(chMap map[byte]struct{}) byte {
	if _, ok := chMap['a']; !ok {
		return 'a'
	}
	if _, ok := chMap['b']; !ok {
		return 'b'
	}
	return 'c'
}
