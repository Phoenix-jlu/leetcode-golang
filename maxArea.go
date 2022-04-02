package main

import (
	"fmt"
	"math"
)
//双层循环超时
//func maxArea(height []int) int {
//	var ans int =0
//	for i :=0 ;i<len(height)-1 ;i++  {
//		if(height[i] > ans / len(height)-1-i) {
//			for j := i + 1; j < len(height); j++ {
//				heightitem := int(math.Min(float64(height[i]), float64(height[j])))
//				if (heightitem*(j-i) > ans) {
//					ans = heightitem * (j - i)
//				}
//			}
//		}
//	}
//	return ans
//}

func maxArea(height []int) int {
	var ans int =0

	for i,j:=0,len(height)-1;i<j; {
		if(ans < (j-i)*int(math.Min(float64(height[i]),float64(height[j])))){
			ans = (j-i)*int(math.Min(float64(height[i]),float64(height[j])))
		}
		if(height[i] <= height[j]){
			i++
			continue
		}
		if(height[i] >= height[j]){
			j--
			continue
		}
	}
	return ans
}



func main()  {
	//testInt := []int{1,8,6,2,5,4,8,3,7}
	testInt := []int{4,3,2,1,4}
	fmt.Println(maxArea(testInt))
}
