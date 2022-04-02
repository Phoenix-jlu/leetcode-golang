package main

import (
	"fmt"
	"math"
)

func maxIncreaseKeepingSkyline(grid [][]int) int {

	var rowSli = make([]int,len(grid))//行
	var colSli = make([]int,len(grid))//列
	var ans int
	//for i:=0 ; i<len(grid) ; i++  {
	//	tmp := 0
	//	for j:=0 ; j<len(grid[i]) ; j++  {
	//		if(tmp < grid[i][j]){
	//			tmp = grid[i][j]
	//		}
	//	}
	//	rowSli = append(rowSli,tmp)
	//}
	//for i:=0 ; i<len(grid[0]) ; i++  {
	//	tmp := 0
	//	for j:=0 ; j<len(grid) ; j++  {
	//		if(tmp < grid[j][i]){
	//			tmp = grid[j][i]
	//		}
	//	}
	//	colSli = append(colSli,tmp)
	//}
	//高端操作
	for i:=0 ; i<len(grid) ; i++  {
		for j:=0 ; j<len(grid[i]) ; j++  {
			rowSli[i] = int(math.Max(float64(rowSli[i]),float64(grid[i][j])))
			colSli[j] = int(math.Max(float64(colSli[j]),float64(grid[i][j])))
		}
	}

	for i:=0 ; i<len(grid) ; i++  {
		for j:=0 ; j<len(grid[i]) ; j++  {
			numMin := int(math.Min(float64(rowSli[i]),float64(colSli[j])))
			ans = ans + numMin - grid[i][j]
 		}
	}
	return ans
}

func main(){
	var test =  [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	fmt.Println(maxIncreaseKeepingSkyline(test))
}
