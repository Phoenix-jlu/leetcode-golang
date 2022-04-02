package main

import "fmt"

func countBattleships(board [][]byte) int {
	var ans int
	var dp = make([][]int,len(board))
	for i:= range dp{
		dp[i] = make([]int,len(board[0]))
	}
	for i:=0;i<len(board) ;i++ {
		for j:=0;j<len(board[0]);j++ {
			if (string(board[i][j]) == "X" && dp[i][j] == 0) {
				defcountBattleships(board,dp,i,j)
				ans++
			}
		}
	}
	return ans
}

func defcountBattleships(board [][]byte,dp [][]int,i,j int){

	if(i >= len(board) || j >= len(board[0]) || i < 0 || j < 0 || dp[i][j] == 1 || string(board[i][j]) == "." ){
		return
	}
	if( string(board[i][j]) == "X"){
		dp[i][j] = 1
		defcountBattleships(board,dp,i+1,j)
		defcountBattleships(board,dp,i-1,j)
		defcountBattleships(board,dp,i,j+1)
		defcountBattleships(board,dp,i,j-1)
	}
}

func main(){
	//var test =  [][]byte{{'X','.','.','X'},{'X','.','.','X'},{'.','.','.','X'}}
	var test =  [][]byte{{'.'}}
	fmt.Println(countBattleships(test))
}

