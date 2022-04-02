package main

import (
	"fmt"
	"sort"
)
//一手顺子
func isNStraightHand(hand []int, groupSize int) bool {
	if(len(hand)%groupSize != 0){
		return false
	}
	var num int
	handStatus := make([]int,len(hand))
	sort.Slice(hand, func(i,j int) bool {
		return hand[i] < hand[j]
	})
	if(groupSize == 1){
		return true
	}
	for(num != len(hand)){
		var first int
		for i:=0;i<len(hand);i++{
			if(handStatus[i] != 1){
				first = hand[i]
				handStatus[i] = 1
				num++
				break
			}
		}
		tmp := groupSize
		for i:=1;i<groupSize;i++{
			for j:=0;j<len(hand);j++{
				if(hand[j] == first + i) && (handStatus[j] != 1){
					num++
					tmp--
					handStatus[j] =1
					break
				}
			}
		}
		if(tmp != 1){
			return false
		}
	}
	return true
}

func main()  {
	//testInt := []int{1,2,3,6,3,3,4,7,8}
	testInt := []int{5,1}
	groupSize := 1
	fmt.Println(isNStraightHand(testInt,groupSize))
}
