package main

//https://leetcode-cn.com/problems/battleships-in-a-board/

func countBattleships(board [][]byte) int {
	ans := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				searchBattleShips(board,i,j)
				ans++
			}
		}
	}
	return ans
}

func searchBattleShips(board [][]byte, row, col int) {
	if !inMetrix(board,row,col) ||board[row][col]!='X' {
		return
	}

	board[row][col]='.'

	searchBattleShips(board,row-1,col)
	searchBattleShips(board,row+1,col)
	searchBattleShips(board,row,col-1)
	searchBattleShips(board,row,col+1)
}

func inMetrix(board [][]byte, row, col int) bool {
	return 0 <= row && row < len(board) && 0 <= col && col < len(board[0])
}
