package main

import (
	"fmt"
)
func countQuadruplets(nums []int) int {
	var num int
	for i:=0;i<len(nums)-3;i++{
		for j:=i+1;j<len(nums)-2 ;j++{
			for k:=j+1;k<len(nums)-1;k++{
				for z:=k+1;z<len(nums);z++{
					if (nums[i]+nums[j]+nums[k] == nums[z]) {
						num++
					}
				}
			}
		}
	}
	return num
}

func main()  {
	testInt := []int{28,8,49,85,37,90,20,8}
	fmt.Println(countQuadruplets(testInt))
}