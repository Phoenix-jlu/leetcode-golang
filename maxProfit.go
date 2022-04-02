package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	if (len(prices) == 1){
		return 0
	}
	var dp = make([][]int,len(prices))
	for i:= range dp{
		dp[i] = make([]int,2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i:=1;i<len(prices) ;i++  {
		dp[i][0] = int(math.Max(float64(dp[i-1][0]),float64(dp[i-1][1]+prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]),float64(dp[i-1][0]-prices[i])))
	}
	return dp[len(prices)-1][0]
}

func main(){
	test := []int{1,7}
	fmt.Println(maxProfit(test))
}