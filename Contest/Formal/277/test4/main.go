package main

func maximumGood(words [][]int) int {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[0]); j++ {
			if words[i][j] == 0 {
				return process(words, i, j)
			}
		}
	}
	return 0
}

func process(words [][]int, row, col int) (res int) {

}
