package main

import "fmt"

func climbStairs(n int) int {
	if (n == 1){
		return 1
	}
	if(n == 2){
		return 2
	}
	var a = make([]int,n)
	a[0] = 1
	a[1] = 2
	for i:=2;i<n ;i++  {
		a[i] = a[i-1] + a[i-2]
	}
	return a[n-1]
}

func main()  {
	fmt.Println(climbStairs(3))
}