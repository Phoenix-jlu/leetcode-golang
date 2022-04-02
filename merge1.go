package main

import (
	"fmt"
	"sort"
)

func merge1(nums1 []int, m int, nums2 []int, n int)  {
	for i:=0;i<len(nums2) ;i++  {
		nums1[m + i] = nums2[i]
	}
	sort.Slice(nums1, func(i,j int) bool {
		return nums1[i] < nums1[j]
	})
}

func main()  {
	var test1 =  []int{1,2,3,0,0,0}
	var test2 =  []int{2,5,6}
	//var test1 =  []int{0}
	//var test2 =  []int{1}
	merge1(test1,3,test2,3)
	fmt.Println(test1)
}
