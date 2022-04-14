package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	ans := 0
	for i := 0;i<len(accounts);i++{
		var tmp = 0
		for j := 0 ;j<len(accounts[0]);j++{
			tmp = accounts[i][j] + tmp
		}
		if (tmp > ans){
			ans = tmp
		}
	}
	return ans
}


func main(){
	var test =  [][]int{{1,2,3}}
	fmt.Println(maximumWealth(test))
}
