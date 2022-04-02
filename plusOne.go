package main

import (
	"fmt"
	"strconv"
	"strings"
)

func plusOne(digits []int) []int {
	ans := make([]int,0)
	var tmp string = ""
	if(len(digits)== 1 && digits[0] == 0){
		return []int{1}
	}
	for _,i := range digits{
		tmp += strconv.Itoa(i)
	}
	inttmp,_ := strconv.Atoi(tmp)

	anstmp := inttmp + 1
	strans := strconv.Itoa(anstmp)
	ansslic := strings.Split(strans,"")
	for i:=0;i<len(ansslic);i++{
		tmpInt,_ := strconv.Atoi(ansslic[i])
		ans = append(ans,tmpInt)
	}
	return ans
}

func main(){
	testInt := []int{7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6}
	ans := plusOne(testInt)
	fmt.Print(ans)
}