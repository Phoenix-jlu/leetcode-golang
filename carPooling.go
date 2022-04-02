package main

import (
	"fmt"
	"sort"
)

func carPooling(trips [][]int, capacity int) bool {
	if(len(trips) <= 1){
		return trips[0][0] <= capacity
	}
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][2] < trips[j][2]
	})
	maxIndex := trips[len(trips)-1][2]
	sort.Slice(trips, func(i, j int) bool {
			return trips[i][1] < trips[j][1]
	})
	minIndex := trips[0][1]
	tripsStatus := make([]int,len(trips))
	var tmp int
	for i:=minIndex;i<=maxIndex;i++{
		for j:=0;j<len(trips);j++  {
			if(trips[j][1] <= i && trips[j][2] > i  && tripsStatus[j] == 0){
				tmp = tmp + trips[j][0]
				tripsStatus[j] = 1
				if(tmp > capacity){
					return false
				}
				continue
			}
			if(trips[j][2] <= i  && tripsStatus[j] == 1){
				tmp = tmp - trips[j][0]
				tripsStatus[j] = 2
				if(tmp > capacity){
					return false
				}
				continue
			}
		}
	}
	return true
}

func main(){
	var test =  [][]int{{3,2,8},{4,4,6},{10,8,9}}
	//var test =  [][]int{{3,2,7},{3,7,9},{8,3,9}}
	//var test =  [][]int{{2,1,5},{3,5,7}}
	fmt.Println(carPooling(test,11))
}


