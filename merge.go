package main

import (
	"fmt"
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	var ans = make([][]int,0)
	if (len(intervals) == 0 || len(intervals) == 1){
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if(intervals[i][0] != intervals[j][0]){
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
    var start int = intervals[0][0]
    var end int = intervals[0][1]
	for i:=1;i<len(intervals) ;i++  {
		if ( end >= intervals[i][0] ) {
			end = int(math.Max(float64(end),float64(intervals[i][1])))
		}else {
			ans = append(ans,[]int{start,end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	ans = append(ans,[]int{start,end})
	return ans
}

func main(){
	//var test =  [][]int{{1,3},{2,6},{5,10},{15,18}}
	//var test =  [][]int{{1,4},{4,5}}
	var test =  [][]int{{1,4},{0,4}}
	fmt.Println(merge(test))
}