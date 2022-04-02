package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	var ans  = true
	if (x < 0) {
		return false
	}
	var s = strconv.Itoa(x)
	for i,j:=0,len(s)-1;i<=j;{
		if(s[i] != s[j]){
			ans = false
			break
		}
		i++
		j--
	}
	return ans
}

func main(){
	//var test =  [][]byte{{'X','.','.','X'},{'X','.','.','X'},{'.','.','.','X'}}
	var test = 1000000001000000001
	fmt.Println(isPalindrome(test))
}

